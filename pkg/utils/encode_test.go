package utils

import (
	"testing"
)

type Sample struct {
	ID    string `json:"id"`
	Color string `json:"color"`
}

func TestCborEncode(t *testing.T) {
	s := Sample{
		ID:    "test",
		Color: "red",
	}

	b := make([]byte, 64)
	e := CborEncoder(&b)
	err := e.Encode(&s)
	if err != nil {
		t.Fatal(err)
	}

	d := CborDecoder(&b)
	var r Sample
	err = d.Decode(&r)
	if err != nil {
		t.Fatal(err)
	}

	if r.ID != s.ID {
		t.FailNow()
	}

	if r.Color != s.Color {
		t.FailNow()
	}
}
