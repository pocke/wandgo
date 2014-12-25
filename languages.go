package wandgo

import (
	"fmt"
	"net/url"

	"github.com/dustin/go-jsonpointer"
)

type Language struct {
	CompilerOptionRaw     bool      `json:"compiler-option-raw"`
	DisplayCompileCommand string    `json:"display-compile-command"`
	Language              string    `json:"language"`
	Name                  string    `json:"name"`
	RuntimeOptionRaw      bool      `json:"runtime-option-raw"`
	Version               string    `json"version"`
	Switches              *Switches `json:"switches"`
}

type Switches struct {
	Single []SingleSwitch
	Select []SelectSwitch
}

type SingleSwitch struct {
	Default bool   `json:"default"`
	Type    string `json:"type"` // single
	Option
}

type SelectSwitch struct {
	Default string   `json:"default"`
	Type    string   `json:"type"` // select
	Options []Option `json:"options"`
}

type Option struct {
	DisplayFlags string `json:"display-flags"`
	DisplayName  string `json:"display-name"`
	Name         string `json:"name"`
}

func (s *Switches) UnmarshalJSON(b []byte) error {
	for i := int64(0); true; i++ {
		pathRoot := fmt.Sprintf("/%d", i)
		pathType := pathRoot + "/type"
		var t string
		err := jsonpointer.FindDecode(b, pathType, &t)
		if err != nil {
			break
		}

		switch t {
		case "single":
			var v SingleSwitch
			err := jsonpointer.FindDecode(b, pathRoot, &v)
			if err != nil {
				return err
			}
			s.Single = append(s.Single, v)
		case "select":
			var v SelectSwitch
			err := jsonpointer.FindDecode(b, pathRoot, &v)
			if err != nil {
				return err
			}
			s.Select = append(s.Select, v)
		default:
			return fmt.Errorf("%s is Unknown type", t)
		}
	}

	return nil
}

func (a *API) GetLanguages() ([]Language, error) {
	langs := make([]Language, 0)
	err := a.apiGet(BaseURL+"/list.json", url.Values{}, &langs)
	if err != nil {
		return nil, err
	}

	return langs, nil
}
