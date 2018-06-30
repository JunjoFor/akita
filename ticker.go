package core

import "sync"

// TickEvent is a generic event that almost all the component can use to
// update their status.
type TickEvent struct {
	*EventBase
}

// NewTickEvent creates a newly created TickEvent
func NewTickEvent(t VTimeInSec, handler Handler) *TickEvent {
	evt := new(TickEvent)
	evt.EventBase = NewEventBase(t, handler)
	return evt
}

// Ticker is a tool that helps a component that executes in a tick-tick fashion
type Ticker struct {
	sync.Mutex
	handler Handler
	freq    Freq
	engine  Engine
	tick    *TickEvent
}

func NewTicker(handler Handler, engine Engine, freq Freq) *Ticker {
	ticker := new(Ticker)

	ticker.handler = handler
	ticker.engine = engine
	ticker.freq = freq

	ticker.tick = NewTickEvent(-1, handler)
	return ticker
}

func (t *Ticker) TickLater(now VTimeInSec) {
	t.Lock()
	defer t.Unlock()

	time := t.freq.NextTick(now)

	if t.tick.Time() >= time {
		return
	}

	t.tick.SetTime(time)
	t.engine.Schedule(t.tick)
}
