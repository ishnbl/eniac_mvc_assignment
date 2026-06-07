package models

import "gorm.io/datatypes"

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Username       string `gorm:"uniqueIndex"`
	Name           string
	HashedPassword string
	FightsWon      int
	Village        Village
}

type Village struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint `gorm:"uniqueIndex"`
	VillageLevel     int
	Gold             int
	Oil              int
	Money            int
	FarmLand         int
	Mines            int
	LevelConstraints datatypes.JSON
	Buildings        []BuildingsVillageMapping
	Defenses         []VillageDefMapping
	Troops           []TroopVillageMapping
}

type Defenses struct {
	ID             uint   `gorm:"primaryKey"`
	Type           string `gorm:"uniqueIndex"`
	DefensivePower int
	Capabilities   datatypes.JSON
	AttackPower    int
	Cost           int
}

type Troops struct {
	ID             uint `gorm:"primaryKey"`
	Type           string
	Health         int
	OffensivePower int
	Capabilities   datatypes.JSON
	Level          int
	Cost           int
}

type Buildings struct {
	ID     uint `gorm:"primaryKey"`
	Width  int
	Height int
	Type   string `gorm:"uniqueIndex"`
	Cost   int
}

type BattleReplay struct {
	ID             uint `gorm:"primaryKey"`
	AttackerID     uint
	DefenderID     uint
	AttackerLoot   int
	DefenderLoot   int
	Winner         bool
	Attacker       Village
	Defender       Village
	ReplayDefenses []ReplayDefenses
	ReplayTroops   []ReplayTroops
}

type ReplayDefenses struct {
	ID             uint `gorm:"primaryKey"`
	AttacksUsed    int
	BattleReplayID uint
	DefensesID     uint
	NumDeployed    int
	Defenses       Defenses
}

type ReplayTroops struct {
	ID               uint `gorm:"primaryKey"`
	AmountDeployed   int
	AttacksAllocated int
	BattleReplayID   uint
	TroopsID         uint
	Troops           Troops
}

type VillageDefMapping struct {
	ID         uint `gorm:"primaryKey"`
	DefensesID uint
	Amount     int
	VillageID  uint
	Defenses   Defenses
}

type TroopVillageMapping struct {
	ID        uint `gorm:"primaryKey"`
	VillageID uint
	TroopsID  uint
	Quantity  int
	Troops    Troops
}

type BuildingsVillageMapping struct {
	ID          uint `gorm:"primaryKey"`
	VillageID   uint
	X           int
	Y           int
	BuildingsID uint
	Buildings   Buildings
}
