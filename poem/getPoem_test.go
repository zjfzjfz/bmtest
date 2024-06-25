package poem_test

import (
	"bmtest/utility"
	"net/http"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Poem API", Ordered, func() {
	var client *resty.Client
	var token string

	BeforeAll(func() {
		client = resty.New()
		token = utility.GetToken()
	})

	Context("sending poem request", func() {
		It("level-1", func() {
			resp, err := client.R().
				SetHeader("Authorization", token).
				Get("http://47.96.125.121:8080/user/poem/kid")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			Expect(resp.Body()).To(Not(BeNil()))
		})

		It("level-2", func() {
			resp, err := client.R().
				SetHeader("Authorization", token).
				Get("http://47.96.125.121:8080/user/poem/teen")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			Expect(resp.Body()).To(Not(BeNil()))
		})
	})
})
