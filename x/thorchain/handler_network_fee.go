package thorchain

import (
	"github.com/blang/semver"

	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
	"gitlab.com/thorchain/thornode/constants"
)

// NetworkFeeHandler a handler to process MsgNetworkFee messages
type NetworkFeeHandler struct {
	mgr Manager
}

// NewNetworkFeeHandler create a new instance of network fee handler
func NewNetworkFeeHandler(mgr Manager) NetworkFeeHandler {
	return NetworkFeeHandler{
		mgr: mgr,
	}
}

// Run is the main entry point for network fee logic
func (h NetworkFeeHandler) Run(ctx cosmos.Context, m cosmos.Msg) (*cosmos.Result, error) {
	msg, ok := m.(*MsgNetworkFee)
	if !ok {
		return nil, errInvalidMessage
	}
	if err := h.validate(ctx, *msg); err != nil {
		ctx.Logger().Error("MsgNetworkFee failed validation", "error", err)
		return nil, err
	}
	result, err := h.handle(ctx, *msg)
	if err != nil {
		ctx.Logger().Error("fail to process MsgNetworkFee", "error", err)
	}
	return result, err
}

func (h NetworkFeeHandler) validate(ctx cosmos.Context, msg MsgNetworkFee) error {
	version := h.mgr.GetVersion()
	if version.GTE(semver.MustParse("0.1.0")) {
		return h.validateV1(ctx, msg)
	}
	return errBadVersion
}

func (h NetworkFeeHandler) validateV1(ctx cosmos.Context, msg MsgNetworkFee) error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	if !isSignedByActiveNodeAccounts(ctx, h.mgr, msg.GetSigners()) {
		return cosmos.ErrUnauthorized(notAuthorized.Error())
	}
	return nil
}

// handle process MsgNetworkFee
func (h NetworkFeeHandler) handle(ctx cosmos.Context, msg MsgNetworkFee) (*cosmos.Result, error) {
	version := h.mgr.GetVersion()
	if version.GTE(semver.MustParse("0.47.0")) {
		return h.handleV47(ctx, msg)
	} else if version.GTE(semver.MustParse("0.1.0")) {
		return h.handleV1(ctx, msg)
	}
	return nil, errBadVersion
}

// handleV47 process MsgNetworkFee
func (h NetworkFeeHandler) handleV47(ctx cosmos.Context, msg MsgNetworkFee) (*cosmos.Result, error) {
	active, err := h.mgr.Keeper().ListActiveValidators(ctx)
	if err != nil {
		err = wrapError(ctx, err, "fail to get list of active node accounts")
		return nil, err
	}

	voter, err := h.mgr.Keeper().GetObservedNetworkFeeVoterV47(ctx, msg.BlockHeight, msg.Chain, int64(msg.TransactionFeeRate))
	if err != nil {
		return nil, err
	}
	observeSlashPoints := h.mgr.GetConstants().GetInt64Value(constants.ObserveSlashPoints)
	observeFlex := h.mgr.GetConstants().GetInt64Value(constants.ObservationDelayFlexibility)
	h.mgr.Slasher().IncSlashPoints(ctx, observeSlashPoints, msg.Signer)
	if !voter.Sign(msg.Signer) {
		ctx.Logger().Info("signer already signed MsgNetworkFee", "signer", msg.Signer.String(), "block height", msg.BlockHeight, "chain", msg.Chain.String())
		return &cosmos.Result{}, nil
	}
	h.mgr.Keeper().SetObservedNetworkFeeVoter(ctx, voter)
	// doesn't have consensus yet
	if !voter.HasConsensus(active) {
		return &cosmos.Result{}, nil
	}

	if voter.BlockHeight > 0 {
		if (voter.BlockHeight + observeFlex) >= common.BlockHeight(ctx) {
			h.mgr.Slasher().DecSlashPoints(ctx, observeSlashPoints, msg.Signer)
		}
		// MsgNetworkFee tx already processed
		return &cosmos.Result{}, nil
	}

	voter.BlockHeight = common.BlockHeight(ctx)
	h.mgr.Keeper().SetObservedNetworkFeeVoter(ctx, voter)
	// decrease the slash points
	h.mgr.Slasher().DecSlashPoints(ctx, observeSlashPoints, voter.GetSigners()...)
	ctx.Logger().Info("update network fee", "chain", msg.Chain.String(), "transaction-size", msg.TransactionSize, "fee-rate", msg.TransactionFeeRate)
	if err := h.mgr.Keeper().SaveNetworkFee(ctx, msg.Chain, NetworkFee{
		Chain:              msg.Chain,
		TransactionSize:    msg.TransactionSize,
		TransactionFeeRate: msg.TransactionFeeRate,
	}); err != nil {
		return nil, ErrInternal(err, "fail to save network fee")
	}
	return &cosmos.Result{}, nil
}
