package api_test

import (
	"context"
	"testing"

	"rilutham/stovia/lib/api"
	pb "rilutham/stovia/lib/rpc"
)

func TestGetFutureValue(t *testing.T) {
	t.Run("get future value of some stock", func(t *testing.T) {
		api := api.NewAPI()
		cases := []struct {
			code, year, quarter string
			targetYear          int32
			futureValue         float64
			recommendedToBuy    bool
		}{
			{"BBCA", "2018", "4", 2024, 1000, true},
		}

		for _, c := range cases {
			req := &pb.FutureValueRequest{
				Code:       c.code,
				Year:       c.year,
				Quarter:    c.quarter,
				TargetYear: c.targetYear,
			}
			resp, err := api.GetFutureValue(context.Background(), req)
			if err != nil {
				t.Errorf("You got unexpected error")
			}
			if resp.FutureValue != c.futureValue {
				t.Errorf("Future value = %v, wanted %v.", resp.FutureValue, c.futureValue)
			}
		}
	})
}
