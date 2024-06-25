package poem_test

import (
	"bmtest/utility"
	"net/http"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Record API", Ordered, func() {
	var client *resty.Client
	var token string

	BeforeAll(func() {
		client = resty.New()
		token = utility.GetToken()
	})

	Context("sending record request", func() {
		It("valid", func() {
			resp, err := client.R().
				SetHeader("Authorization", token).
				SetQueryParam("correctRate", "0.80").
				Post("http://47.96.125.121:8080/user/poem/record")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})
	})
})
