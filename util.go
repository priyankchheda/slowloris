package main

import (
	"io"
	"io/ioutil"
)

// readResponse reads twitter raw json response
func readResponse(data io.Reader) ([]byte, error) {
	bits, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}

	err = CheckForResponseError(bits)
	if err != nil {
		return nil, err
	}

	return bits, nil
}
