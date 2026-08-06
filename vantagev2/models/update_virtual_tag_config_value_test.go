package models

import (
	"encoding/json"
	"testing"
)

func TestUpdateVirtualTagConfigValueMarshalsOnlySetFields(t *testing.T) {
	payload, err := json.Marshal(&UpdateVirtualTagConfigValue{Name: "platform"})
	if err != nil {
		t.Fatal(err)
	}

	if string(payload) != `{"name":"platform"}` {
		t.Fatalf("expected only the set field, got %s", payload)
	}
}
