package models

import (
	"encoding/json"
	"testing"
)

func TestUpdateReportForecastMarshalsFalseSetAsDefault(t *testing.T) {
	value := false
	payload, err := json.Marshal(&UpdateReportForecast{SetAsDefault: &value})
	if err != nil {
		t.Fatal(err)
	}

	if string(payload) != `{"set_as_default":false}` {
		t.Fatalf("expected explicit false set_as_default, got %s", payload)
	}
}

func TestUpdateReportForecastOmitsUnsetSetAsDefault(t *testing.T) {
	payload, err := json.Marshal(&UpdateReportForecast{})
	if err != nil {
		t.Fatal(err)
	}

	if string(payload) != `{}` {
		t.Fatalf("expected unset set_as_default to be omitted, got %s", payload)
	}
}
