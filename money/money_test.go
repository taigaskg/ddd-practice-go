package money

import (
	"fmt"
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
	m := Money{10, "a"}
	// n := Money{12, "aa"}
	mm, err := m.Add(&Money{})
	fmt.Println(mm, err)
}
