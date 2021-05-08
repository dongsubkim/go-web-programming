package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test824TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "824TestGinkgo Suite")
}
