// Copyright 2018 The SET Team Authors
// This file is part of the SET project.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package rpcapi

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/Second-Earth/setchain/common"
	"github.com/Second-Earth/setchain/types"
	"github.com/Second-Earth/setchain/utils/rlp"
)

// PublicsetAPI offers and API for the transaction pool. It only operates on data that is non confidential.
type PublicsetAPI struct {
	b Backend
}

// NewPublicsetAPI creates a new tx pool service that gives information about the transaction pool.
func NewPublicsetAPI(b Backend) *PublicsetAPI {
	return &PublicsetAPI{b}
}

// GasPrice returns a suggestion for a gas price.
func (s *PublicsetAPI) GasPrice(ctx context.Context) (*big.Int, error) {
	return s.b.SuggestPrice(ctx)
}

// SendRawTransaction will add the signed transaction to the transaction pool.
// The sender is responsible for signing the transaction and using the correct nonce.
func (s *PublicsetAPI) SendRawTransaction(ctx context.Context, encodedTx hexutil.Bytes) (common.Hash, error) {
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(encodedTx, tx); err != nil {
		return common.Hash{}, err
	}
	return submitTransaction(ctx, s.b, tx)
}
