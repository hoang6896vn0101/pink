package model_repos

import (
	e "pink/domains/entities"
	u "pink/domains/models"
)

type Weather u.WeatheModel

type WatherModelRepo interface {
	GetAllWeather() []UserModel
}

func (u Weather) GetAllWeather() []Weather {
	season := e.SeasonEntity{}
	sys := e.SysEntity{}
	wind := e.WindEntity{}
	weathers := []Weather{{Season: season, Sys: sys, Wind: wind}}
	return weathers
}
