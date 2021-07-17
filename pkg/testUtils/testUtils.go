package testUtils

import "testing"

func FailOnError(t *testing.T, err error, reason string) {
	t.Logf("%s, found %v", reason, err)
	t.Fail()
}
