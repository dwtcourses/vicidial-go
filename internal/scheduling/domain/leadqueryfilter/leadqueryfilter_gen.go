// Code generated by ddd-domain-gen, DO NOT EDIT.
package leadqueryfilter

import "fmt"

// Constructors ...

// New returns a guaranteed-to-be-valid LeadQueryFilter or an error
func New(filter string) (*LeadQueryFilter, error) {
	l := &LeadQueryFilter{filter: filter}
	return l, nil
}

// MustNew returns a guaranteed-to-be-valid LeadQueryFilter or panics
func MustNew(filter string) *LeadQueryFilter {
	l, err := New(filter)
	if err != nil {
		panic(err)
	}
	return l
}

// Marshalers ...

// UnmarshalFromRepository unmarshals LeadQueryFilter from the repository so that non-constructable
// private fields can still be initialized from (private) repository state
//
// Important: DO NEVER USE THIS METHOD EXCEPT FROM THE REPOSITORY
// Reason: This method initializes private state, so you could corrupt the domain.
func UnmarshalFromRepository(filter string) *LeadQueryFilter {
	l := MustNew(filter)
	return l
}

// Accessors ...

// Utilities ...

// Equal answers whether v is equivalent to l
// Always returns false if v is not a LeadQueryFilter
func (l LeadQueryFilter) Equal(v interface{}) bool {
	other, ok := v.(LeadQueryFilter)
	if !ok {
		return false
	}
	return l
}

// String implements the fmt.Stringer interface and returns the native format of LeadQueryFilter
func (l LeadQueryFilter) String() string {
	return fmt.Sprintf("")
}
