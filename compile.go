package wandgo

type CompileParam struct {
	Compiler          string `json:"compiler"`
	Code              string `json:"code"`
	Options           string `json:"options"`
	Stdin             string `json:"stdin"`
	CompilerOptionRaw string `json:"compiler-option-raw"`
	RuntimeOptionRaw  string `json:"runtime-option-raw"`
	save              bool   `json:"false"`
}

type CompileResult struct {
	Status          string `json:"status"`
	Signal          string `json:"signal"`
	CompilerOutput  string `json:"compiler_output"`
	CompilerError   string `json:"compiler_error"`
	CompilerMessage string `json:"compiler_message"`
	ProgramOutput   string `json:"program_output"`
	ProgramError    string `json:"program_error"`
	ProgramMessage  string `json:"program_message"`
	Permlink        string `json:"permlink"`
	URL             string `json:"url"`
}

func (a *API) Compile(p *CompileParam) (*CompileResult, error) {
	res := &CompileResult{}

	err := a.apiPostAsJSON(BaseURL+"/compile.json", p, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
