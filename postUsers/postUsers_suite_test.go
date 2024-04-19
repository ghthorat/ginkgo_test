package main_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPostUsers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PostUsers Suite")
}
