package wandgo

import (
	"fmt"
	"net/url"

	"github.com/joeshaw/iso8601"
)

type Permlink struct {
	Parameter *PermlinkParam `json:"parameter"`
	Result    *CompileResult `json:"result"`
}

type PermlinkParam struct {
	CompileParam
	CreatedAt *iso8601.Time `json:"created-at"`
}

func (a *API) GetPermlink(link string) (*Permlink, error) {
	endpoint := fmt.Sprintf("%s/permlink/%s", BaseURL, link)
	v := url.Values{}
	res := &Permlink{}

	err := a.apiGet(endpoint, v, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
