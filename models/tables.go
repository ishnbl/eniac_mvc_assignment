package models

import (
	"gorm.io/datatypes"
)

type User struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	username        string `gorm:"unique"`
	hashed_password string
	VillageID       uint
	Village         Village
}

type Village struct {
	ID            uint `gorm:"primaryKey"`
	resources     uint
	village_level uint
	vill_map      datatypes.JSON
	yield         uint
	DefensesID    uint
	TroopsID      uint
	Defenses      Defenses
	Troops        Troops
}

type Defenses struct {
	ID               uint `gorm:"primaryKey"`
	water_crocodiles uint
	wall_archers     uint
	CastleCannons    uint
}

type Troops struct {
	ID        uint `gorm:"primaryKey"`
	archers   uint
	swordsmen uint
	beasts    uint
	ninjas    uint
	cavalry   uint
}

type Fights struct {
	ID       uint `gorm:"primaryKey"`
	WinnerID uint
	LooserID uint

	Winner User
	Looser User
}

type Buildings struct {
	ID             uint `gorm:"primaryKey"`
	building_level uint
	x              uint
	y              uint
	height         uint
	width          uint
	VillageID      uint
	Village        Village
}
