package fastjson

import (
	"io"

	"github.com/goccy/go-json"
	"github.com/pucora/lura/v2/encoding"
)

const Name = "fast-json"

func Register() error {
	return encoding.GetRegister().Register(Name, NewDecoder)
}

func NewDecoder(isCollection bool) func(io.Reader, *map[string]interface{}) error {
	if isCollection {
		return CollectionDecoder
	}
	return Decoder
}

func Decoder(r io.Reader, v *map[string]interface{}) error {
	d := json.NewDecoder(r)
	d.UseNumber()
	return d.Decode(v)
}

func CollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	var collection []interface{}
	d := json.NewDecoder(r)
	d.UseNumber()
	if err := d.Decode(&collection); err != nil {
		return err
	}
	*v = map[string]interface{}{"collection": collection}
	return nil
}
