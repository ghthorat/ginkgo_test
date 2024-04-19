package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInterview(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interview Suite")
}
