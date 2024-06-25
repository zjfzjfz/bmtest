package user_test

import (
	"encoding/json"
	"net/http"

	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Login API", func() {
	var client *resty.Client

	BeforeEach(func() {
		client = resty.New()
	})

	Context("sending login request", func() {
		It("invalid", func() {
			payload := map[string]interface{}{
				"email":    "christianlsl@foxmail.com",
				"password": "1234",
			}

			resp, _ := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(payload).
				Post("http://47.96.125.121:8080/user/login")

			Expect(resp.StatusCode()).To(Equal(http.StatusOK))

			var result map[string]interface{}
			_ = json.Unmarshal(resp.Body(), &result)
			Expect(result["success"]).To(BeFalse())
		})

		It("valid", func() {
			payload := map[string]interface{}{
				"email":    "zjf15821145685@163.com",
				"password": "1234",
			}

			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(payload).
				Post("http://47.96.125.121:8080/user/login")

			Expect(err).To(BeNil())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))

			var result map[string]interface{}
			err = json.Unmarshal(resp.Body(), &result)

			Expect(err).To(BeNil())
			Expect(result["success"]).To(BeTrue())
			Expect(result["data"]).To(Not(BeNil()))
		})
	})
})
