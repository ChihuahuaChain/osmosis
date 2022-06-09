package balancer

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	ErrMsgFormatFailedInterimLiquidityUpdate = errMsgFormatFailedInterimLiquidityUpdate

	CalcPoolSharesOutGivenSingleAssetIn = calcPoolSharesOutGivenSingleAssetIn
	CalcSingleAssetInGivenPoolSharesOut = calcSingleAssetInGivenPoolSharesOut
	UpdateIntermediaryPoolAssets        = updateIntermediaryPoolAssets
)

func (p *Pool) CalcSingleAssetJoin(tokenIn sdk.Coin, swapFee sdk.Dec, tokenInPoolAsset PoolAsset, totalShares sdk.Int) (numShares sdk.Int, err error) {
	return p.calcSingleAssetJoin(tokenIn, swapFee, tokenInPoolAsset, totalShares)
}

func (p *Pool) CalcJoinSingleAssetTokensIn(tokensIn sdk.Coins, totalSharesSoFar sdk.Int, poolAssetsByDenom map[string]PoolAsset, swapFee sdk.Dec) (sdk.Int, sdk.Coins, error) {
	return p.calcJoinSingleAssetTokensIn(tokensIn, totalSharesSoFar, poolAssetsByDenom, swapFee)
}
