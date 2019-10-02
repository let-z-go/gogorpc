package stream

import (
	"sync"
	"time"

	"github.com/rs/zerolog"

	"github.com/let-z-go/gogorpc/internal/transport"
)

type Options struct {
	Transport                 *transport.Options
	Logger                    *zerolog.Logger
	ActiveHangupTimeout       time.Duration
	IncomingKeepaliveInterval time.Duration
	OutgoingKeepaliveInterval time.Duration
	IncomingConcurrencyLimit  int
	OutgoingConcurrencyLimit  int

	normalizeOnce sync.Once
}

func (self *Options) Normalize() *Options {
	self.normalizeOnce.Do(func() {
		if self.Transport == nil {
			self.Transport = &defaultTransportOptions
		}

		self.Transport.Normalize()

		if self.Logger == nil {
			self.Logger = self.Transport.Logger
		}

		normalizeDurValue(&self.ActiveHangupTimeout, defaultActiveHangupTimeout, minActiveHangupTimeout, maxActiveHangupTimeout)
		normalizeDurValue(&self.IncomingKeepaliveInterval, defaultKeepaliveInterval, minKeepaliveInterval, maxKeepaliveInterval)
		normalizeDurValue(&self.OutgoingKeepaliveInterval, defaultKeepaliveInterval, minKeepaliveInterval, maxKeepaliveInterval)
		normalizeIntValue(&self.IncomingConcurrencyLimit, defaultConcurrencyLimit, minConcurrencyLimit, maxConcurrencyLimit)
		normalizeIntValue(&self.OutgoingConcurrencyLimit, defaultConcurrencyLimit, minConcurrencyLimit, maxConcurrencyLimit)
	})

	return self
}

const (
	defaultActiveHangupTimeout = 3 * time.Second
	minActiveHangupTimeout     = 1 * time.Second
	maxActiveHangupTimeout     = 5 * time.Second
)

const (
	defaultKeepaliveInterval = 15 * time.Second
	minKeepaliveInterval     = 5 * time.Second
	maxKeepaliveInterval     = 60 * time.Second
)

const (
	defaultConcurrencyLimit = 1 << 17
	minConcurrencyLimit     = 1
	maxConcurrencyLimit     = 1 << 20
)

var defaultTransportOptions transport.Options

func normalizeDurValue(value *time.Duration, defaultValue, minValue, maxValue time.Duration) {
	if *value == 0 {
		*value = defaultValue
		return
	}

	if *value < minValue {
		*value = minValue
		return
	}

	if *value > maxValue {
		*value = maxValue
		return
	}
}

func normalizeIntValue(value *int, defaultValue, minValue, maxValue int) {
	if *value == 0 {
		*value = defaultValue
		return
	}

	if *value < minValue {
		*value = minValue
		return
	}

	if *value > maxValue {
		*value = maxValue
		return
	}
}
