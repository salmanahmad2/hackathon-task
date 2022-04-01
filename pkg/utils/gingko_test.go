package utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
This test suite also serve as a demo for how to use Ginkgo Test Framework
https://github.com/onsi/ginkgo
API Documentation: https://godoc.org/github.com/onsi/ginkgo
*/

// Entry point for the Ginkgo Tests. Only one per package is necessary.
func TestGoUtilsSuite(t *testing.T) {
	// Connects Ginkgo to Gomega (the library for making assertions)
	RegisterFailHandler(Fail)
	// Runs all the Ginkgo Tests (var _ = Describe(...))
	RunSpecs(t, "Go Utils Test Suite")
}
