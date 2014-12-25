package wandgo_test

import (
	"testing"

	"github.com/pocke/wandgo"
)

var api = wandgo.NewAPI()

func TestGetLanguages(t *testing.T) {
	langs, err := api.GetLanguages()
	if err != nil {
		t.Logf("langs: %+v", langs)
		t.Error(err)
	}
}
