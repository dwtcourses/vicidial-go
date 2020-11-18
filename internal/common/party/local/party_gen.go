// Code generated by ddd-domain-gen, DO NOT EDIT.
package local

import (
	"errors"
	"fmt"
	gouuid "github.com/satori/go.uuid"
	"reflect"
)

// Constructors ...

// New returns a guaranteed-to-be-valid LocalParty or an error
func New(uuid gouuid.UUID, name string, endpointUUID string, pType LocalPartyType) (*LocalParty, error) {
	if reflect.ValueOf(uuid).IsZero() {
		return nil, errors.New("missing party UUID")
	}
	if reflect.ValueOf(name).IsZero() {
		return nil, errors.New("missing party name")
	}
	if reflect.ValueOf(endpointUUID).IsZero() {
		return nil, errors.New("missing endpoint UUID")
	}
	if reflect.ValueOf(pType).IsZero() {
		return nil, errors.New("missing party type")
	}
	l := &LocalParty{
		endpointUUID: endpointUUID,
		name:         name,
		pType:        pType,
		uuid:         uuid,
	}
	return l, nil
}

// MustNew returns a guaranteed-to-be-valid LocalParty or panics
func MustNew(uuid gouuid.UUID, name string, endpointUUID string, pType LocalPartyType) *LocalParty {
	l, err := New(uuid, name, endpointUUID, pType)
	if err != nil {
		panic(err)
	}
	return l
}

// Marshalers ...

// UnmarshalFromRepository unmarshals LocalParty from the repository so that non-constructable
// private fields can still be initialized from (private) repository state
//
// Important: DO NEVER USE THIS METHOD EXCEPT FROM THE REPOSITORY
// Reason: This method initializes private state, so you could corrupt the domain.
func UnmarshalFromRepository(uuid gouuid.UUID, name string, endpointUUID string, pType LocalPartyType) *LocalParty {
	l := MustNew(uuid, name, endpointUUID, pType)
	return l
}

// Accessors ...

// Utilities ...

// Equal answers whether v is equivalent to l
// Always returns false if v is not a LocalParty
func (l LocalParty) Equal(v interface{}) bool {
	other, ok := v.(LocalParty)
	if !ok {
		return false
	}
	if l.uuid != other.uuid {
		return false
	}
	return l
}

// String implements the fmt.Stringer interface and returns the native format of LocalParty
func (l LocalParty) String() string {
	return fmt.Sprintf("%s %s ", l.name, l.pType)
}