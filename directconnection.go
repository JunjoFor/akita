package akita

import (
	"log"
	"sync"
)

// DirectConnection connects two components without latency
type DirectConnection struct {
	sync.Mutex
	*HookableBase

	endPoints map[Port]bool
	engine    Engine
}

// PlugIn marks the port connects to this DirectConnection.
func (c *DirectConnection) PlugIn(port Port) {
	c.Lock()
	defer c.Unlock()

	c.endPoints[port] = true
	port.SetConnection(c)
}

// Unplug marks the port no longer connects to this DirectConnection.
func (c *DirectConnection) Unplug(port Port) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.endPoints[port]; !ok {
		log.Panicf("connectable if not attached")
	}

	delete(c.endPoints, port)
	port.SetConnection(c)
}

// NotifyAvailable is called by a port to notify that the connection can
// deliver to the port again.
func (c *DirectConnection) NotifyAvailable(now VTimeInSec, port Port) {
	for p := range c.endPoints {
		p.NotifyAvailable(now)
	}
}

// Send of a DirectConnection schedules a DeliveryEvent immediately
func (c *DirectConnection) Send(msg Msg) *SendError {
	if msg.Meta().Dst == nil {
		log.Panic("destination is null")
	}

	// if _, found := c.endPoints[msg.Dst()]; !found {
	// 	log.Panicf("destination %s not connected, "+
	// 		"msg ID %s, "+
	// 		"msguest from %s",
	// 		msg.Dst().Comp.Name(),
	// 		msg.GetID(),
	// 		msg.Dst().Comp.Name(),
	// 	)
	// }

	msg.Meta().RecvTime = msg.Meta().SendTime
	return msg.Meta().Dst.Recv(msg)
}

// Handle defines how the DirectConnection handles events
func (c *DirectConnection) Handle(evt Event) error {
	return nil
}

// NewDirectConnection creates a new DirectConnection object
func NewDirectConnection(engine Engine) *DirectConnection {
	c := new(DirectConnection)
	c.HookableBase = NewHookableBase()
	c.endPoints = make(map[Port]bool)
	c.engine = engine
	return c
}
