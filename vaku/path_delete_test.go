package vaku_test

import (
	"testing"

	"github.com/Lingrino/vaku/vaku"
	"github.com/stretchr/testify/assert"
)

type TestPathDeleteData struct {
	input     *vaku.PathInput
	outputErr bool
}

func TestPathDelete(t *testing.T) {
	c := clientInitForTests(t)
	defer seed(t, c)

	tests := map[int]TestPathDeleteData{
		1: {
			input:     vaku.NewPathInput("secretv1/test/foo"),
			outputErr: false,
		},
		2: {
			input:     vaku.NewPathInput("secretv2/test/foo"),
			outputErr: false,
		},
		3: {
			input:     vaku.NewPathInput("secretv1/doesnotexist"),
			outputErr: false,
		},
		4: {
			input:     vaku.NewPathInput("secretv2/doesnotexist"),
			outputErr: false,
		},
		5: {
			input:     vaku.NewPathInput("secretdoesnotexist/test/foo"),
			outputErr: true,
		},
	}

	for _, d := range tests {
		e := c.PathDelete(d.input)
		r, re := c.PathRead(d.input)
		if d.outputErr {
			assert.Error(t, e)
		} else {
			if re == nil {
				assert.Equal(t, "SECRET_HAS_BEEN_DELETED", r["VAKU_STATUS"])
			} else {
				assert.Error(t, re)
			}
			assert.NoError(t, e)
		}
	}
}