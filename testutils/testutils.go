package testutils

import (
	"encoding/json"
	"io/ioutil"

	"go.i3wm.org/i3/v4"
)

// TreeFromJSONFile loads a i3.Tree from a JSON file
func TreeFromJSONFile(filepath string) *i3.Tree {
	data, err := ioutil.ReadFile(filepath)
	must(err)

	return unmarshalTree(data)
}

func unmarshalTree(b []byte) *i3.Tree {
	var root i3.Node
	must(json.Unmarshal(b, &root))

	return &i3.Tree{Root: &root}
}

func MarshalTree(t *i3.Tree) []byte {
	data, err := json.Marshal(t)
	must(err)

	return data
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func MaybeWriteGolden(original string, data []byte, update bool) string {
	if update {
		golden := original + ".golden"
		must(ioutil.WriteFile(golden, data, 0640))
	}

	// Read file
	read, err := ioutil.ReadFile(original)
	must(err)

	return string(read)
}
