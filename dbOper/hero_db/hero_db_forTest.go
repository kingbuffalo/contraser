package hero_db

func GetHeroVOsForTest1() [](*HeroVO) {
	heroVO1 := &HeroVO{
		PlayerId: -1,
		Id:       1,
		Exp:      1,
		Level:    1,
		Star:     1,
		ArmyId:   1,
	}
	return []*HeroVO{
		heroVO1,
	}
}
