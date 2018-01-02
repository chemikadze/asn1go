package asn1go

import (
	"encoding/asn1"
	"io/ioutil"
	"os"
)

func MarshalToFile(val interface{}, path string, mode os.FileMode) error {
	data, err := asn1.Marshal(val)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, data, mode)
	if err != nil {
		return err
	}
	return nil
}
