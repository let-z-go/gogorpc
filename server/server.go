package server

import (
	"context"
	"errors"
	"net"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/let-z-go/pbrpc/channel"
)

type Server struct {
	options  *Options
	rawURL   string
	ctx      context.Context
	cancel   context.CancelFunc
	shutdown shutdown
}

func (self *Server) Init(options *Options, rawURL string) *Server {
	self.options = options.Normalize()
	self.rawURL = rawURL
	self.ctx, self.cancel = context.WithCancel(context.Background())
	self.shutdown.Init(self.options.ShutdownTimeout, self.ctx)
	return self
}

func (self *Server) Close() {
	self.cancel()
	self.shutdown.Close()
}

func (self *Server) Run() (err error) {
	if self.shutdown.IsClosed() {
		self.options.Logger.Error().
			Str("server_url", self.rawURL).
			Msg("server_already_closed")
		return ErrClosed
	}

	atomic.AddInt32(&self.shutdown.Counter, 1)

	defer func() {
		self.options.Logger.Info().
			Str("server_url", self.rawURL).
			Msg("server_closed")

		if self.shutdown.IsClosed() {
			err = ErrClosed
		} else {
			self.cancel()
		}

		atomic.AddInt32(&self.shutdown.Counter, -1)
	}()

	if self.shutdown.IsClosed() {
		return
	}

	url_, err := url.Parse(self.rawURL)

	if err != nil {
		self.options.Logger.Error().
			Err(err).
			Str("server_url", self.rawURL).
			Msg("server_invalid_url")
		return err
	}

	acceptor, err := GetAcceptor(url_.Scheme)

	if err != nil {
		self.options.Logger.Error().
			Err(err).
			Str("server_url", self.rawURL).
			Msg("server_bad_scheme")
		return err
	}

	err = acceptor(self.ctx, url_, &self.shutdown.Counter, func(connection net.Conn) {
		channel_ := new(channel.Channel).Init(self.options.Channel)
		defer channel_.Close()

		if err := channel_.Accept(self.shutdown.Ctx, connection); err != nil {
			self.options.Logger.Warn().Err(err).
				Str("server_url", self.rawURL).
				Str("transport_id", channel_.GetTransportID().String()).
				Msg("server_channel_accept_failed")
			return
		}

		err := channel_.Process(self.shutdown.Ctx)
		self.options.Logger.Warn().Err(err).
			Str("server_url", self.rawURL).
			Str("transport_id", channel_.GetTransportID().String()).
			Msg("server_channel_process_failed")
	})

	self.options.Logger.Error().Err(err).
		Str("server_url", self.rawURL).
		Msg("server_accept_failed")
	return err
}

func (self *Server) WaitForShutdown() bool {
	return self.shutdown.WaitFor()
}

var ErrClosed = errors.New("pbrpc/server: closed")

type shutdown struct {
	Ctx     context.Context
	Counter int32

	isClosed int32
}

func (self *shutdown) Init(timeout time.Duration, ctx context.Context) {
	if timeout < 0 {
		self.Ctx = context.Background()
	} else {
		var cancel context.CancelFunc
		self.Ctx, cancel = context.WithCancel(context.Background())

		go func() {
			select {
			case <-ctx.Done():
				time.Sleep(timeout)
				cancel()
			}
		}()
	}
}

func (self *shutdown) Close() {
	atomic.StoreInt32(&self.isClosed, 1)
}

func (self *shutdown) WaitFor() bool {
	if self.Ctx.Err() != nil {
		return false
	}

	ticker := time.NewTicker(shutdownPollInterval)
	defer ticker.Stop()

	for {
		if atomic.LoadInt32(&self.Counter) == 0 && self.IsClosed() {
			return true
		}

		select {
		case <-self.Ctx.Done():
			return false
		case <-ticker.C:
		}
	}
}

func (self *shutdown) IsClosed() bool {
	return atomic.LoadInt32(&self.isClosed) == 1
}

const shutdownPollInterval = 500 * time.Millisecond