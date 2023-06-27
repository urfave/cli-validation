package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testcase[T any] struct {
	name        string
	f           func(T) error
	input       T
	errExpected bool
}

func test_with_input[T any](t *testing.T, tcases []testcase[T]) {
	r := require.New(t)
	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.f(tc.input)
			if tc.errExpected {
				r.Error(err)
			} else {
				r.NoError(err)
			}
		})
	}
}

func TestMin(t *testing.T) {

	inputs := []testcase[int]{
		{
			name:  "min lower limit",
			f:     Min[int](10),
			input: 10,
		},
		{
			name:  "min normal",
			f:     Min[int](10),
			input: 11,
		},
		{
			name:        "min failure",
			f:           Min[int](10),
			input:       8,
			errExpected: true,
		},
	}

	test_with_input[int](t, inputs)
}

func TestMax(t *testing.T) {

	inputs := []testcase[float64]{
		{
			name:  "max upper limit",
			f:     Max[float64](10.0),
			input: 10.0,
		},
		{
			name:  "max normal",
			f:     Max[float64](10.0),
			input: 9.0,
		},
		{
			name:        "max failure",
			f:           Max[float64](10.0),
			input:       10.00001,
			errExpected: true,
		},
	}

	test_with_input[float64](t, inputs)
}

func TestRange(t *testing.T) {

	inputs := []testcase[uint64]{
		{
			name:  "lower limit",
			f:     RangeInclusive[uint64](10, 16),
			input: 10,
		},
		{
			name:  "upper limit",
			f:     RangeInclusive[uint64](10, 16),
			input: 16,
		},
		{
			name:        "lower limit failure",
			f:           RangeInclusive[uint64](10, 16),
			input:       9,
			errExpected: true,
		},
		{
			name:        "upper limit failure",
			f:           RangeInclusive[uint64](10, 16),
			input:       17,
			errExpected: true,
		},
		{
			name:  "normal",
			f:     RangeInclusive[uint64](10, 16),
			input: 12,
		},
	}

	test_with_input[uint64](t, inputs)
}

func TestDisjointRange(t *testing.T) {

	inputs := []testcase[int16]{
		{
			name:  "lower limit lower range",
			f:     ValidationChainAny[int16](RangeInclusive[int16](10, 16), RangeInclusive[int16](56, 67)),
			input: 10,
		},
		{
			name:  "upper limit lower range",
			f:     ValidationChainAny[int16](RangeInclusive[int16](10, 16), RangeInclusive[int16](56, 67)),
			input: 16,
		},
		{
			name:  "lower limit upper range",
			f:     ValidationChainAny[int16](RangeInclusive[int16](10, 16), RangeInclusive[int16](56, 67)),
			input: 56,
		},
		{
			name:  "upper limit upper range",
			f:     ValidationChainAny[int16](RangeInclusive[int16](10, 16), RangeInclusive[int16](56, 67)),
			input: 67,
		},
		{
			name:  "normal lower range",
			f:     ValidationChainAny[int16](RangeInclusive[int16](10, 16), RangeInclusive[int16](56, 67)),
			input: 13,
		},
		{
			name:  "normal upper range",
			f:     ValidationChainAny[int16](RangeInclusive[int16](10, 16), RangeInclusive[int16](56, 67)),
			input: 60,
		},
		{
			name:        "failure below lower range",
			f:           ValidationChainAny[int16](RangeInclusive[int16](10, 16), RangeInclusive[int16](56, 67)),
			input:       1,
			errExpected: true,
		},
		{
			name:        "failure in between lower and upper range",
			f:           ValidationChainAny[int16](RangeInclusive[int16](10, 16), RangeInclusive[int16](56, 67)),
			input:       20,
			errExpected: true,
		},
		{
			name:        "failure above upper range",
			f:           ValidationChainAny[int16](RangeInclusive[int16](10, 16), RangeInclusive[int16](56, 67)),
			input:       70,
			errExpected: true,
		},
	}

	test_with_input[int16](t, inputs)
}