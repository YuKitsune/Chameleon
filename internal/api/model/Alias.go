package model

import (
	"gorm.io/gorm"
	"regexp"
)

type Alias struct {
	gorm.Model
	UserID                 uint
	Username               string
	SenderWhitelistPattern string
}

func (a *Alias) SenderIsAllowed(sender string) (*bool, error) {
	r, err := regexp.Compile(a.SenderWhitelistPattern)
	if err != nil {
		return nil, err
	}

	match := r.MatchString(sender)
	return &match, nil
}