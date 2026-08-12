package models

import (
	"encoding/json"
	"testing"
)

func TestUpdateTagMarshalsFalseBooleans(t *testing.T) {
	falseVal := false
	trueVal := true

	payload, err := json.Marshal(&UpdateTag{
		TagKey:    "environment",
		Preferred: &falseVal,
		Hidden:    &trueVal,
	})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if decoded["preferred"] != false {
		t.Fatalf("expected preferred=false, got %#v", decoded["preferred"])
	}
	if decoded["hidden"] != true {
		t.Fatalf("expected hidden=true, got %#v", decoded["hidden"])
	}
}

func TestUpdateTagOmitsNilBooleans(t *testing.T) {
	payload, err := json.Marshal(&UpdateTag{
		TagKey:    "environment",
		Preferred: nil,
		Hidden:    nil,
	})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}

	if _, ok := decoded["preferred"]; ok {
		t.Fatalf("expected preferred to be omitted, got %#v", decoded["preferred"])
	}
	if _, ok := decoded["hidden"]; ok {
		t.Fatalf("expected hidden to be omitted, got %#v", decoded["hidden"])
	}
}
