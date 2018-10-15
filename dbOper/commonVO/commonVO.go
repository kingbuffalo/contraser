package commonVO

type RobAbleResVO struct {
	Wood  int32 `json:"wood"`
	Gold  int32 `json:"gold"`
	Grain int32 `json:"grain"`
}

type MapNpcHeroVO struct {
	HeroId int32 `json:"hero_id"`
	Troops int32 `json:"troops"`
	Level  int32 `json:"level"`
}
