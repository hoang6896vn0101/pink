package modelrepos

import (
	e "pink/domains/entities"
	u "pink/domains/models"
)

// Weather model
type Weather u.WeatheModel

// WatherModelRepo interface
type WatherModelRepo interface {
	GetAllWeather() []UserModel
}

// GetAllWeather func
func (u Weather) GetAllWeather() []Weather {
	season := e.SeasonEntity{}
	sys := e.SysEntity{}
	wind := e.WindEntity{}
	weathers := []Weather{{Season: season, Sys: sys, Wind: wind}}
	return weathers
}
