package nhlApi

import (
	"fmt"
	"testing"

	mock "github.com/jarcoal/httpmock"
)

func TestGetAllTeams(t *testing.T) {
	mock.Activate()

	mock.RegisterResponder("GET", BaseURL+"/teams", GetAllTeamsMock)

	var tests = []struct {
		want int
	}{
		{1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("table-style")
		t.Run(testname, func(t *testing.T) {
			ans, _ := GetAllTeams()

			if ans[0].ID != tt.want {
				t.Errorf("got: %d; want %d", ans[0].ID, tt.want)
			}

		})
	}

	defer mock.DeactivateAndReset()
}
