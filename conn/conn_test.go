package conn_test

import "testing"
import "github.com/onsi/gomega"
import "github.com/onsi/ginkgo"

func TestConn(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Request System")
}
