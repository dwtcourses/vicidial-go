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
	// ErrNotAuthorizedToStartRecording signals that the command issuer is not authorized to start the recording
	ErrNotAuthorizedToStartRecording = cerrors.CommonError("ErrNotAuthorizedToStartRecording")
)

// StartRecordingHandler knows how to start a recording
type StartRecordingHandler struct {
	aggregate domain.Repository
	commMgr   domain.CommManager
	policy    app.Policy
}

// NewStartRecordingHandler returs a StartRecordingHandler
func NewStartRecordingHandler(aggregate domain.Repository, commMgr domain.CommManager) StartRecordingHandler {
	if aggregate == nil {
		panic("nil aggregate")
	}
	if commMgr == nil {
		panic("nil commMgr")
	}
	return StartRecordingHandler{aggregate: aggregate, commMgr: commMgr}
}

// Handle starts a recording
func (h StartRecordingHandler) Handle(ctx context.Context, livecallToStartRecordingFor, userID uuid.UUID, elevationToken string) error {

	if err := h.aggregate.Update(ctx, livecallToStartRecordingFor, func(l *livecall.Livecall) error {
		if ok := h.policy.Can(ctx, "StartRecording", userID, elevationToken, *l); !ok {
			return ErrNotAuthorizedToStartRecording
		}
		if err := l.StartRecording(h.commMgr); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}