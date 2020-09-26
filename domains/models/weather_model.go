package models

import e "pink/domains/entities"

type WeatheModel struct {
	Season e.SeasonEntity
	Sys    e.SysEntity
	Wind   e.WindEntity
}
