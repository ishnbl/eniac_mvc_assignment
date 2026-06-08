package models

import (
	"gorm.io/datatypes"
)

type retTroops struct {
	Quantity       int
	Type           string
	Health         int
	OffensivePower int
	Capabilities   datatypes.JSON
	Level          int
}

func GetTroops(username string) []retTroops {
	var user User
	var village Village

	var TroopMappings []TroopVillageMapping
	var returnTroops []retTroops
	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)
	DB.Where(&TroopVillageMapping{VillageID: village.ID}).Find(&TroopMappings)
	for i := 0; i < len(TroopMappings); i++ {
		var troop Troops
		DB.Where(&Troops{ID: TroopMappings[i].TroopsID}).First(&troop)
		returnTroops = append(returnTroops, retTroops{
			Quantity:       TroopMappings[i].Quantity,
			Type:           troop.Type,
			Health:         troop.Health,
			OffensivePower: troop.OffensivePower,
			Capabilities:   troop.Capabilities,
			Level:          troop.Level,
		})
	}
	return returnTroops
}

type retDefenses struct {
	Amount         int
	Type           string
	DefensivePower int
	AttackPower    int
	Capabilities   datatypes.JSON
}

func GetDefenses(username string) []retDefenses {
	var user User
	var village Village

	var DefMappings []VillageDefMapping
	var returnDefenses []retDefenses
	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)
	DB.Where(&VillageDefMapping{VillageID: village.ID}).Find(&DefMappings)
	for i := 0; i < len(DefMappings); i++ {
		var defense Defenses
		DB.Where(&Defenses{ID: DefMappings[i].DefensesID}).First(&defense)
		returnDefenses = append(returnDefenses, retDefenses{
			Amount:         DefMappings[i].Amount,
			Type:           defense.Type,
			DefensivePower: defense.DefensivePower,
			AttackPower:    defense.AttackPower,
			Capabilities:   defense.Capabilities,
		})
	}
	return returnDefenses
}

type RetBuilding struct {
	X      int
	Y      int
	Type   string
	Width  int
	Height int
	Cost   int
}

func GetBuildings(username string) []RetBuilding {
	var user User
	var village Village

	var BuildingMappings []BuildingsVillageMapping
	var returnBuildings []RetBuilding
	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)
	DB.Where(&BuildingsVillageMapping{VillageID: village.ID}).Find(&BuildingMappings)
	for i := 0; i < len(BuildingMappings); i++ {
		var building Buildings
		DB.Where(&Buildings{ID: BuildingMappings[i].BuildingsID}).First(&building)
		returnBuildings = append(returnBuildings, RetBuilding{
			X:      BuildingMappings[i].X,
			Y:      BuildingMappings[i].Y,
			Type:   building.Type,
			Width:  building.Width,
			Height: building.Height,
			Cost:   building.Cost,
		})
	}
	return returnBuildings
}

func GetVillage(username string) Village {
	var user User
	var village Village
	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)
	return village
}

type ReturnVillage struct {
	VillageReturned   Village
	BuildingsReturned []RetBuilding
}

func GetVillageBuildings(username string) ReturnVillage {
	var user User
	var village Village
	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)
	return_buildings := GetBuildings(username)
	return ReturnVillage{
		VillageReturned:   village,
		BuildingsReturned: return_buildings,
	}
}

func GetShopTroops(username string) []Troops {
	var user User
	var village Village
	var troops []Troops
	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)
	DB.Where("level <= ?", village.VillageLevel).Find(&troops)
	return troops
}

func GetShopDefenses(username string) []Defenses {
	var defenses []Defenses
	DB.Find(&defenses)
	return defenses
}

func CreateTroop(username string, payloadType string, payloadQuantity int) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	var troopType Troops

	DB.Where(&Troops{Type: payloadType}).First(&troopType)
	if troopType.ID == 0 {
		return false
	}
	if troopType.Cost > village.Money {
		return false
	}

	saveTroop := TroopVillageMapping{
		VillageID: village.ID,
		TroopsID:  troopType.ID,
		Quantity:  payloadQuantity,
	}

	DB.Create(&saveTroop)
	village.Money = village.Money - troopType.Cost
	DB.Save(&village)
	return true
}

func CreateDefense(username string, payloadType string, payloadAmount int) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	var defenseType Defenses

	DB.Where(&Defenses{Type: payloadType}).First(&defenseType)

	if defenseType.ID == 0 {
		return false
	}
	defense := VillageDefMapping{
		VillageID:  village.ID,
		Amount:     payloadAmount,
		DefensesID: defenseType.ID,
	}

	if defenseType.Cost > village.Money {
		return false
	}

	village.Money = village.Money - defenseType.Cost
	result := DB.Create(&defense)
	if result.Error != nil {
		return false
	}
	DB.Save(&village)
	return true
}

func CreateBuilding(username string, payloadType string, payloadX int, payloadY int) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	var buildingType Buildings
	DB.Where(&Buildings{Type: payloadType}).First(&buildingType)
	if buildingType.ID == 0 {
		return false
	}
	if buildingType.Cost > village.Money {
		return false
	}
	village.Money -= buildingType.Cost
	building := BuildingsVillageMapping{
		VillageID:   village.ID,
		BuildingsID: buildingType.ID,
		X:           payloadX,
		Y:           payloadY,
	}

	buildings := []BuildingsVillageMapping{}
	DB.Where(&BuildingsVillageMapping{VillageID: village.ID}).Find(&buildings)

	for _, b := range buildings {
		if (payloadX < b.X+buildingType.Width) && (b.X < payloadX+buildingType.Width) &&
			(payloadY < b.Y+buildingType.Height) && (b.Y < payloadY+buildingType.Height) {
			return false
		}
	}

	DB.Save(&village)
	result := DB.Create(&building)
	if result.Error != nil {
		return false
	}
	return true
}
