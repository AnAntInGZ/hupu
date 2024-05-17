package dao

import "alex.tse/model"

func ListRaceTypes() ([]model.RaceType, error) {
	db := GetDBInstance()
	var raceTypes []model.RaceType
	if err := db.Model(&model.RaceType{}).Find(&raceTypes).Error; err != nil {
		return nil, err
	}
	return raceTypes, nil
}
