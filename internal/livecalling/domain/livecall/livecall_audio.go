// Copyright © 2020 David Arnold <dar@xoe.solutions>
// SPDX-License-Identifier: MIT

package livecall

import (
	cerrors "github.com/xoe-labs/vicidial-go/internal/common/errors"
	"github.com/xoe-labs/vicidial-go/internal/livecalling/domain"
)

const (
	// ErrUnableToPlayAudio signals that the audio could not be played
	ErrUnableToPlayAudio = cerrors.CommonError("ErrUnableToPlayAudio")
)

// livecall play audio  domain topic
type livecallPlayAudio struct {
	// currently no persisted data for play audio topic
}

// PlayAudioToRemoteParty plays an audio through a communications manager to the remote party
func (l *Livecall) PlayAudioToRemoteParty(cm domain.CommManager) error {
	// TODO: implement play audio to remote party on communications manager
	return nil
}

// PlayAudioToLocalParty plays an audio through a communications manager to the local party
func (l *Livecall) PlayAudioToLocalParty(cm domain.CommManager) error {
	// TODO: implement play audio to local party on communications manager
	return nil
}

// PlayAudioToBothParties plays an audio through a communications manager to both parties
func (l *Livecall) PlayAudioToBothParties(cm domain.CommManager) error {
	// TODO: implement play audio to both parties on communications manager
	return nil
}
