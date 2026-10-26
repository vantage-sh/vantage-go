package contract_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	models "github.com/vantage-sh/vantage-go/vantagev2/models"
	metrics "github.com/vantage-sh/vantage-go/vantagev2/vantage/business_metrics"
)

func TestGeneratedMetronomeClient(t *testing.T) {
	for _, kind := range []string{"revenue", "usage"} {
		t.Run(kind, func(t *testing.T) {
			fields := map[string]any{"metric_kind": kind, "metric": "daily_revenue", "label_source": "customer_id", "billable_metric_id": nil, "billable_metric_name": nil}
			if kind == "usage" {
				fields["billable_metric_id"] = "c46e063a-284b-4d91-8828-eac166092c7a"
				fields["billable_metric_name"] = "Tokens"
			}
			calls := []string{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls = append(calls, r.Method)
				if r.Method != http.MethodGet {
					var body map[string]any
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Error(err)
					}
					config, ok := body["metronome_metric_fields"].(map[string]any)
					if !ok {
						t.Errorf("missing config: %#v", body)
					}
					if r.Method == http.MethodPost && config["metric_kind"] != kind {
						t.Errorf("create kind: %#v", config)
					}
					if r.Method == http.MethodPut {
						if !reflect.DeepEqual(config, map[string]any{"label_source": ""}) {
							t.Errorf("partial update = %#v", config)
						}
						if !reflect.DeepEqual(body["cost_report_tokens_with_metadata"], []any{}) {
							t.Errorf("clear links = %#v", body)
						}
					}
				}
				w.Header().Set("Content-Type", "application/json")
				if r.Method == http.MethodPost {
					w.WriteHeader(http.StatusCreated)
				}
				_ = json.NewEncoder(w).Encode(map[string]any{"token": "bsnss_mtrc_test", "title": "Metronome", "import_type": "metronome_metrics", "integration_token": "accss_crdntl_test", "metronome_metric_fields": fields, "cost_report_tokens_with_metadata": []any{}})
			}))
			defer server.Close()
			u, _ := url.Parse(server.URL)
			client := metrics.New(httptransport.New(u.Host, "/v2", []string{"http"}), strfmt.Default)
			create := &models.CreateBusinessMetric{Title: swag.String("Metronome"), MetronomeMetricFields: &models.CreateBusinessMetricMetronomeMetricFields{IntegrationToken: swag.String("accss_crdntl_test"), MetricKind: swag.String(kind)}}
			if kind == "usage" {
				create.MetronomeMetricFields.BillableMetricID = swag.String(fields["billable_metric_id"].(string))
				create.MetronomeMetricFields.BillableMetricName = swag.String("Tokens")
			}
			if err := create.Validate(strfmt.Default); err != nil {
				t.Fatal(err)
			}
			out, err := client.CreateBusinessMetric(metrics.NewCreateBusinessMetricParams().WithCreateBusinessMetric(create), nil)
			if err != nil {
				t.Fatal(err)
			}
			if err := out.Payload.MetronomeMetricFields.Validate(strfmt.Default); err != nil {
				t.Fatal(err)
			}
			read, err := client.GetBusinessMetric(metrics.NewGetBusinessMetricParams().WithBusinessMetricToken(out.Payload.Token), nil)
			if err != nil {
				t.Fatal(err)
			}
			if read.Payload.MetronomeMetricFields.MetricKind != kind {
				t.Fatalf("read config = %#v", read.Payload.MetronomeMetricFields)
			}
			_, err = client.UpdateBusinessMetric(metrics.NewUpdateBusinessMetricParams().WithBusinessMetricToken(out.Payload.Token).WithUpdateBusinessMetric(&models.UpdateBusinessMetric{MetronomeMetricFields: &models.UpdateBusinessMetricMetronomeMetricFields{LabelSource: swag.String("")}, CostReportTokensWithMetadata: []*models.UpdateBusinessMetricCostReportTokensWithMetadataItems0{}}), nil)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(calls, []string{"POST", "GET", "PUT"}) {
				t.Fatal(calls)
			}
		})
	}
}

func TestGeneratedOmissionAndClearing(t *testing.T) {
	for _, tc := range []struct {
		name    string
		metric  models.UpdateBusinessMetric
		reports string
	}{
		{name: "nil reports", reports: "null"},
		{name: "empty reports", metric: models.UpdateBusinessMetric{CostReportTokensWithMetadata: []*models.UpdateBusinessMetricCostReportTokensWithMetadataItems0{}}, reports: "[]"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b, err := json.Marshal(tc.metric)
			if err != nil {
				t.Fatal(err)
			}
			var got map[string]json.RawMessage
			if err = json.Unmarshal(b, &got); err != nil {
				t.Fatal(err)
			}
			if _, ok := got["metronome_metric_fields"]; ok {
				t.Fatalf("omitted configuration serialized: %s", b)
			}
			if string(got["cost_report_tokens_with_metadata"]) != tc.reports {
				t.Fatal(string(b))
			}
		})
	}
}
