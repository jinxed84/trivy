package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_IntJSON(t *testing.T) {
	val := Int(0x66, NewMisconfigMetadata(NewRange("main.tf", 123, 123, "", nil), ""))
	data, err := json.Marshal(val)
	require.NoError(t, err)

	var restored IntValue
	err = json.Unmarshal(data, &restored)
	require.NoError(t, err)

	assert.Equal(t, val, restored)
}