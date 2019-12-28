package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseSyscallDuation(t *testing.T) {
	cases := []struct {
		line     string
		duration time.Duration
		error    bool
	}{
		{"", 0, true},
		{"test <abc>", 0, true},
		{"test <0123", 0, true},
		{"test 0123>", 0, true},
		{"<test< line >> < <1.2>\n", time.Millisecond * 1200, false},
		{
			"21:21:31 getsockopt(1, SOL_SOCKET, SO_ERROR, [0], [4]) = 0 <0.000016>",
			16 * time.Microsecond,
			false,
		},
	}
	assert := assert.New(t)
	for i, c := range cases {
		d, err := parseSyscallDuration(c.line)
		if c.error {
			assert.Errorf(err, "test %s", i)
		} else {
			assert.NoErrorf(err, "test %d", i)
		}
		assert.Equalf(c.duration, d, "test %d", i)
	}
}
