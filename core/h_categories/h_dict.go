package h_categories

import (
	"encoding/json"
)

type Dict map[string]interface{}

func ByteToDict(b []byte) (*Dict, error) {
	var (
		returnDict *Dict
		err        error
	)
	json.Unmarshal(b, returnDict)
	return returnDict, err
}
