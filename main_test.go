package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	const equal, notEqual = true, false
	cases := map[string]struct {
		fn1  FullName
		fn2  FullName
		want bool
	}{
		"同姓同名（アルファベット）": {FullName{"john", "smith"}, FullName{"john", "smith"}, equal},
		"同姓同名（日本語）":     {FullName{"坂口", "たいが"}, FullName{"坂口", "たいが"}, equal},
		"姓名逆":           {FullName{"john", "smith"}, FullName{"smith", "john"}, notEqual},
		"異姓異名":          {FullName{"john", "smith"}, FullName{"adam", "bob"}, notEqual},
	}

	for k, v := range cases {
		name := k
		tc := v
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := tc.fn1.Equal(&tc.fn2)
			assert.Equal(t, got, tc.want)
		})
	}
}

func TestNewFullName(t *testing.T) {
	cases := map[string]struct {
		firstName string
		lastName  string
		isError   bool
	}{
		"first and last name is valid": {"first", "last", false},
		"first name is empty":          {"", "last", true},
		"last name is empty":           {"first", "", true},
	}

	for k, v := range cases {
		name := k
		tc := v
		t.Run(name, func(t *testing.T) {
			fn, err := NewFullName(tc.firstName, tc.lastName)
			if tc.isError {
				assert.Nil(t, fn)
				assert.Error(t, err)
			} else {
				assert.NotNil(t, fn)
				assert.Equal(t, &FullName{tc.firstName, tc.lastName}, fn)
				assert.NoError(t, err)
			}
		})
	}
}
