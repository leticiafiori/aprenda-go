package transfer

import "time"

type Transfer struct {
	id                     int
	account_origin_id      string
	account_destination_id string
	amount                 float64
	created_at             time.Time
}
