package schema

import "time"

type UserFilters struct {
	ID            string
	Username      string
	Email         string
	CreatedBefore time.Time
	CreatedAfter  time.Time
}
