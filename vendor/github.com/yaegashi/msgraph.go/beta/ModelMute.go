// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MuteParticipantOperation undocumented
type MuteParticipantOperation struct {
	// CommsOperation is the base model of MuteParticipantOperation
	CommsOperation
}

// MuteParticipantsOperation undocumented
type MuteParticipantsOperation struct {
	// CommsOperation is the base model of MuteParticipantsOperation
	CommsOperation
	// Participants undocumented
	Participants []string `json:"participants,omitempty"`
}
