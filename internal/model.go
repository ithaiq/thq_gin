package internal

import "encoding/json"

type Model interface {
	String() string
}

type Models string

func MakeModels(v interface{}) Models {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return Models(b)
}
