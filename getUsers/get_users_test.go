package main_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Define a struct for the User object
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

var _ = Describe("GET users", func() {
	Context("When sending a GET request to https://reqres.in/api/users", func() {
		It("should return status code 200 and a list of users", func() {
			// Create a request to the API endpoint
			req, err := http.NewRequest("GET", "https://reqres.in/api/users", nil)
			Expect(err).NotTo(HaveOccurred())

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Serve the request to the handler
			handler := http.DefaultServeMux
			handler.ServeHTTP(rr, req)

			// Check the status code
			Expect(rr.Code).To(Equal(http.StatusOK))

			// Read the response body
			body, err := ioutil.ReadAll(rr.Body)
			Expect(err).NotTo(HaveOccurred())

			// Define a struct to unmarshal the JSON response into
			type UsersResponse struct {
				Page       int    `json:"page"`
				PerPage    int    `json:"per_page"`
				Total      int    `json:"total"`
				TotalPages int    `json:"total_pages"`
				Data       []User `json:"data"`
			}

			// Unmarshal the JSON response
			var usersResponse UsersResponse
			err = json.Unmarshal(body, &usersResponse)
			Expect(err).NotTo(HaveOccurred())

			// Verify the contents of the response
			Expect(usersResponse.Page).To(BeNumerically(">", 0))
			Expect(usersResponse.PerPage).To(BeNumerically(">", 0))
			Expect(usersResponse.Total).To(BeNumerically(">=", len(usersResponse.Data)))
			Expect(usersResponse.TotalPages).To(BeNumerically(">", 0))

			// Check if at least one user is returned
			Expect(len(usersResponse.Data)).To(BeNumerically(">", 0))

			// Verify the structure of the first user
			firstUser := usersResponse.Data[0]
			Expect(firstUser.ID).To(BeNumerically(">", 0))
			Expect(firstUser.Email).NotTo(BeEmpty())
			Expect(firstUser.FirstName).NotTo(BeEmpty())
			Expect(firstUser.LastName).NotTo(BeEmpty())
			Expect(firstUser.Avatar).NotTo(BeEmpty())
		})
	})
})
