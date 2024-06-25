package user_test

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Register API", func() {
	var client *resty.Client

	BeforeEach(func() {
		client = resty.New()
	})

	Context("sending registration request", func() {
		It("invalid", func() {
			email := "sjjj@163.com"

			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetQueryParam("email", email).
				Post("http://47.96.125.121:8080/user/code")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))

			var result map[string]interface{}
			err = json.Unmarshal(resp.Body(), &result)
			Expect(err).To(BeNil())
			Expect(result["success"]).To(BeTrue())

			payload := map[string]interface{}{
				"email":     email,
				"code":      "",
				"nick_name": "sjj",
				"password":  "1234",
				"age":       22,
			}

			resp, err = client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(payload).
				Post("http://47.96.125.121:8080/user/register")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})
	})
})
