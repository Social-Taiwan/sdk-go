package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "insurance"
	StoreKey   = ModuleName
)

var (
	// Keys for store prefixes
	BidsKey = []byte{0x01}

	// Key for insurance prefixes
	InsuranceFundPrefixKey = []byte{0x02}

	// Key for insurance redemption prefixes
	RedemptionSchedulePrefixKey = []byte{0x03, 0x05}

	GlobalShareDenomIdPrefixKey         = []byte{0x04, 0x00}
	GlobalRedemptionScheduleIdPrefixKey = []byte{0x05, 0x00}
)

// GetRedemptionScheduleKey provides the key to store a single pending redemption
func GetRedemptionScheduleKey(redemptionID uint64, claimTime time.Time) []byte {
	key := append(RedemptionSchedulePrefixKey, sdk.FormatTimeBytes(claimTime)...)
	key = append(key, sdk.Uint64ToBigEndian(redemptionID)...)
	return key
}

// GetRedemptionScheduleKey provides the key to store a single pending redemption
func (sh RedemptionSchedule) GetRedemptionScheduleKey() []byte {
	return GetRedemptionScheduleKey(sh.Id, sh.ClaimableRedemptionTime)
}
