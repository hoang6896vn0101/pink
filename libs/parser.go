package libs

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// YAMLParse function
// Arguments:
// 1. out -> instance of struct.
// 2. path -> file path.
// Return:
// - void
func YAMLParse(out interface{}, path string) {
	_byte, err := ioutil.ReadFile(path)
	if err != nil {
		panic("error")
	}
	err = yaml.Unmarshal(_byte, out)
	if err != nil {
		panic("error")
	}
}
