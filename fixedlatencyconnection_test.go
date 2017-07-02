package core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gitlab.com/yaotsu/core"
)

var _ = Describe("FixedLatencyConnection", func() {

	var (
		comp1      *core.MockComponent
		comp2      *core.MockComponent
		comp3      *core.MockComponent
		freq       core.Freq
		connection *core.FixedLatencyConnection
		engine     *core.MockEngine
	)

	BeforeEach(func() {
		comp1 = core.NewMockComponent("comp1")
		comp2 = core.NewMockComponent("comp2")
		comp3 = core.NewMockComponent("comp3")
		engine = core.NewMockEngine()

		freq = core.GHz
		latency := 2
		connection = core.NewFixedLatencyConnection(engine, latency, freq)
		connection.Attach(comp1)
		connection.Attach(comp2)
	})

	It("should give error is detaching a not attached component", func() {
		Expect(func() { connection.Detach(comp3) }).To(Panic())
	})

	It("should detach", func() {
		// Normal detaching
		Expect(func() { connection.Detach(comp1) }).NotTo(Panic())

		// Detaching again should give error
		Expect(func() { connection.Detach(comp1) }).To(Panic())
	})

	// It("should send with latency", func() {
	// 	req := NewMockRequest()
	// 	req.SetSrc(comp2)
	// 	req.SetDst(comp1)
	// 	req.SetSendTime(2.0)

	// 	errToRet := core.NewError("something", true, 10)
	// 	comp1.ToReceiveReq(req, errToRet)

	// 	err := connection.Send(req)

	// 	Expect(err).To(BeIdenticalTo(errToRet))
	// })

})
