package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add users", func() {
	Context("When sending a POST request to https://reqres.in/api/users", func() {
		It("should create a new user and return status code 201", func() {
			// Define the JSON payload
			payload := map[string]string{
				"name": "Ganesh",
				"job":  "Automation Tester",
			}

			// Convert payload to JSON
			jsonData, err := json.Marshal(payload)
			Expect(err).NotTo(HaveOccurred())

			// Create a request to the API endpoint with the JSON payload
			req, err := http.NewRequest("POST", "https://reqres.in/api/users", bytes.NewBuffer(jsonData))
			Expect(err).NotTo(HaveOccurred())

			// Set Content-Type header
			req.Header.Set("Content-Type", "application/json")

			// Send the request
			client := &http.Client{}
			resp, err := client.Do(req)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			// Check the status code
			Expect(resp.StatusCode).To(Equal(http.StatusCreated))
		})
	})
})
