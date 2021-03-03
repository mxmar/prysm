package types

import (
	types "github.com/prysmaticlabs/eth2-types"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

// ChunkKind to differentiate what kind of span we are working
// with for slashing detection, either min or max span.
type ChunkKind uint

const (
	MinSpan ChunkKind = iota
	MaxSpan
)

// IndexedAttestationWrapper contains an indexed attestation with its
// signing root to reduce duplicated computation.
type IndexedAttestationWrapper struct {
	IndexedAttestation *ethpb.IndexedAttestation
	SigningRoot        [32]byte
}

// AttesterDoubleVote represents a double vote instance
// which is a slashable event for attesters.
type AttesterDoubleVote struct {
	ValidatorIndex  types.ValidatorIndex
	Target          types.Epoch
	SigningRoot     [32]byte
	PrevSigningRoot [32]byte
}

// DoubleBlockProposal containing an incoming and an existing proposal's signing root.
type DoubleBlockProposal struct {
	Slot                types.Slot
	ProposerIndex       types.ValidatorIndex
	IncomingSigningRoot [32]byte
	ExistingSigningRoot [32]byte
}

// AttestedEpochForValidator encapsulates a previously attested epoch
// for a validator index.
type AttestedEpochForValidator struct {
	ValidatorIndex types.ValidatorIndex
	Epoch          types.Epoch
}

// SignedBlockHeaderWrapper contains an signed beacon block header with its
// signing root to reduce duplicated computation.
type SignedBlockHeaderWrapper struct {
	SignedBeaconBlockHeader *ethpb.SignedBeaconBlockHeader
	SigningRoot             [32]byte
}

// Slashing represents a compact format with all the information
// needed to understand a slashable offense in eth2.
type Slashing struct {
	Kind            SlashingKind
	ValidatorIndex  types.ValidatorIndex
	PrevSourceEpoch types.Epoch
	PrevTargetEpoch types.Epoch
	SourceEpoch     types.Epoch
	TargetEpoch     types.Epoch
	SigningRoot     [32]byte
	PrevSigningRoot [32]byte
	Slot            types.Slot
}

// SlashingKind is an enum representing the type of slashable
// offense detected by slasher, useful for conditionals or for logging.
type SlashingKind int

const (
	NotSlashable SlashingKind = iota
	DoubleVote
	SurroundingVote
	SurroundedVote
	DoubleProposal
)

func (k SlashingKind) String() string {
	switch k {
	case NotSlashable:
		return "NOT_SLASHABLE"
	case DoubleVote:
		return "DOUBLE_VOTE"
	case SurroundingVote:
		return "SURROUNDING_VOTE"
	case SurroundedVote:
		return "SURROUNDED_VOTE"
	default:
		return "UNKNOWN"
	}
}
