package calculate_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCalculate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculate Suite")
}
