package swapi_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSwapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Swapi Suite")
}
