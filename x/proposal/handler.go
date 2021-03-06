package proposal

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateProposal:
			return handleMsgCreateProposal(ctx, keeper, msg)
		case MsgVote:
			return handleMsgVote(ctx, keeper, msg)
		case MsgTryExecuteProposal:
			return handleMsgTryExecuteProposal(ctx, keeper, msg)
		case MsgWithdrawProposal:
			return handleMsgWithdrawProposal(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized data Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}


func handleMsgCreateProposal(ctx sdk.Context, keeper Keeper, msg MsgCreateProposal) sdk.Result {
	return keeper.Propose(ctx, msg.Proposer, msg.Action)
}

func handleMsgVote(ctx sdk.Context, keeper Keeper, msg MsgVote) sdk.Result {
	return keeper.Vote(ctx, msg.ProposalId, msg.Voter, msg.Vote)
}

func handleMsgTryExecuteProposal(ctx sdk.Context, keeper Keeper, msg MsgTryExecuteProposal) sdk.Result {
	return keeper.TryExecute(ctx, msg.ProposalId)
}

func handleMsgWithdrawProposal(ctx sdk.Context, keeper Keeper, msg MsgWithdrawProposal) sdk.Result {
	return keeper.Withdraw(ctx, msg.ProposalId, msg.Proposer)
}

