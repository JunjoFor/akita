package conn

import (
	"fmt"
)

// DirectConnection provides a way to connect two component directly so that
// no latency would happen.
type DirectConnection struct {
	*BasicConn
}

// NewDirectConnection creates a new DirectConnection object
func NewDirectConnection() *DirectConnection {
	c := DirectConnection{NewBasicConn()}
	return &c
}

// CanSend of the DirectConnection only checks if the receiver can process the
// request.
func (c *DirectConnection) CanSend(req *Request) *ConnError {
	_, ok := c.BasicConn.connectables[req.From]
	if !ok {
		return &ConnError{"Source " + req.From.Name() + " is not connected",
			false, 0}
	}

	dst, err := c.getDest(req)
	if err != nil {
		_ = fmt.Errorf("%v", err)
		return &ConnError{err.Error(), false, 0}
	}

	return dst.CanRecv(req)
}

// Send of a DirectConnection invokes receiver's Recv method
func (c *DirectConnection) Send(req *Request) *ConnError {
	if req.To == nil {
		return NewConnError("Destination of a request is not known.", false, 0)
	}
	return req.To.Recv(req)
}
