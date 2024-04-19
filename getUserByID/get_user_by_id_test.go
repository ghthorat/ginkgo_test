package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GET https://reqres.in/api/users/1", func() {
	Context("When sending a GET request to https://reqres.in/api/users/1", func() {
		It("should return status code 200 and user details", func() {
			// Create a request to the API endpoint
			req, err := http.NewRequest("GET", "https://reqres.in/api/users/1", nil)
			Expect(err).NotTo(HaveOccurred())

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Serve the request to the handler
			handler := http.DefaultServeMux
			handler.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusOK))

			// Parse the response body
			var responseBody map[string]interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
			Expect(err).NotTo(HaveOccurred())

			// Validate user details
			Expect(responseBody).To(HaveKeyWithValue("data", HaveKeyWithValue("id", 1)))
			Expect(responseBody).To(HaveKeyWithValue("data", HaveKeyWithValue("email", "george.bluth@reqres.in")))
			Expect(responseBody).To(HaveKeyWithValue("data", HaveKeyWithValue("first_name", "George")))
			Expect(responseBody).To(HaveKeyWithValue("data", HaveKeyWithValue("last_name", "Bluth")))
			Expect(responseBody).To(HaveKeyWithValue("data", HaveKeyWithValue("avatar", "https://reqres.in/img/faces/1-image.jpg")))
		})
	})
})
