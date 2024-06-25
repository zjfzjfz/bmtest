package calculate_test

import (
	"bmtest/utility"
	"net/http"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Math API", Ordered, func() {
	var client *resty.Client
	var token string

	BeforeAll(func() {
		client = resty.New()
		token = utility.GetToken()
	})

	Context("sending math request", func() {
		It("level-1", func() {
			resp, err := client.R().
				SetHeader("Authorization", token).
				SetPathParam("level", "1").
				Get("http://47.96.125.121:8080/user/math/{level}")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			Expect(resp.Body()).To(Not(BeNil()))
		})

		It("level-2", func() {
			resp, err := client.R().
				SetHeader("Authorization", token).
				SetPathParam("level", "2").
				Get("http://47.96.125.121:8080/user/math/{level}")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			Expect(resp.Body()).To(Not(BeNil()))
		})

		It("level-3", func() {
			resp, err := client.R().
				SetHeader("Authorization", token).
				SetPathParam("level", "3").
				Get("http://47.96.125.121:8080/user/math/{level}")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			Expect(resp.Body()).To(Not(BeNil()))
		})

		It("level-4", func() {
			resp, err := client.R().
				SetHeader("Authorization", token).
				SetPathParam("level", "4").
				Get("http://47.96.125.121:8080/user/math/{level}")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			Expect(resp.Body()).To(Not(BeNil()))
		})
	})
})
