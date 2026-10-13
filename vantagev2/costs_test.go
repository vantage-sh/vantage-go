package vantagev2_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	costsv2 "github.com/vantage-sh/vantage-go/vantagev2/vantage/costs"
)

func TestGetCostsGroupingSortQuery(t *testing.T) {
	for _, tc := range []struct {
		name      string
		column    string
		direction string
	}{
		{name: "omitted"},
		{name: "ascending", column: "service", direction: "asc"},
		{name: "descending", column: "service", direction: "desc"},
		{name: "column with default direction", column: "tag:owner's team"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			requests := make(chan url.Values, 1)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet || r.URL.Path != "/v2/costs" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				requests <- r.URL.Query()
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `{"costs":[],"total_cost":{"amount":"0","currency":"USD"},"links":{}}`)
			}))
			defer server.Close()

			host, err := url.Parse(server.URL)
			if err != nil {
				t.Fatal(err)
			}
			client := costsv2.NewClientWithBearerToken(host.Host, "/v2", "http", "test-token")
			token := "rprt_123"
			params := costsv2.NewGetCostsParams().WithCostReportToken(&token)
			expected := url.Values{"cost_report_token": {token}}
			if tc.column != "" {
				params.WithSortColumn(&tc.column)
				expected.Set("sort_column", tc.column)
			}
			if tc.direction != "" {
				params.WithSortDirection(&tc.direction)
				expected.Set("sort_direction", tc.direction)
			}

			if _, err := client.GetCosts(params, nil); err != nil {
				t.Fatal(err)
			}
			if got := <-requests; !reflect.DeepEqual(expected, got) {
				t.Fatalf("expected query %v, got %v", expected, got)
			}
		})
	}
}
