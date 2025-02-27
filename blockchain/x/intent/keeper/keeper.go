// Copyright 2024
//
// This file includes work covered by the following copyright and permission notices:
//
// Copyright 2023 Qredo Ltd.
// Licensed under the Apache License, Version 2.0;
//
// This file is part of the Warden Protocol library.
//
// The Warden Protocol library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Warden Protocol library. If not, see https://github.com/warden-protocol/wardenprotocol/blob/main/LICENSE
package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/warden-protocol/wardenprotocol/intent"
	"github.com/warden-protocol/wardenprotocol/repo"
	"github.com/warden-protocol/wardenprotocol/x/intent/types"
)

type (
	Keeper struct {
		cdc                     codec.BinaryCodec
		storeKey                storetypes.StoreKey
		memKey                  storetypes.StoreKey
		paramstore              paramtypes.Subspace
		actionHandlers          map[string]func(sdk.Context, *types.Action, *cdctypes.Any) (any, error)
		intentGeneratorHandlers map[string]func(sdk.Context, *cdctypes.Any) (intent.Intent, error)
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		actionHandlers:          make(map[string]func(sdk.Context, *types.Action, *cdctypes.Any) (any, error)),
		intentGeneratorHandlers: make(map[string]func(sdk.Context, *cdctypes.Any) (intent.Intent, error)),
	}
}

func (Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) IntentRepo() *repo.ObjectRepo[*types.Intent] {
	return &repo.ObjectRepo[*types.Intent]{
		Constructor: func() *types.Intent { return &types.Intent{} },
		StoreKey:    k.storeKey,
		Cdc:         k.cdc,
		CountKey:    types.KeyPrefix(types.IntentCountKey),
		ObjKey:      types.KeyPrefix(types.IntentKey),
	}
}
