package emitter

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/integration/makefakegenesis"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/inter/pos"
	"github.com/stretchr/testify/assert"
)

func TestDivideValidators(t *testing.T) {
	gValidators := makefakegenesis.GetFakeValidators(10)
	weights := []uint32{25, 24, 20, 15, 5, 4, 3, 2, 1, 1}
	vv := pos.NewBuilder()
	for i, v := range gValidators {
		vv.Set(v.ID, pos.Weight(weights[i]))
	}
	validators := vv.Build()

	testCases := []struct {
		thresholdBasisPoints uint64
		lowGroupSize         idx.Validator
		highGroupSize        idx.Validator
		lowGroupWeight       pos.Weight
		highGroupWeight      pos.Weight
	}{
		{100, 10, 0, 100, 0},
		{50, 10, 0, 100, 0},
		{25, 10, 0, 100, 0},
		{24, 9, 1, 75, 25},
		{20, 8, 2, 51, 49},
		{15, 7, 3, 31, 69},
		{5, 6, 4, 16, 84},
		{4, 5, 5, 11, 89},
		{1, 2, 8, 2, 98},
		{0, 0, 10, 0, 100},
	}

	for i, tc := range testCases {
		low, high := divideValidators(validators, tc.thresholdBasisPoints)
		assert.Equal(t, low.Len(), tc.lowGroupSize, "test case: %d ls.size", i)
		assert.Equal(t, low.TotalWeight(), tc.lowGroupWeight, "test case: %d ls.weight", i)
		assert.Equal(t, high.Len(), tc.highGroupSize, "test case: %d hs.size", i)
		assert.Equal(t, high.TotalWeight(), tc.highGroupWeight, "test case: %d hs.weight", i)
	}

}
