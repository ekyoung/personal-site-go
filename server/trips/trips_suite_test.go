package trips_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/reporters"
	"testing"

	"os"
	"path/filepath"
)

func TestTrips(t *testing.T) {
	RegisterFailHandler(Fail)
	testReportsPath, _ := filepath.Abs("../../test-reports")
	os.MkdirAll(testReportsPath, 0644)
	junitReporter := reporters.NewJUnitReporter(filepath.Join(testReportsPath, "server-trips-junit.xml"))
	RunSpecsWithDefaultAndCustomReporters(t, "Trips Suite", []Reporter{junitReporter})
}
