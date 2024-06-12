package category

import "time"

type Category struct {
	Id         int       `db:"category_id" json:"id"`
	Name       string    `db:"name" json:"name"`
	LastUpdate time.Time `db:"last_update" json:"lastUpdate"`
}
