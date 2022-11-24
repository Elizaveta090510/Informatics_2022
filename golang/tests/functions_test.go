package internal_test

import (
	"fmt"
	"testing"

	"isuct.ru/informatics2022/internal"
)

func TestEvaluateFunctionTableDriven(t *testing.T) {
	var tests = []struct {
		a, b, x float64
		want    string
	}{
		{0.4, 0.8, 3.2, "0.991503"},
		{0.4, 0.8, 3.8, "0.903289"},
		{0.4, 0.8, 4.4, "0.810890"},
		{0.4, 0.8, 5.0, "0.721277"},
		{0.4, 0.8, 5.6, "0.637818"},
		{0.4, 0.8, 6.2, "0.561889"},
		{0.4, 0.8, 4.48, "0.798680"},
		{0.4, 0.8, 3.56, "0.939642"},
		{0.4, 0.8, 2.78, "1.043989"},
		{0.4, 0.8, 5.28, "0.681446"},
		{0.4, 0.8, 3.21, "0.990128"},
	}

	for _, testInput := range tests {
		result := fmt.Sprintf("%.6f", internal.EvaluateFunction(testInput.a, testInput.b, testInput.x))
		if result != testInput.want {
			t.Errorf("got %s, want %s", result, testInput.want)
		}
	}
}
