package signature_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJws(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jws Suite")
}
