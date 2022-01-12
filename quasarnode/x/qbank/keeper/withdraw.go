package keeper

import (
	"encoding/binary"

	"github.com/abag/quasarnode/x/qbank/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetWithdrawCount get the total number of withdraw
func (k Keeper) GetWithdrawCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.QbankGlobalKBP)
	byteKey := types.CreateWithdrawCountKey()
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetWithdrawCount set the total number of withdraw
func (k Keeper) SetWithdrawCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.QbankGlobalKBP)
	byteKey := types.CreateWithdrawCountKey()
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendWithdraw appends a withdraw in the store with a new id and update the count
func (k Keeper) AppendWithdraw(
	ctx sdk.Context,
	withdraw types.Withdraw,
) uint64 {
	// Create the withdraw
	count := k.GetWithdrawCount(ctx)

	// Set the ID of the appended value
	withdraw.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WithdrawKeyKBP)
	appendedValue := k.cdc.MustMarshal(&withdraw)
	store.Set(types.CreateIDKey(withdraw.Id), appendedValue)

	// Update withdraw count
	k.SetWithdrawCount(ctx, count+1)

	return count
}

// SetWithdraw set a specific withdraw in the store
func (k Keeper) SetWithdraw(ctx sdk.Context, withdraw types.Withdraw) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WithdrawKeyKBP)
	b := k.cdc.MustMarshal(&withdraw)
	store.Set(types.CreateIDKey(withdraw.Id), b)
}

// GetWithdraw returns a withdraw from its id
func (k Keeper) GetWithdraw(ctx sdk.Context, id uint64) (val types.Withdraw, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WithdrawKeyKBP)
	b := store.Get(types.CreateIDKey(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWithdraw removes a withdraw from the store
func (k Keeper) RemoveWithdraw(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WithdrawKeyKBP)
	store.Delete(types.CreateIDKey(id))
}

// GetAllWithdraw returns all withdraw
func (k Keeper) GetAllWithdraw(ctx sdk.Context) (list []types.Withdraw) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.WithdrawKeyKBP)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Withdraw
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetWithdrawIDBytes returns the byte representation of the ID
func GetWithdrawIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetWithdrawIDFromBytes returns ID in uint64 format from a byte array
func GetWithdrawIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
