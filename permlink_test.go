package wandgo_test

import "testing"

func TestGetPermlink(t *testing.T) {
	perm, err := api.GetPermlink("e2NwL3avMIW0I1sS")
	if err != nil {
		t.Logf("Permlink: %+v", perm)
		t.Errorf("%+v", err)
	}
}
