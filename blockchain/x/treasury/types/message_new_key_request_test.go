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
package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/warden-protocol/wardenprotocol/testutil/sample"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMsgNewKeyRequest_NewMsgNewKeyRequest(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgNewKeyRequest
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewKeyRequest{
				Creator:       sample.AccAddress(),
				SpaceAddr: "wardenspace14a2hpadpsy9h5sm54xj",
				KeychainAddr:   "wardenkeychain1ph63us46lyw56lmt585",
				KeyType:       1,
				Btl:           1000,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &MsgNewKeyRequest{tt.msg.Creator, tt.msg.SpaceAddr, tt.msg.KeychainAddr, tt.msg.KeyType, tt.msg.Btl}
			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgNewKeyRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgNewKeyRequest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgNewKeyRequest{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgNewKeyRequest{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
