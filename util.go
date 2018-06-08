package main

import "encoding/json"

func marshal(o interface{}) string {
	var bytes, err = json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

