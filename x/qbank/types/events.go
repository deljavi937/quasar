package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeEvtFundsDeposited = "funds_deposited"
	TypeEvtFundsWithdrawn = "funds_withdrawn"

	AttributeValueCategory          = ModuleName
	AttributeKeyDepositedCoin       = "deposited_coin"
	AttributeKeyDepositLockupPeriod = "deposit_lockup_period"
	AttributeKeyDepositEpoch        = "deposit_epoch"
)

func CreateFundsDepositedEvent(ctx sdk.Context, sender sdk.AccAddress, coin sdk.Coin, lockupPeriod LockupTypes, currentEpoch uint64) sdk.Event {
	return sdk.NewEvent(
		TypeEvtFundsDeposited,
		sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
		sdk.NewAttribute(sdk.AttributeKeySender, sender.String()),
		sdk.NewAttribute(AttributeKeyDepositedCoin, coin.String()),
		sdk.NewAttribute(AttributeKeyDepositLockupPeriod, lockupPeriod.String()),
		sdk.NewAttribute(AttributeKeyDepositEpoch, strconv.FormatUint(currentEpoch, 10)),
	)
}
