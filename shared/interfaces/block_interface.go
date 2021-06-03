package interfaces

import (
	types "github.com/prysmaticlabs/eth2-types"
	ethpb "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
	"google.golang.org/protobuf/proto"
)

// SignedBeaconBlock is an interface describing the method set of
// a signed beacon block.
type SignedBeaconBlock interface {
	Block() BeaconBlock
	Signature() []byte
	IsNil() bool
	Copy() SignedBeaconBlock
	MarshalSSZ() ([]byte, error)
	Proto() proto.Message
	PbPhase0Block() (*ethpb.SignedBeaconBlock, error)
	Version() int
}

// BeaconBlock describes an interface which states the methods
// employed by an object that is a beacon block.
type BeaconBlock interface {
	Slot() types.Slot
	ProposerIndex() types.ValidatorIndex
	ParentRoot() []byte
	StateRoot() []byte
	Body() BeaconBlockBody
	IsNil() bool
	HashTreeRoot() ([32]byte, error)
	MarshalSSZ() ([]byte, error)
	Proto() proto.Message
}

// BeaconBlockBody describes the method set employed by an object
// that is a beacon block body.
type BeaconBlockBody interface {
	RandaoReveal() []byte
	Eth1Data() *ethpb.Eth1Data
	Graffiti() []byte
	ProposerSlashings() []*ethpb.ProposerSlashing
	AttesterSlashings() []*ethpb.AttesterSlashing
	Attestations() []*ethpb.Attestation
	Deposits() []*ethpb.Deposit
	VoluntaryExits() []*ethpb.SignedVoluntaryExit
	SyncAggregate() (*ethpb.SyncAggregate, error)
	IsNil() bool
	HashTreeRoot() ([32]byte, error)
	Proto() proto.Message
}
