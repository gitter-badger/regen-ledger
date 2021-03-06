package data

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"math"
	"strconv"
)

// query endpoints supported by the governance Querier
const (
	QueryData = "get"
	QueryDataBlockHeight = "block-height"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryData:
			return queryData(ctx, path[1:], req, keeper)
		case QueryDataBlockHeight:
			return queryDataBlockHeight(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown data query endpoint")
		}
	}
}

func fromHex(str string) []byte {
	bz, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return bz
}

func queryData(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]

	value := keeper.GetData(ctx, fromHex(id))

	if len(value) == 0 {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve id")
	}

	return value, nil
}

func queryDataBlockHeight(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	id := path[0]

	value := keeper.GetDataBlockHeight(ctx, fromHex(id))

	if value == math.MaxInt64 {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve id")
	}

	return []byte(strconv.FormatInt(value, 10)), nil
}
