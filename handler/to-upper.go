package main

import (
	"encoding/json"
	"strings"
)

func jsonToUpper(s string) string {

	var f interface{}

	if err := json.Unmarshal([]byte(s), &f); err != nil {
		panic(err)
	}

	if b, err := json.Marshal(f); err != nil {
		panic(err)
	} else {
		return strings.ToUpper(string(b))
	}

}
