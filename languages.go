package wandgo

import "net/url"

type Language struct {
	CompilerOptionRaw     bool   `json:"compiler-option-raw"`
	DisplayCompileCommand string `json:"display-compile-command"`
	Language              string `json:"language"`
	Name                  string `json:"name"`
	RuntimeOptionRaw      bool   `json:"runtime-option-raw"`
	Version               string `json"version"`
	// Switches
}

func (a *API) GetLanguages() ([]Language, error) {
	langs := make([]Language, 0)
	err := a.apiGet(BaseURL+"/list.json", url.Values{}, &langs)
	if err != nil {
		return nil, err
	}

	return langs, nil
}
