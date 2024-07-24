package core

import (
	"blockchain/types"
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     999383838,
	}

	buf := &bytes.Buffer{}

	if err := h.EncodeBinary(buf); err != nil {
		t.Fatalf("EncodeBinary: %v", err)
	}

	hDecode := &Header{}

	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}
