package wandgo_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/pocke/wandgo"
)

var api = wandgo.NewAPI()

func TestSiwtchesUnmarshalJSON(t *testing.T) {
	jsonStr := `
[
  {
    "default": true,
    "display-flags": "-I/usr/local/sprout",
    "display-name": "Sprout",
    "name": "sprout",
    "type": "single"
  },
  {
    "default": "gnu++1y",
    "options": [
    {
      "display-flags": "-std=c++98 -pedantic",
      "display-name": "C++03",
      "name": "c++98"
    },
    {
      "display-flags": "-std=gnu++98",
      "display-name": "C++03(GNU)",
      "name": "gnu++98"
    },
    {
      "display-flags": "-std=c++11 -pedantic",
      "display-name": "C++11",
      "name": "c++11"
    },
    {
      "display-flags": "-std=gnu++11",
      "display-name": "C++11(GNU)",
      "name": "gnu++11"
    },
    {
      "display-flags": "-std=c++1y -pedantic",
      "display-name": "C++1y",
      "name": "c++1y"
    },
    {
      "display-flags": "-std=gnu++1y",
      "display-name": "C++1y(GNU)",
      "name": "gnu++1y"
    }
    ],
    "type": "select"
  }
]`
	sws := &wandgo.Switches{}

	err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(sws)
	if err != nil {
		t.Log("%+v\n", sws)
		t.Error(err)
	}

	if len(sws.Single) != 1 || sws.Single[0].Name != "sprout" {
		t.Errorf("Unmarshal of Switches.Single failed")
	}

	if len(sws.Select) != 1 || len(sws.Select[0].Options) == 0 {
		t.Errorf("Unmarshal of Switches.Select failed")
	}
}

func TestGetLanguages(t *testing.T) {
	langs, err := api.GetLanguages()
	if err != nil {
		t.Logf("langs: %+v\n", langs)
		t.Error(err)
	}
}
