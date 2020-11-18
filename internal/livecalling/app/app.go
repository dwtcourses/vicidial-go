// Copyright © 2020 David Arnold <dar@xoe.solutions>
// SPDX-License-Identifier: MIT

package app

import (
	"github.com/xoe-labs/vicidial-go/internal/livecalling/app/command"
	"github.com/xoe-labs/vicidial-go/internal/livecalling/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	HangUp command.HangUpHandler

	StartRecording command.StartRecordingHandler
	StopRecording  command.StopRecordingHandler

	PlayAudioToBothParties command.PlayAudioToBothPartiesHandler
	PlayAudioToRemoteParty command.PlayAudioToRemotePartyHandler
	PlayAudioToLocalParty  command.PlayAudioToLocalPartyHandler
}

type Queries struct {
	// HourAvailability      query.HourAvailabilityHandler
	// TrainerAvailableHours query.AvailableHoursHandler
}