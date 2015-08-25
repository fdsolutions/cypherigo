package cypherigo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCypherigo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cypherigo Suite")
}
