// Copyright © 2020 David Arnold <dar@xoe.solutions>
// SPDX-License-Identifier: MIT

package command

import (
	"context"

	"github.com/satori/go.uuid"

	cerrors "github.com/xoe-labs/vicidial-go/internal/common/errors"

	// import adapters to other services (if any) here

	"github.com/xoe-labs/vicidial-go/internal/livecalling/app"

	"github.com/xoe-labs/vicidial-go/internal/livecalling/domain"
	"github.com/xoe-labs/vicidial-go/internal/livecalling/domain/livecall"
)

const (
	// ErrNotAuthorizedToPlayAudioToRemoteParty signals that the command issuer is not authorized to play the audio to the remote party
	ErrNotAuthorizedToPlayAudioToRemoteParty = cerrors.CommonError("ErrNotAuthorizedToPlayAudioToRemoteParty")
)

// PlayAudioToRemotePartyHandler knows how to play an audio
type PlayAudioToRemotePartyHandler struct {
	aggregate domain.Repository
	commMgr   domain.CommManager
	policy    app.Policy
}

// NewPlayAudioToRemotePartyHandler returs a PlayAudioToRemotePartyHandler
func NewPlayAudioToRemotePartyHandler(aggregate domain.Repository, commMgr domain.CommManager) PlayAudioToRemotePartyHandler {
	if aggregate == nil {
		panic("nil aggregate")
	}
	if commMgr == nil {
		panic("nil commMgr")
	}
	return PlayAudioToRemotePartyHandler{aggregate: aggregate, commMgr: commMgr}
}

// Handle plays an audio to the remote party
func (h PlayAudioToRemotePartyHandler) Handle(ctx context.Context, livecallToPlayAudioToRemotePartyFor, userID uuid.UUID, elevationToken string) error {

	if err := h.aggregate.Update(ctx, livecallToPlayAudioToRemotePartyFor, func(l *livecall.Livecall) error {
		if ok := h.policy.Can(ctx, "PlayAudioToRemoteParty", userID, elevationToken, *l); !ok {
			return ErrNotAuthorizedToPlayAudioToRemoteParty
		}
		if err := l.PlayAudioToRemoteParty(h.commMgr); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}