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

// nolint:stylecheck,st1003
// revive:disable-next-line var-naming
func (sr *SignRequest) SetId(id uint64) { sr.Id = id }

// nolint:stylecheck,st1003
// revive:disable-next-line var-naming
func (str *SignTransactionRequest) SetId(id uint64) { str.Id = id }

func NewMsgFulfilSignatureRequestPayload(signedData []byte) isMsgFulfilSignatureRequest_Result {
	return &MsgFulfilSignatureRequest_Payload{
		Payload: &MsgSignedData{
			SignedData: signedData,
		},
	}
}

func NewMsgFulfilSignatureRequestReject(msg string) isMsgFulfilSignatureRequest_Result {
	return &MsgFulfilSignatureRequest_RejectReason{
		RejectReason: msg,
	}
}
