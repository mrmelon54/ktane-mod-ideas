package structure

import "time"

type User struct {
	Id                   int64     `xorm:"pk autoincr" json:"id"`
	DiscordId            string    `json:"-"`
	DiscordName          string    `json:"discord_name,omitempty"`
	DiscordDiscriminator string    `json:"discord_discriminator,omitempty"`
	Picture              string    `json:"picture,omitempty"`
	Banned               bool      `json:"banned,omitempty"`
	Admin                bool      `json:"admin,omitempty"`
	CreatedAt            time.Time `xorm:"created index" json:"-"`
	UpdatedAt            time.Time `xorm:"updated index" json:"-"`
}
