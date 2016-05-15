package trips_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTrips(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trips Suite")
}
