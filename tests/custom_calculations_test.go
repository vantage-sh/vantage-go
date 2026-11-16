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
	"github.com/vantage-sh/vantage-go/vantagev2/models"
	vantage "github.com/vantage-sh/vantage-go/vantagev2/vantage"
	reports "github.com/vantage-sh/vantage-go/vantagev2/vantage/costs"
	calculations "github.com/vantage-sh/vantage-go/vantagev2/vantage/custom_calculations"
)

func TestCalculationUsageUnits(t *testing.T) {
	if vantage.DefaultBasePath != "/v2" {
		t.Fatalf("default base path = %q", vantage.DefaultBasePath)
	}
	for _, unit := range []any{"Hours", []any{"Hours", "Hrs"}, nil} {
		filter := map[string]any{"usage_unit": unit, "usage_unit_numerator": unit, "usage_unit_denominator": unit}
		for _, model := range []any{
			&models.CreateCustomCalculation{Title: swag.String("Usage"), Filter: filter},
			&models.UpdateCustomCalculation{Filter: filter},
			&models.CustomCalculation{Filter: filter},
		} {
			data, err := json.Marshal(model)
			if err != nil {
				t.Fatal(err)
			}
			var got struct {
				Filter map[string]any `json:"filter"`
			}
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got.Filter, filter) {
				t.Fatalf("%T filter = %#v", model, got.Filter)
			}
		}
	}
	if err := (&models.CreateCustomCalculation{}).Validate(strfmt.Default); err == nil {
		t.Fatal("missing required title/filter accepted")
	}
	if err := (&models.UpdateCustomCalculation{Operation: "invalid"}).Validate(strfmt.Default); err == nil {
		t.Fatal("invalid operation accepted")
	}
}

// These assertions record the generated representation, including its limits:
// nil description cannot express explicit null; nil attachments emit null.
func TestCalculationDescriptionSerialization(t *testing.T) {
	for _, tc := range []struct{ input, want string }{
		{`{}`, ""}, {`{"description":null}`, ""}, {`{"description":""}`, `""`}, {`{"description":"Details"}`, `"Details"`},
	} {
		for _, model := range []any{&models.CreateCustomCalculation{}, &models.UpdateCustomCalculation{}} {
			if err := json.Unmarshal([]byte(tc.input), model); err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(model)
			if err != nil {
				t.Fatal(err)
			}
			var got map[string]json.RawMessage
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatal(err)
			}
			if string(got["description"]) != tc.want {
				t.Fatalf("%T %s => %s", model, tc.input, data)
			}
		}
	}
	for _, input := range []string{`null`, `""`, `"Details"`} {
		var response models.CustomCalculation
		if err := json.Unmarshal([]byte(`{"description":`+input+`}`), &response); err != nil {
			t.Fatal(err)
		}
		data, err := json.Marshal(response)
		if err != nil {
			t.Fatal(err)
		}
		var got map[string]json.RawMessage
		if err := json.Unmarshal(data, &got); err != nil {
			t.Fatal(err)
		}
		if string(got["description"]) != input {
			t.Fatalf("description %s => %s", input, data)
		}
	}
}

func TestCalculationHTTPContracts(t *testing.T) {
	filter := map[string]any{"usage_unit": []any{"Hours", "Hrs"}, "usage_unit_numerator": []any{"Tokens", "tokens"}, "usage_unit_denominator": nil}
	response := &models.CustomCalculation{Token: "cstm_clcltn_test", Title: "Usage", Description: swag.String("Details"), Filter: filter, Operation: "sum", DisplayAs: "ratio", UnitScale: 1, IsReadOnly: true, CostReportTokens: []string{"rprt_test"}, CreatedAt: "2026-11-16T00:00:00Z", UpdatedAt: "2026-11-16T00:00:00Z"}
	calls := []string{}
	var wantBody map[string]any
	status := http.StatusOK
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls = append(calls, r.Method+" "+r.URL.RequestURI())
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("authorization missing: %v", r.Header)
		}
		if r.Method == http.MethodPost || r.Method == http.MethodPut {
			var got map[string]any
			if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, wantBody) {
				t.Errorf("body = %#v; want %#v", got, wantBody)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if status >= 400 {
			_ = json.NewEncoder(w).Encode(map[string]any{"errors": []string{"denied"}})
			return
		}
		if status == http.StatusNoContent {
			return
		}
		if r.URL.Path == "/v2/custom_calculations" && r.Method == http.MethodGet {
			_ = json.NewEncoder(w).Encode(&models.CustomCalculations{CustomCalculations: []*models.CustomCalculation{response}})
		} else if r.URL.Path == "/v2/cost_reports/rprt_test" {
			_ = json.NewEncoder(w).Encode(map[string]any{"token": "rprt_test"})
		} else {
			_ = json.NewEncoder(w).Encode(response)
		}
	}))
	defer server.Close()
	u, _ := url.Parse(server.URL)
	transport := httptransport.New(u.Host, "/v2", []string{"http"})
	client := calculations.New(transport, strfmt.Default)
	reportClient := reports.New(transport, strfmt.Default)
	auth := httptransport.BearerToken("test-token")
	check := func(got *models.CustomCalculation, err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, response) {
			t.Fatalf("decoded response = %#v", got)
		}
		if err := got.Validate(strfmt.Default); err != nil {
			t.Fatal(err)
		}
		editable := *got
		editable.IsReadOnly = false
		if err := editable.Validate(strfmt.Default); err == nil {
			t.Fatal("required false boolean validator behavior changed; revisit limitation")
		}
		unlinked := *got
		unlinked.CostReportTokens = []string{}
		if err := unlinked.Validate(strfmt.Default); err != nil {
			t.Fatalf("valid empty response attachments: %v", err)
		}
	}
	status = http.StatusCreated
	wantBody = map[string]any{"title": "Usage", "filter": filter, "description": "Details"}
	created, err := client.CreateCustomCalculation(calculations.NewCreateCustomCalculationParams().WithCreateCustomCalculation(&models.CreateCustomCalculation{Title: swag.String("Usage"), Filter: filter, Description: swag.String("Details")}), auth)
	if err != nil {
		t.Fatal(err)
	}
	check(created.Payload, nil)
	status = http.StatusOK
	read, err := client.GetCustomCalculation(calculations.NewGetCustomCalculationParams().WithCustomCalculationToken(response.Token), auth)
	if err != nil {
		t.Fatal(err)
	}
	check(read.Payload, nil)
	list, err := client.GetCustomCalculations(calculations.NewGetCustomCalculationsParams().WithPage(swag.Int32(2)).WithLimit(swag.Int32(10)), auth)
	if err != nil {
		t.Fatal(err)
	}
	if len(list.Payload.CustomCalculations) != 1 {
		t.Fatal(list.Payload)
	}
	check(list.Payload.CustomCalculations[0], nil)
	for _, tc := range []struct {
		name, input string
		want        any
	}{
		{"omitted", ``, nil}, {"null", `null`, nil}, {"empty", `[]`, []any{}}, {"populated", `["rprt_test"]`, []any{"rprt_test"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var update models.UpdateCustomCalculation
			input := `{"description":"","filter":{"usage_unit":["Hours","Hrs"],"usage_unit_numerator":["Tokens","tokens"],"usage_unit_denominator":null}`
			if tc.input != "" {
				input += `,"cost_report_tokens":` + tc.input
			}
			input += `}`
			if err := json.Unmarshal([]byte(input), &update); err != nil {
				t.Fatal(err)
			}
			wantBody = map[string]any{"description": "", "filter": filter, "cost_report_tokens": tc.want}
			updated, err := client.UpdateCustomCalculation(calculations.NewUpdateCustomCalculationParams().WithCustomCalculationToken(response.Token).WithUpdateCustomCalculation(&update), auth)
			if err != nil {
				t.Fatal(err)
			}
			check(updated.Payload, nil)
			var report models.UpdateCostReport
			input = `{}`
			if tc.input != "" {
				input = `{"custom_calculation_tokens":` + tc.input + `}`
			}
			if err := json.Unmarshal([]byte(input), &report); err != nil {
				t.Fatal(err)
			}
			wantBody = map[string]any{"custom_calculation_tokens": tc.want, "business_metric_tokens_with_metadata": nil, "saved_filter_tokens": nil, "groupings": ""}
			_, err = reportClient.UpdateCostReport(reports.NewUpdateCostReportParams().WithCostReportToken("rprt_test").WithUpdateCostReport(&report), auth)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
	status = http.StatusNoContent
	_, err = client.DeleteCustomCalculation(calculations.NewDeleteCustomCalculationParams().WithCustomCalculationToken(response.Token), auth)
	if err != nil {
		t.Fatal(err)
	}
	wantCalls := []string{"POST /v2/custom_calculations", "GET /v2/custom_calculations/cstm_clcltn_test", "GET /v2/custom_calculations?limit=10&page=2"}
	for i := 0; i < 4; i++ {
		wantCalls = append(wantCalls, "PUT /v2/custom_calculations/cstm_clcltn_test", "PUT /v2/cost_reports/rprt_test")
	}
	wantCalls = append(wantCalls, "DELETE /v2/custom_calculations/cstm_clcltn_test")
	if !reflect.DeepEqual(calls, wantCalls) {
		t.Fatalf("calls = %v", calls)
	}
	status = http.StatusForbidden
	_, err = client.GetCustomCalculation(calculations.NewGetCustomCalculationParams().WithCustomCalculationToken(response.Token), auth)
	if _, ok := err.(*calculations.GetCustomCalculationForbidden); !ok {
		t.Fatalf("403 error = %T %v", err, err)
	}
}

// Upstream still marks this array omitempty: an explicit reset is lost.
func TestGeneratedTableColumnResetLimitation(t *testing.T) {
	data, err := json.Marshal(&models.UpdateCostReport{Settings: &models.UpdateCostReportSettings{TableColumnOrder: []string{}}})
	if err != nil {
		t.Fatal(err)
	}
	var got struct {
		Settings map[string]any `json:"settings"`
	}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if _, ok := got.Settings["table_column_order"]; ok {
		t.Fatal("representation changed; revisit upstream reset limitation")
	}
}

func TestCalculationAttachmentResponseSerialization(t *testing.T) {
	for _, tc := range []struct{ input, want string }{{``, `null`}, {`null`, `null`}, {`[]`, `[]`}, {`["token"]`, `["token"]`}} {
		for _, model := range []struct {
			value any
			field string
		}{{&models.CustomCalculation{}, "cost_report_tokens"}, {&models.CostReport{}, "custom_calculation_tokens"}} {
			input := `{}`
			if tc.input != "" {
				input = `{"` + model.field + `":` + tc.input + `}`
			}
			if err := json.Unmarshal([]byte(input), model.value); err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(model.value)
			if err != nil {
				t.Fatal(err)
			}
			var got map[string]json.RawMessage
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatal(err)
			}
			if string(got[model.field]) != tc.want {
				t.Fatalf("%T %s => %s", model.value, input, data)
			}
		}
	}
}
