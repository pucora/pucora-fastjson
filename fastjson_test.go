package fastjson_test

import (
	"bytes"
	"testing"

	fastjson "github.com/pucora/pucora-fastjson"
)

func TestDecoder(t *testing.T) {
	input := `{"user": {"id": 123, "name": "test"}, "active": true}`
	buf := bytes.NewBufferString(input)

	var v map[string]interface{}
	err := fastjson.Decoder(buf, &v)
	if err != nil {
		t.Fatal(err)
	}

	// We'll just verify some fields to avoid json.Number deep equality headaches
	if v["active"] != true {
		t.Errorf("expected active=true")
	}
	userMap, ok := v["user"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected user to be map")
	}
	if userMap["name"] != "test" {
		t.Errorf("expected name=test")
	}
}

func TestCollectionDecoder(t *testing.T) {
	input := `[{"id": 1}, {"id": 2}]`
	buf := bytes.NewBufferString(input)

	var v map[string]interface{}
	err := fastjson.CollectionDecoder(buf, &v)
	if err != nil {
		t.Fatal(err)
	}

	collection, ok := v["collection"].([]interface{})
	if !ok {
		t.Fatalf("expected collection to be slice")
	}
	if len(collection) != 2 {
		t.Errorf("expected collection length 2, got %d", len(collection))
	}
}
