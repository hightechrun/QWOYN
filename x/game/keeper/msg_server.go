package keeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmic-horizon/coho/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) TransferModuleOwnership(goCtx context.Context, msg *types.MsgTransferModuleOwnership) (*types.MsgTransferModuleOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Sender != params.Owner {
		return nil, types.ErrNotModuleOwner
	}
	params.Owner = msg.NewOwner
	m.SetParamSet(ctx, params)
	return &types.MsgTransferModuleOwnershipResponse{}, nil
}

func (m msgServer) WhitelistNftContracts(goCtx context.Context, msg *types.MsgWhitelistNftContracts) (*types.MsgWhitelistNftContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Sender != params.Owner {
		return nil, types.ErrNotModuleOwner
	}

	moduleAddr := m.AccountKeeper.GetModuleAddress(types.ModuleName)
	for _, contract := range msg.Contracts {
		contractAddr, err := sdk.AccAddressFromBech32(contract)
		if err != nil {
			return nil, err
		}

		minterJSON, err := m.WasmViewer.QuerySmart(ctx, contractAddr, []byte(`{"minter": {}}`))

		var parsed map[string]string
		err = json.Unmarshal(minterJSON, &parsed)
		if err != nil {
			return nil, err
		}
		if parsed["minter"] != moduleAddr.String() {
			fmt.Println("minter", parsed["minter"])
			return nil, types.ErrMinterIsNotModuleAddress
		}

		contractInfoJSON, err := m.WasmViewer.QuerySmart(ctx, contractAddr, []byte(`{"contract_info": {}}`))
		err = json.Unmarshal(contractInfoJSON, &parsed)
		if err != nil {
			return nil, err
		}
		if parsed["owner"] != moduleAddr.String() {
			fmt.Println("owner", parsed["owner"])
			return nil, types.ErrOwnerIsNotModuleAddress
		}

		m.SetWhitelistedContract(ctx, contract)
	}
	return &types.MsgWhitelistNftContractsResponse{}, nil
}

func (m msgServer) RemoveWhitelistedNftContracts(goCtx context.Context, msg *types.MsgRemoveWhitelistedNftContracts) (*types.MsgRemoveWhitelistedNftContractsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Sender != params.Owner {
		return nil, types.ErrNotModuleOwner
	}
	for _, contract := range msg.Contracts {
		m.DeleteWhitelistedContract(ctx, contract)
	}
	return &types.MsgRemoveWhitelistedNftContractsResponse{}, nil
}

func (m msgServer) DepositNft(goCtx context.Context, msg *types.MsgDepositNft) (*types.MsgDepositNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.Keeper.DepositNft(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositNftResponse{}, nil
}

func (m msgServer) WithdrawUpdatedNft(goCtx context.Context, msg *types.MsgWithdrawUpdatedNft) (*types.MsgWithdrawUpdatedNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.Keeper.WithdrawUpdatedNft(ctx, msg)
	if err != nil {
		return nil, err
	}
	return &types.MsgWithdrawUpdatedNftResponse{}, nil
}

func (m msgServer) DepositToken(goCtx context.Context, msg *types.MsgDepositToken) (*types.MsgDepositTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Amount.Denom != params.DepositDenom {
		return nil, types.ErrInvalidDepositDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.BankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	m.IncreaseDeposit(ctx, sender, msg.Amount)

	return &types.MsgDepositTokenResponse{}, nil
}

func (m msgServer) WithdrawToken(goCtx context.Context, msg *types.MsgWithdrawToken) (*types.MsgWithdrawTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := m.GetParamSet(ctx)
	if msg.Amount.Denom != params.DepositDenom {
		return nil, types.ErrInvalidWithdrawDenom
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	err = m.Keeper.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, sdk.Coins{msg.Amount})
	if err != nil {
		return nil, err
	}

	m.DecreaseDeposit(ctx, sender, msg.Amount)

	return &types.MsgWithdrawTokenResponse{}, nil
}
