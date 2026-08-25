package models

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
)

func TestCreateBudgetPeriodCadenceJSONAndValidation(t *testing.T) {
	start := strfmt.Date(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC))
	name := "Engineering"
	budget := &CreateBudget{
		Name: &name,
		PeriodCadence: &CreateBudgetPeriodCadence{
			StartsAt:      &start,
			IntervalCount: 2,
			IntervalUnit:  CreateBudgetPeriodCadenceIntervalUnitWeek,
		},
	}

	if err := budget.Validate(strfmt.Default); err != nil {
		t.Fatalf("expected budget to validate: %v", err)
	}

	body, err := json.Marshal(budget)
	if err != nil {
		t.Fatalf("expected budget to marshal: %v", err)
	}
	if !strings.Contains(string(body), `"period_cadence":{"interval_count":2,"interval_unit":"week","starts_at":"2026-01-01"}`) {
		t.Fatalf("expected period_cadence in JSON, got %s", body)
	}
}

func TestCreateBudgetPeriodCadenceValidatesIntervalUnit(t *testing.T) {
	name := "Engineering"
	budget := &CreateBudget{
		Name: &name,
		PeriodCadence: &CreateBudgetPeriodCadence{
			IntervalUnit: "quarter",
		},
	}

	if err := budget.Validate(strfmt.Default); err == nil {
		t.Fatal("expected invalid interval_unit to fail validation")
	}
}

func TestBudgetRequiresPeriodCadence(t *testing.T) {
	name := "Engineering"
	budget := &Budget{
		BudgetAlertTokens: []string{},
		ChildBudgetTokens: []string{},
		CreatedAt:         "2026-01-01T00:00:00Z",
		Name:              &name,
		Periods:           []*BudgetPeriod{},
		Token:             "bdgt_123",
		WorkspaceToken:    "wrkspc_123",
	}

	if err := budget.Validate(strfmt.Default); err == nil {
		t.Fatal("expected missing period_cadence to fail validation")
	}
}

func TestBudgetAllowsNullPeriodCadenceStartsAt(t *testing.T) {
	name := "Engineering"
	budget := &Budget{
		BudgetAlertTokens: []string{},
		ChildBudgetTokens: []string{},
		CreatedAt:         "2026-01-01T00:00:00Z",
		Name:              &name,
		PeriodCadence: &PeriodCadence{
			IntervalCount: 1,
			IntervalUnit:  "month",
		},
		Periods:        []*BudgetPeriod{},
		Token:          "bdgt_123",
		WorkspaceToken: "wrkspc_123",
	}

	if err := budget.Validate(strfmt.Default); err != nil {
		t.Fatalf("expected null starts_at to validate: %v", err)
	}

	body, err := json.Marshal(budget)
	if err != nil {
		t.Fatalf("expected budget to marshal: %v", err)
	}
	if !strings.Contains(string(body), `"starts_at":null`) {
		t.Fatalf("expected starts_at to marshal as null, got %s", body)
	}
}

func TestBudgetPeriodsMarshalNullEndAt(t *testing.T) {
	start := strfmt.Date(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC))
	amount := 123.45

	testCases := []struct {
		name  string
		value interface{}
	}{
		{
			name: "create budget period",
			value: &CreateBudgetPeriodsItems0{
				Amount:  &amount,
				StartAt: &start,
			},
		},
		{
			name: "update budget period",
			value: &UpdateBudgetPeriodsItems0{
				Amount:  &amount,
				StartAt: &start,
			},
		},
	}

	for _, tc := range testCases {
		body, err := json.Marshal(tc.value)
		if err != nil {
			t.Fatalf("expected %s to marshal: %v", tc.name, err)
		}
		if !strings.Contains(string(body), `"end_at":null`) {
			t.Fatalf("expected %s end_at to marshal as null, got %s", tc.name, body)
		}
	}
}
