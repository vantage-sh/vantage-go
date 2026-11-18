package vantagev2_test

import (
	"encoding/json"
	"testing"

	"github.com/vantage-sh/vantage-go/vantagev2/models"
)

func TestCostReportColumnOrderSerialization(t *testing.T) {
	for _, operation := range []struct {
		name  string
		model func() interface{}
	}{
		{"create", func() interface{} { return &models.CreateCostReport{} }},
		{"update", func() interface{} { return &models.UpdateCostReport{} }},
	} {
		for _, tc := range []struct {
			name     string
			settings string
			want     string
		}{
			{"omitted", `{}`, ""},
			{"reset", `{"table_column_order":[]}`, `[]`},
			{"ordered", `{"table_column_order":["service","provider","costAmount"]}`, `["service","provider","costAmount"]`},
			{"reordered", `{"table_column_order":["costAmount","provider","service"]}`, `["costAmount","provider","service"]`},
		} {
			t.Run(operation.name+"/"+tc.name, func(t *testing.T) {
				model := operation.model()
				if err := json.Unmarshal([]byte(`{"settings":`+tc.settings+`}`), model); err != nil {
					t.Fatal(err)
				}
				encoded, err := json.Marshal(model)
				if err != nil {
					t.Fatal(err)
				}
				var body struct {
					Settings map[string]json.RawMessage `json:"settings"`
				}
				if err := json.Unmarshal(encoded, &body); err != nil {
					t.Fatal(err)
				}
				if got := string(body.Settings["table_column_order"]); got != tc.want {
					t.Fatalf("table_column_order = %q, want %q", got, tc.want)
				}
			})
		}
	}
}
