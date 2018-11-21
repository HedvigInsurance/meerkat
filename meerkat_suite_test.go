package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMeerkat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Meerkat Suite")
}
