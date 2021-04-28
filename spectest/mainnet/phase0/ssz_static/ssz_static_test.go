package ssz_static

import (
	"testing"

	"github.com/prysmaticlabs/prysm/spectest/shared/phase0/ssz_static"
)

func TestMainnet_Phase0_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "mainnet")
}
