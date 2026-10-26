package models

import (
	"encoding/json"
	"testing"
)

func TestUpdateBusinessMetricReportAttachmentsOmission(t *testing.T) {
	token := "rprt_1"
	tests := []struct {
		name    string
		reports []*UpdateBusinessMetricCostReportTokensWithMetadataItems0
		want    string
	}{
		{name: "omitted", want: `{"title":"Metric"}`},
		{name: "clear", reports: []*UpdateBusinessMetricCostReportTokensWithMetadataItems0{}, want: `{"title":"Metric","cost_report_tokens_with_metadata":[]}`},
		{name: "replace", reports: []*UpdateBusinessMetricCostReportTokensWithMetadataItems0{{CostReportToken: &token, LabelFilter: []string{}}}, want: `{"title":"Metric","cost_report_tokens_with_metadata":[{"cost_report_token":"rprt_1","label_filter":[]}]}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(&UpdateBusinessMetric{Title: "Metric", CostReportTokensWithMetadata: tt.reports})
			if err != nil {
				t.Fatal(err)
			}
			if string(body) != tt.want {
				t.Fatalf("got %s, want %s", body, tt.want)
			}
		})
	}
}
