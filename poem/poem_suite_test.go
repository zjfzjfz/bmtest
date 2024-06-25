package poem_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPoem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Poem Suite")
}
