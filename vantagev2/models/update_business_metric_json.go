package models

import "encoding/json"

// MarshalJSON preserves an explicit empty report list so it clears attachments,
// while an omitted list leaves existing attachments unchanged.
func (m *UpdateBusinessMetric) MarshalJSON() ([]byte, error) {
	type model UpdateBusinessMetric
	var reports *[]*UpdateBusinessMetricCostReportTokensWithMetadataItems0
	if m.CostReportTokensWithMetadata != nil {
		reports = &m.CostReportTokensWithMetadata
	}

	return json.Marshal(struct {
		*model
		CostReportTokensWithMetadata *[]*UpdateBusinessMetricCostReportTokensWithMetadataItems0 `json:"cost_report_tokens_with_metadata,omitempty"`
	}{(*model)(m), reports})
}
