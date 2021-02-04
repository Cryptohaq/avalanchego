// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcdb

import (
	"github.com/ava-labs/avalanchego/database"
)

var (
	errCodeToError = map[uint32]error{
		0: nil,
		1: database.ErrClosed,
		2: database.ErrNotFound,
		3: database.ErrAvoidCorruption,
	}
	errorToErrCode = map[error]uint32{
		nil:                         0,
		database.ErrClosed:          1,
		database.ErrNotFound:        2,
		database.ErrAvoidCorruption: 3,
	}
)

func errorToRPCError(err error) error {
	switch err {
	case database.ErrClosed, database.ErrNotFound, database.ErrAvoidCorruption:
		return nil
	default:
		return err
	}
}
