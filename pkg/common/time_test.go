// Copyright (C) Subhajit DasGupta 2021

package common

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestTimeParseText(t *testing.T) {
	assert := require.New(t)

	t1 := Time{}

	str := t1.TextValue()

	t2 := Time{}

	t2.Parse(str)

	assert.Equal(t1, t2)
	assert.Equal(t1.TextValue(), t2.TextValue())
	assert.Equal("", cmp.Diff(t1, t2))

	content, err := t1.MarshalText()
	assert.Nil(err)

	assert.Nil(t2.UnmarshalText(content))

	assert.Equal(t1, t2)
	assert.Equal(t1.TextValue(), t2.TextValue())
	assert.Equal("", cmp.Diff(t1, t2))

	content, err = json.Marshal(t1)
	assert.Nil(err)

	assert.Nil(json.Unmarshal(content, &t2))
	assert.Equal(t1.TextValue(), t2.TextValue())
	assert.Equal("", cmp.Diff(t1, t2))

	t1 = Time{}
	content, err = json.Marshal(t1)
	assert.Nil(err)

	assert.Nil(json.Unmarshal(content, &t2))
	assert.Equal("", cmp.Diff(t1, t2))

	t1 = Now()
	content, err = json.Marshal(t1)
	assert.Nil(err)

	assert.Nil(json.Unmarshal(content, &t2))
	assert.Equal("", cmp.Diff(t1, t2))
}
