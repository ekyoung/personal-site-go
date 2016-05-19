package trips_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/reporters"
	"testing"
)

func TestTrips(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter(`..\..\test-reports\server-trips-junit.xml`)
	RunSpecsWithDefaultAndCustomReporters(t, "Trips Suite", []Reporter{junitReporter})
}
