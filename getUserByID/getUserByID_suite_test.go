package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGetUserByID(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GetUserByID Suite")
}
