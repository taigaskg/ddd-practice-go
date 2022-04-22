package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cases := map[string]struct {
		amount   uint32
		currency string
		isError  bool
	}{
		"empty currency": {10, "", true},
		"valid":          {10, "aaa", false},
		"amount is zero": {0, "aaa", false},
	}

	for k, v := range cases {
		name := k
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			m, err := New(tc.amount, tc.currency)
			if tc.isError {
				assert.Nil(t, m)
				assert.Error(t, err)
			} else {
				assert.NotNil(t, m)
				assert.Equal(t, &Money{tc.amount, tc.currency}, m)
				assert.NoError(t, err)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	cases := map[string]struct {
		m1             Money
		m2             Money
		expectedAmount uint32
		isError        bool
	}{
		"same currency":      {Money{10, "jpy"}, Money{20, "jpy"}, 30, false},
		"different currency": {Money{10, "jpy"}, Money{20, "usd"}, 0, true},
	}

	for k, v := range cases {
		name := k
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			m, err := tc.m1.Add(&tc.m2)
			if tc.isError {
				assert.Nil(t, m)
				assert.Error(t, err)
			} else {
				assert.NotNil(t, m)
				assert.Equal(t, tc.expectedAmount, m.Amount)
				assert.NoError(t, err)
			}
		})
	}
}
