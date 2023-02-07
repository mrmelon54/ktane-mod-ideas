package structure

import "time"

type Vote struct {
	Id        int64     `xorm:"pk autoincr" json:"id"`
	Owner     int64     `json:"owner"`
	Idea      int64     `json:"idea"`
	Value     int       `json:"value"`
	CreatedAt time.Time `xorm:"created index" json:"-"`
	UpdatedAt time.Time `xorm:"updated index" json:"-"`
}
