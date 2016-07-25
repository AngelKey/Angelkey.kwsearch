package index

import (
	"crypto/sha256"
	"search/util"
	"testing"

	"github.com/jxguan/go-datastructures/bitarray"
)

// TestMarshalAndUnmarshal tests the `Marshal` and `Unmarshal` functions.
// Checks that after a pair of `Marshal` and `Unmarshal` operations, the orignal
// SecureIndex is correctly reconstructed from the byte slice.
func TestMarshalAndUnmarshal(t *testing.T) {
	si := new(SecureIndex)
	si.BloomFilter = bitarray.NewSparseBitArray()
	for i := 0; i < 1000; i++ {
		si.BloomFilter.SetBit(util.RandUint64n(1000000))
	}
	si.DocID = 42
	si.Size = uint64(1900000)
	si.Hash = sha256.New
	bytes := si.Marshal()
	si2 := Unmarshal(bytes)
	if si2.DocID != si.DocID {
		t.Fatalf("DocID does not match")
	}
	if si2.Hash().Size() != si.Hash().Size() {
		t.Fatalf("Hash does not match")
	}
	if si2.Size != si.Size {
		t.Fatalf("Size does not match")
	}
	if !si2.BloomFilter.Equals(si.BloomFilter) {
		t.Fatalf("BloomFilter does not mtach")
	}
}
