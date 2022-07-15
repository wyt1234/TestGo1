package util

import "testing"

func TestNewPerfTimer(t *testing.T) {
	perfTimer := NewPerfTimer("celery create client")
	defer perfTimer.SoFar()
	// do something
}
