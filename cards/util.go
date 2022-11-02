package cards

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func dotaccess(val interface{}, data []byte, dotnote string) error {
	keys := strings.Split(dotnote, ".")
	var raw json.RawMessage = data
	for _, key := range keys {
		var mapraw map[string]*json.RawMessage
		err := json.Unmarshal(raw, &mapraw)
		if err != nil {
			return err
		}
		if r, ok := mapraw[key]; ok {
			raw = *r
		} else {
			return fmt.Errorf("%v key does not exist in json", key)
		}
	}
	return json.Unmarshal(raw, val)
}

func JSONMarshal(i interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(i)
	return buffer.Bytes(), err
}

func stringOr(a, b string) string {
	if len(a) > 0 {
		return a
	}
	return b

}
