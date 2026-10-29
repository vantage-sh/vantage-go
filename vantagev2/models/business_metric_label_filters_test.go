package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBusinessMetricCanonicalLabelFilterRequests(t *testing.T) {
	token := "rprt_test"
	filters := map[string][]string{"label": {"production", ""}}
	for _, legacy := range [][]string{nil, {}} {
		requests := map[string]any{
			"create": &CreateBusinessMetric{CostReportTokensWithMetadata: []*CreateBusinessMetricCostReportTokensWithMetadataItems0{{CostReportToken: &token, LabelFilter: legacy, LabelFilters: filters}}},
			"update": &UpdateBusinessMetric{CostReportTokensWithMetadata: []*UpdateBusinessMetricCostReportTokensWithMetadataItems0{{CostReportToken: &token, LabelFilter: legacy, LabelFilters: filters}}},
		}
		for name, request := range requests {
			t.Run(name, func(t *testing.T) {
				wire, err := json.Marshal(request)
				if err != nil {
					t.Fatal(err)
				}
				var decoded struct {
					Attachments []map[string]json.RawMessage `json:"cost_report_tokens_with_metadata"`
				}
				if err := json.Unmarshal(wire, &decoded); err != nil {
					t.Fatal(err)
				}
				attachment := decoded.Attachments[0]
				if _, exists := attachment["label_filter"]; exists {
					t.Fatalf("unexpected singular alias in canonical request: %s", wire)
				}
				var actual map[string][]string
				if err := json.Unmarshal(attachment["label_filters"], &actual); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(actual, filters) {
					t.Fatalf("label_filters = %#v, want %#v", actual, filters)
				}
			})
		}
	}
}

func TestBusinessMetricLegacyLabelFilterRequest(t *testing.T) {
	token := "rprt_test"
	for _, request := range []any{
		&CreateBusinessMetricCostReportTokensWithMetadataItems0{CostReportToken: &token, LabelFilter: []string{"production", ""}},
		&UpdateBusinessMetricCostReportTokensWithMetadataItems0{CostReportToken: &token, LabelFilter: []string{"production", ""}},
	} {
		wire, err := json.Marshal(request)
		if err != nil {
			t.Fatal(err)
		}
		var decoded map[string]json.RawMessage
		if err := json.Unmarshal(wire, &decoded); err != nil {
			t.Fatal(err)
		}
		if _, exists := decoded["label_filters"]; exists {
			t.Fatalf("unexpected canonical map in singular request: %s", wire)
		}
		var labels []string
		if err := json.Unmarshal(decoded["label_filter"], &labels); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(labels, []string{"production", ""}) {
			t.Fatalf("label_filter = %#v", labels)
		}
	}
}
