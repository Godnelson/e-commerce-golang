package costumer

import "time"

type Costumer struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
