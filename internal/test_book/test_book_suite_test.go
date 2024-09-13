package test_book_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTestBook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestBook Suite")
}
