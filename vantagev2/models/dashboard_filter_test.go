package models

import (
	"encoding/json"
	"testing"
)

func TestDashboardFilterJSON(t *testing.T) {
	filter := "vantage.provider = 'aws'"
	for _, tc := range []struct {
		name    string
		body    interface{}
		want    interface{}
		present bool
	}{
		{"create omitted", &CreateDashboard{}, nil, false},
		{"update omitted", &UpdateDashboard{}, nil, false},
		{"create set", &CreateDashboard{Filter: &filter}, filter, true},
		{"update set", &UpdateDashboard{Filter: &filter}, filter, true},
		{"read nullable", &Dashboard{}, nil, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			encoded, err := json.Marshal(tc.body)
			if err != nil {
				t.Fatal(err)
			}
			var body map[string]interface{}
			if err := json.Unmarshal(encoded, &body); err != nil {
				t.Fatal(err)
			}
			value, present := body["filter"]
			if present != tc.present || value != tc.want {
				t.Fatalf("filter=%#v present=%t", value, present)
			}
		})
	}
}
