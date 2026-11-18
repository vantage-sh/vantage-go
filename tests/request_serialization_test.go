// These contracts exercise omission and clearing semantics in generated request models.
package contract_test

import (
	"encoding/json"
	"testing"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

func TestRecommendationViewFilterSerialization(t *testing.T) {
	fields := []struct {
		name    string
		request func(*[]string) *models.UpdateRecommendationView
	}{
		{"provider_ids", func(value *[]string) *models.UpdateRecommendationView {
			return &models.UpdateRecommendationView{ProviderIds: value}
		}},
		{"billing_account_ids", func(value *[]string) *models.UpdateRecommendationView {
			return &models.UpdateRecommendationView{BillingAccountIds: value}
		}},
		{"account_ids", func(value *[]string) *models.UpdateRecommendationView {
			return &models.UpdateRecommendationView{AccountIds: value}
		}},
		{"regions", func(value *[]string) *models.UpdateRecommendationView {
			return &models.UpdateRecommendationView{Regions: value}
		}},
		{"types", func(value *[]string) *models.UpdateRecommendationView {
			return &models.UpdateRecommendationView{Types: value}
		}},
	}

	for _, field := range fields {
		t.Run(field.name, func(t *testing.T) {
			for _, test := range optionalStringLists() {
				t.Run(test.name, func(t *testing.T) {
					assertJSONField(t, field.request(test.value), test.want, field.name)
				})
			}
		})
	}
}

func TestCostReportTableColumnOrderSerialization(t *testing.T) {
	for _, test := range optionalStringLists() {
		t.Run(test.name, func(t *testing.T) {
			request := &models.UpdateCostReport{
				Settings: &models.UpdateCostReportSettings{TableColumnOrder: test.value},
			}
			assertJSONField(t, request, test.want, "settings", "table_column_order")
		})
	}
}

func TestCostReportCompletePeriodSerialization(t *testing.T) {
	completePeriod := false
	request := &models.UpdateCostReport{
		Settings: &models.UpdateCostReportSettings{CompletePeriod: &completePeriod},
	}
	assertJSONField(t, request, "false", "settings", "complete_period")
	request.Settings.CompletePeriod = nil
	assertJSONField(t, request, "", "settings", "complete_period")
}

func TestCalculationAttachmentSerialization(t *testing.T) {
	requests := []struct {
		name    string
		field   string
		request func([]string) any
	}{
		{"create_calculation", "cost_report_tokens", func(value []string) any {
			return &models.CreateCustomCalculation{CostReportTokens: value}
		}},
		{"update_calculation", "cost_report_tokens", func(value []string) any {
			return &models.UpdateCustomCalculation{CostReportTokens: value}
		}},
		{"create_cost_report", "custom_calculation_tokens", func(value []string) any {
			return &models.CreateCostReport{CustomCalculationTokens: value}
		}},
		{"update_cost_report", "custom_calculation_tokens", func(value []string) any {
			return &models.UpdateCostReport{CustomCalculationTokens: value}
		}},
	}
	tests := []struct {
		name  string
		value []string
		want  string
	}{
		{"unchanged", nil, "null"},
		{"clear", []string{}, "[]"},
		{"replace", []string{"first", "second"}, `["first","second"]`},
	}

	for _, request := range requests {
		t.Run(request.name, func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					assertJSONField(t, request.request(test.value), test.want, request.field)
				})
			}
		})
	}
}

func TestCalculationPartialUpdateDoesNotSendDefaults(t *testing.T) {
	request := &models.UpdateCustomCalculation{Title: "Renamed"}
	assertJSONField(t, request, `"Renamed"`, "title")
	for _, field := range []string{"operation", "display_as", "unit_scale", "description", "filter"} {
		assertJSONField(t, request, "", field)
	}
}

func TestCalculationUsageUnitTypes(t *testing.T) {
	fields := []struct {
		name  string
		write func(any) *models.CustomCalculationFilter
		read  func(*models.CustomCalculationFilter) any
	}{
		{"usage_unit", func(value any) *models.CustomCalculationFilter {
			return &models.CustomCalculationFilter{UsageUnit: value}
		}, func(filter *models.CustomCalculationFilter) any { return filter.UsageUnit }},
		{"usage_unit_numerator", func(value any) *models.CustomCalculationFilter {
			return &models.CustomCalculationFilter{UsageUnitNumerator: value}
		}, func(filter *models.CustomCalculationFilter) any { return filter.UsageUnitNumerator }},
		{"usage_unit_denominator", func(value any) *models.CustomCalculationFilter {
			return &models.CustomCalculationFilter{UsageUnitDenominator: value}
		}, func(filter *models.CustomCalculationFilter) any { return filter.UsageUnitDenominator }},
	}

	for _, field := range fields {
		for _, value := range []string{`"GB"`, `["GB","GiB"]`, `null`} {
			t.Run(field.name+"/"+value, func(t *testing.T) {
				request := &models.UpdateCustomCalculation{Filter: field.write(json.RawMessage(value))}
				assertJSONField(t, request, value, "filter", field.name)

				var response models.CustomCalculation
				body := `{"filter":{"` + field.name + `":` + value + `}}`
				if err := json.Unmarshal([]byte(body), &response); err != nil {
					t.Fatalf("decode calculation response: %v", err)
				}
				if response.Filter == nil {
					t.Fatal("missing structured calculation filter")
				}
				decoded, err := json.Marshal(field.read(response.Filter))
				if err != nil || string(decoded) != value {
					t.Fatalf("decoded %s = %s, want %s (error: %v)", field.name, decoded, value, err)
				}
			})
		}
	}
}

func optionalStringLists() []struct {
	name  string
	value *[]string
	want  string
} {
	empty := []string{}
	populated := []string{"first", "second"}
	var null []string

	return []struct {
		name  string
		value *[]string
		want  string
	}{
		{"omitted", nil, ""},
		{"empty", &empty, "[]"},
		{"populated", &populated, `["first","second"]`},
		{"null", &null, "null"},
	}
}

func assertJSONField(t *testing.T, request any, want string, path ...string) {
	t.Helper()
	encoded, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}

	for i, field := range path {
		var fields map[string]json.RawMessage
		if err := json.Unmarshal(encoded, &fields); err != nil {
			t.Fatalf("decode request at %v: %v", path[:i], err)
		}
		value, present := fields[field]
		if i == len(path)-1 {
			if want == "" {
				if present {
					t.Fatalf("%v must be omitted; got %s", path, value)
				}
				return
			}
			if !present || string(value) != want {
				t.Fatalf("%v = %s (present: %v), want %s", path, value, present, want)
			}
			return
		}
		if !present {
			t.Fatalf("missing request object %v", path[:i+1])
		}
		encoded = value
	}
}
