package jutil_test

import (
	"bytes"
	"github.com/jchavannes/jgo/jerr"
	"github.com/jchavannes/jgo/jutil"
	"testing"
)

func TestHasPrefix(t *testing.T) {
	var b = []byte{0x6A, 0x02, 0x6d, 0x04}
	var prefix = []byte{0x6A}
	if !jutil.HasPrefix(b, prefix) {
		t.Error(jerr.Newf("error prefix did not match (b: %x, prefix: %x)", b, prefix))
	}
	if jutil.HasPrefix(prefix, b) {
		t.Error(jerr.Newf("error prefix unexpectedly matched (b: %x, prefix: %x)", prefix, b))
	}
}

func TestGetUint32Data(t *testing.T) {
	jutil.GetUint32Data(23)
	jutil.GetUint32Data(1234567890)
}

func TestBytePadPrefix(t *testing.T) {
	var input = []byte{0x00, 0x01}
	originalInputSize := len(input)
	const PadSize = 10
	paddedInput := jutil.BytePadPrefix(input, PadSize)
	if len(paddedInput) != PadSize {
		t.Error(jerr.Newf("error padded input does not match pad size: %d %d", len(paddedInput), PadSize))
		return
	}
	if len(input) != originalInputSize {
		t.Error(jerr.Newf("error input size does not match original input size: %d %d", len(input), originalInputSize))
		return
	}
	unpaddedInput := jutil.ByteUnPad(paddedInput)
	if !bytes.Equal(unpaddedInput, input) {
		t.Error(jerr.Newf("error unpadded input does not match original input"))
	}
}
