package models

import "gorm.io/datatypes"

type User struct {
	ID             uint `gorm:"primaryKey"`
	Username       string
	Name           string
	HashedPassword string
	FightsWon      int
	Village        Village
}

type Village struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint
	VillageLevel int
	Gold         int
	Oil          int
	Money        int
	FarmLand     int
	Mines        int
	Map          datatypes.JSON
	Defenses     []Defenses
	Troops       []Troops
	Buildings    []Buildings
}

type Defenses struct {
	ID             uint `gorm:"primaryKey"`
	VillageID      uint
	Type           string
	DefensivePower int
	Amount         int
}

type Troops struct {
	ID             uint `gorm:"primaryKey"`
	VillageID      uint
	Type           string
	Health         int
	OffensivePower int
	Level          int
	Quantity       int
	Attacks        datatypes.JSON
}

type Buildings struct {
	ID            uint `gorm:"primaryKey"`
	VillageID     uint
	BuildingLevel int
	X             int
	Y             int
	Width         int
	Height        int
}

type BattleReplay struct {
	ID              uint `gorm:"primaryKey"`
	WinnerVillageID uint
	LoserVillageID  uint
	WinnerVillage   Village `gorm:"foreignKey:WinnerVillageID"`
	LoserVillage    Village `gorm:"foreignKey:LoserVillageID"`
	ReplayDefenses  []ReplayDefenses
	ReplayTroops    []ReplayTroops
}

type ReplayDefenses struct {
	ID             uint `gorm:"primaryKey"`
	BattleReplayID uint
	VillageID      uint
	Type           string
	DefensivePower int
	AmountDeployed int
	AttacksUsed    int
}

type ReplayTroops struct {
	ID               uint `gorm:"primaryKey"`
	BattleReplayID   uint
	VillageID        uint
	Type             string
	Health           int
	OffensivePower   int
	Level            int
	AmountDeployed   int
	AttacksAllocated int
}
