package structure

import (
	"time"
)

type IdeaState int

const (
	IdeaStateUnknown IdeaState = iota
	IdeaStateRed
	IdeaStateYellow
	IdeaStateGreen
)

type Idea struct {
	Id          int64     `xorm:"pk autoincr" json:"id"`
	Owner       int64     `json:"owner"`
	Author      string    `xorm:"-" json:"author"`
	State       IdeaState `json:"state"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Notes       string    `json:"notes"`
	ManualUrl   string    `json:"manual_url"`
	CreatedAt   time.Time `xorm:"created" json:"-"`
	UpdatedAt   time.Time `xorm:"updated" json:"-"`
}
