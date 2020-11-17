// GENERATED BY metatag, DO NOT EDIT
// (or edit away - I'm a comment, not a cop)

package livecall

import (
	"fmt"
	"github.com/xoe-labs/vicidial-go/internal/common/party"
)

// Equal answers whether v is equivalent to l.
// Always returns false if v is not a Livecall.
func (l Livecall) Equal(v interface{}) bool {
	l2, ok := v.(Livecall)
	if !ok {
		return false
	}
	if l.uuid != l2.uuid {
		return false
	}
	return true
}

// Lead returns the value of lead.
func (l Livecall) Lead() party.RemoteParty {
	return l.lead
}

// String returns the "native" format of Livecall. Implements the fmt.Stringer interface.
func (l Livecall) String() string {
	return fmt.Sprintf("%v %v", l.lead, l.localParty)
}

// LocalParty returns the value of localParty.
func (l Livecall) LocalParty() party.LocalParty {
	return l.localParty
}

// SetResultSentinel sets the given value as resultSentinel.
func (l *Livecall) SetResultSentinel(s string) {
	l.resultSentinel = s
}