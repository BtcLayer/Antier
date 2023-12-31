package commitment

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

const (
	// tapBranchPreimageLen is the length of a TapBranch preimage, excluding
	// the TapBranchTag.
	tapBranchPreimageLen = 64
)

var (
	// ErrInvalidTaprootProof is an error returned upon verifying an invalid
	// Taproot proof.
	ErrInvalidTaprootProof = errors.New("invalid taproot proof")

	// ErrInvalidTapscriptProof is an error returned upon attempting to
	// prove a malformed TapscriptProof.
	ErrInvalidTapscriptProof = errors.New("invalid tapscript proof")

	// ErrInvalidEmptyTapscriptPreimage is an error returned upon attempting
	// to generate the tap hash of an empty preimage.
	ErrInvalidEmptyTapscriptPreimage = errors.New(
		"invalid empty tapscript preimage",
	)

	// ErrInvalidTapscriptPreimageLen is an error returned upon attempting
	// to generate the tap hash of a branch preimage with invalid length.
	ErrInvalidTapscriptPreimageLen = errors.New(
		"invalid tapscript preimage length",
	)

	// ErrPreimageIsTapCommitment is an error returned when a tapscript
	// preimage is a valid Taproot Asset commitment.
	ErrPreimageIsTapCommitment = errors.New(
		"preimage is a Taproot Asset commitment",
	)
)

// TapscriptPreimageType denotes the type of tapscript sibling preimage.
type TapscriptPreimageType uint8

const (
	// LeafPreimage is a pre-image that's a leaf script.
	LeafPreimage TapscriptPreimageType = 0

	// BranchPreimage is a pre-image that's a branch, so it's actually
	// 64-bytes of two child pre-images.
	BranchPreimage TapscriptPreimageType = 1
)

// String returns a human-readable string for the TapscriptPreimageType.
func (t TapscriptPreimageType) String() string {
	switch t {
	case LeafPreimage:
		return "LeafPreimage"

	case BranchPreimage:
		return "BranchPreimage"

	default:
		return fmt.Sprintf("UnKnownSiblingType(%d)", t)
	}
}

// TapscriptPreimage wraps a pre-image byte slice with a type byte that self
// identifies what type of pre-image it is.
type TapscriptPreimage struct {
	// SiblingPreimage is the pre-image itself. This will be either 32 or
	// 64 bytes.
	SiblingPreimage []byte

	// SiblingType is the type of the pre-image.
	SiblingType TapscriptPreimageType
}

// NewPreimageFromLeaf creates a new TapscriptPreimage from a single tap leaf.
func NewPreimageFromLeaf(leaf txscript.TapLeaf) *TapscriptPreimage {
	// The leaf encoding is: leafVersion || compactSizeof(script) ||
	// script, where compactSizeof returns the compact size needed to
	// encode the value.
	var encodedLeaf bytes.Buffer

	_ = encodedLeaf.WriteByte(byte(leaf.LeafVersion))
	_ = wire.WriteVarBytes(&encodedLeaf, 0, leaf.Script)

	return &TapscriptPreimage{
		SiblingPreimage: encodedLeaf.Bytes(),
		SiblingType:     LeafPreimage,
	}
}

// NewPreimageFromBranch creates a new TapscriptPreimage from a tap branch.
func NewPreimageFromBranch(branch txscript.TapBranch) *TapscriptPreimage {
	var (
		encodedBranch bytes.Buffer
		leftHash      = branch.Left().TapHash()
		rightHash     = branch.Right().TapHash()
	)
	_, _ = encodedBranch.Write(leftHash[:])
	_, _ = encodedBranch.Write(rightHash[:])

	return &TapscriptPreimage{
		SiblingPreimage: encodedBranch.Bytes(),
		SiblingType:     BranchPreimage,
	}
}

// IsEmpty returns true if the sibling pre-image is empty.
func (t *TapscriptPreimage) IsEmpty() bool {
	if t == nil {
		return true
	}

	return len(t.SiblingPreimage) == 0
}

// TapHash returns the tap hash of this preimage according to its type.
func (t *TapscriptPreimage) TapHash() (*chainhash.Hash, error) {
	if t.IsEmpty() {
		return nil, ErrInvalidEmptyTapscriptPreimage
	}

	switch t.SiblingType {
	// The sibling is actually a leaf pre-image, so we'll verify that it
	// isn't a Taproot Asset commitment, and then hash it with the
	// commitment to obtain our root.
	case LeafPreimage:
		return TapLeafHash(t.SiblingPreimage)

	// The sibling is actually a branch pre-image, so we'll verify that the
	// branch pre-image is 64-bytes (the two 32-byte hashes of the left
	// and right nodes), and then derive the key from that.
	case BranchPreimage:
		return TapBranchHash(t.SiblingPreimage)

	default:
		return nil, fmt.Errorf("unknown sibling type: <%d>",
			t.SiblingType)
	}
}

// VerifyNoCommitment verifies that the preimage is not a Taproot Asset
// commitment.
func (t *TapscriptPreimage) VerifyNoCommitment() error {
	if IsTaprootAssetCommitmentScript(t.SiblingPreimage) {
		return ErrPreimageIsTapCommitment
	}

	return nil
}

// MaybeDecodeTapscriptPreimage returns the decoded preimage and hash of the
// Tapscript sibling or nil if the encoded content is empty.
func MaybeDecodeTapscriptPreimage(encoded []byte) (*TapscriptPreimage,
	*chainhash.Hash, error) {

	if len(encoded) == 0 {
		return nil, nil, nil
	}

	var (
		preimage = &TapscriptPreimage{}
		scratch  [8]byte
	)
	err := TapscriptPreimageDecoder(
		bytes.NewReader(encoded), &preimage, &scratch,
		uint64(len(encoded)),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding tapscript "+
			"preimage: %w", err)
	}

	tapHash, err := preimage.TapHash()
	if err != nil {
		return nil, nil, fmt.Errorf("error deriving tap hash "+
			"from preimage: %w", err)
	}

	return preimage, tapHash, nil
}

// MaybeEncodeTapscriptPreimage returns the encoded preimage and hash of the
// Tapscript sibling or nil if the preimage is nil.
func MaybeEncodeTapscriptPreimage(t *TapscriptPreimage) ([]byte,
	*chainhash.Hash, error) {

	if t == nil {
		return nil, nil, nil
	}

	tapHash, err := t.TapHash()
	if err != nil {
		return nil, nil, fmt.Errorf("error deriving tap hash "+
			"from preimage: %w", err)
	}

	var (
		b       bytes.Buffer
		scratch [8]byte
	)
	if err := TapscriptPreimageEncoder(&b, &t, &scratch); err != nil {
		return nil, tapHash, fmt.Errorf("error encoding tapscript "+
			"preimage: %w", err)
	}

	return b.Bytes(), tapHash, nil
}

// NewTapBranchHash takes the raw tap hashes of the left and right nodes and
// hashes them into a branch.
func NewTapBranchHash(l, r chainhash.Hash) chainhash.Hash {
	if bytes.Compare(l[:], r[:]) > 0 {
		l, r = r, l
	}

	return *chainhash.TaggedHash(chainhash.TagTapBranch, l[:], r[:])
}

// TapBranchHash computes the TapHash of a TapBranch node from its preimage
// if possible, otherwise an error is returned.
func TapBranchHash(preimage []byte) (*chainhash.Hash, error) {
	// Empty preimage or leaf preimage, return typed error.
	if len(preimage) != tapBranchPreimageLen {
		return nil, ErrInvalidTapscriptPreimageLen
	}

	left := (*chainhash.Hash)(preimage[:chainhash.HashSize])
	right := (*chainhash.Hash)(preimage[chainhash.HashSize:])
	h := NewTapBranchHash(*left, *right)

	return &h, nil
}

// TapLeafHash computes the TapHash of a TapLeaf node from its preimage
// if possible, otherwise an error is returned.
func TapLeafHash(preimage []byte) (*chainhash.Hash, error) {
	// Empty preimage.
	if len(preimage) == 0 {
		return nil, ErrInvalidEmptyTapscriptPreimage
	}

	// Enforce that it is not including another Taproot Asset commitment.
	if bytes.Contains(preimage, TaprootAssetsMarker[:]) {
		return nil, ErrInvalidTaprootProof
	}

	version := txscript.TapscriptLeafVersion(preimage[0])

	// The script is encoded with a leading VarByte that indicates its total
	// length.
	remaining := preimage[1:]
	script, err := wire.ReadVarBytes(
		bytes.NewReader(remaining), 0, uint32(len(remaining)), "script",
	)
	if err != nil {
		return nil, fmt.Errorf("error decoding leaf pre-image: %w", err)
	}

	h := txscript.NewTapLeaf(version, script).TapHash()
	return &h, nil
}
