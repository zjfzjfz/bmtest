package calculate_test

import (
	"bmtest/utility"
	"net/http"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Score API", Ordered, func() {
	var client *resty.Client
	var token string

	BeforeAll(func() {
		client = resty.New()
		token = utility.GetToken()
	})

	Context("sending score request", func() {
		It("valid", func() {
			resp, err := client.R().
				SetHeader("Authorization", token).
				Get("http://47.96.125.121:8080/user/math/score")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})
	})
})
