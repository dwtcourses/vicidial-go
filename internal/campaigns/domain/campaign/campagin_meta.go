// GENERATED BY metatag, DO NOT EDIT
// (or edit away - I'm a comment, not a cop)

package campaign

import (
	"github.com/xoe-labs/vicidial-go/internal/campaigns/domain/lead"
)

// FilterLeads returns a copy of leads, omitting elements that are rejected by the given function.
func (c Campaign) FilterLeads(fn func(lead.Lead) bool) []lead.Lead {
	return c.FilterLeadsN(fn, -1)
}

// FilterLeadsN returns a copy of leads, omitting elements that are rejected by the given function.
// The n argument determines the maximum number of elements to return (n < 1: all elements).
func (c Campaign) FilterLeadsN(fn func(lead.Lead) bool, n int) []lead.Lead {
	cap := n
	if n < 1 {
		cap = len(c.leads)
	}
	result := make([]lead.Lead, 0, cap)
	for i := range c.leads {
		if fn(c.leads[i]) {
			if result = append(result, c.leads[i]); len(result) >= cap {
				break
			}
		}
	}
	return result
}
