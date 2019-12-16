package problems

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPower(t *testing.T) {
	tests := []struct {
		data float64
		e    int
		ret  float64
	}{
		{-123.123, 0, 1},
		{0, 123, 0},
		{0.2, 2, 0.04},
	}

	for _, tt := range tests {
		assert.Equal(t, 0, int(math.Abs(Power(tt.data, tt.e)-tt.ret)*1000000))
	}
}
