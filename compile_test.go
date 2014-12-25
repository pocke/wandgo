package wandgo_test

import (
	"testing"

	"github.com/pocke/wandgo"
)

func TestCompile(t *testing.T) {
	p := &wandgo.CompileParam{
		Code:              "#include <iostream>\nint main() { int x = 0; std::cout << \"hoge\" << std::endl; }",
		Options:           "warning,gnu++1y",
		Compiler:          "gcc-head",
		CompilerOptionRaw: "-Dx=hogefuga\n-O3",
	}

	res, err := api.Compile(p)
	if err != nil {
		t.Logf("Response: %+v", res)
		t.Error(err)
	}
}

func TestCompileWhenLanguageDoesNotExists(t *testing.T) {
	p := &wandgo.CompileParam{
		Code:     `puts "hoge"`,
		Compiler: `rubyyyyy`,
	}

	res, err := api.Compile(p)
	t.Log(err)
	if err == nil {
		t.Logf("Response: %+v", res)
		t.Error("rubyyyyy doesn't exists. But not return error")
	}
}
