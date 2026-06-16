package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/datatypes"
)

type RetTroops struct {
	Quantity       int
	Type           string
	Health         int
	OffensivePower int
	Level          int
	Skills         []Capability
}

func GetTroops(username string) []RetTroops {
	var user User
	var village Village

	var TroopMappings []TroopVillageMapping
	var returnTroops []RetTroops
	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)
	DB.Where(&TroopVillageMapping{VillageID: village.ID}).Find(&TroopMappings)
	for i := 0; i < len(TroopMappings); i++ {
		var troop Troops
		DB.Where(&Troops{ID: TroopMappings[i].TroopsID}).First(&troop)
		var skills []Capability
		json.Unmarshal(troop.Capabilities, &skills)
		returnTroops = append(returnTroops, RetTroops{
			Quantity:       TroopMappings[i].Quantity,
			Type:           troop.Type,
			Health:         troop.Health,
			OffensivePower: troop.OffensivePower,
			Level:          troop.Level,
			Skills:         skills,
		})
	}
	return returnTroops
}

type RetDefenses struct {
	Amount         int
	Type           string
	DefensivePower int
	AttackPower    int
	Capabilities   datatypes.JSON
}

func GetDefenses(username string) []RetDefenses {
	var user User
	var village Village

	var DefMappings []VillageDefMapping
	var returnDefenses []RetDefenses
	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)
	DB.Where(&VillageDefMapping{VillageID: village.ID}).Find(&DefMappings)
	for i := 0; i < len(DefMappings); i++ {
		var defense Defenses
		DB.Where(&Defenses{ID: DefMappings[i].DefensesID}).First(&defense)
		returnDefenses = append(returnDefenses, RetDefenses{
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
	ID     uint
	X      int
	Y      int
	Type   string
	Width  int
	Height int
	Level  int
}

func GetBuildings(username string) []RetBuilding {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	var buildings []Building
	var returnBuildings []RetBuilding

	DB.Where(&Building{VillageID: village.ID}).Find(&buildings)
	for i := 0; i < len(buildings); i++ {
		var levelSpec LevelSpecific
		DB.Where(&LevelSpecific{Level: buildings[i].Level}).First(&levelSpec)
		returnBuildings = append(returnBuildings, RetBuilding{
			X:      buildings[i].X,
			Y:      buildings[i].Y,
			Type:   buildings[i].BuildingType,
			Width:  levelSpec.W,
			Height: levelSpec.H,
			Level:  buildings[i].Level,
			ID:     buildings[i].ID,
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

type RetShopBuilding struct {
	Type   string
	Width  int
	Height int
	Cost   int
}

var bulType = []string{"Gold", "Oil", "Farm", "Armoury", "Barrack"}

func GetShopBuildings(username string) []RetShopBuilding {
	var levelSpec LevelSpecific
	DB.Where(&LevelSpecific{Level: 1}).First(&levelSpec)

	var returnBuildings []RetShopBuilding
	for _, bType := range bulType {
		returnBuildings = append(returnBuildings, RetShopBuilding{
			Type:   bType,
			Width:  levelSpec.W,
			Height: levelSpec.H,
			Cost:   levelSpec.UpgCost,
		})
	}
	return returnBuildings
}

func UpgradeVillage(username string) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	village.Money = village.Money - village.VillageLevel*1000
	if village.Money < 0 || village.VillageLevel > 3 {
		return false
	}
	village.VillageLevel++
	constraintsJSON, _ := json.Marshal(LCons[village.VillageLevel-1])
	village.LevelConstraints = datatypes.JSON(constraintsJSON)
	DB.Save(&village)
	return true
}

func CreateTroop(username string, payloadType string, payloadQuantity int, payloadLevel int) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	var troopType Troops
	DB.Where(&Troops{Type: payloadType, Level: payloadLevel}).First(&troopType)
	var troptot int
	mytroops := GetTroops(username)
	for i := 0; i < len(mytroops); i++ {
		troptot += mytroops[i].Quantity
	}
	buildings := GetBuildings(username)
	barrack := 0
	for i := 0; i < len(buildings); i++ {
		if buildings[i].Type == "Barrack" {
			barrack++
		}
	}
	if troopType.ID == 0 {
		return false
	}
	if troopType.Cost > village.Money || (troptot+payloadQuantity) > barrack*10 {
		return false
	}

	saveTroop := TroopVillageMapping{
		VillageID: village.ID,
		TroopsID:  troopType.ID,
		Quantity:  payloadQuantity,
	}

	DB.Create(&saveTroop)
	village.Money = village.Money - troopType.Cost*payloadQuantity
	DB.Save(&village)
	return true
}

func CreateDefense(username string, payloadType string, payloadAmount int) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	var defenseType Defenses
	deftot := 0
	mydef := GetDefenses(username)
	for i := 0; i < len(mydef); i++ {
		deftot += mydef[i].Amount
	}
	DB.Where(&Defenses{Type: payloadType}).First(&defenseType)
	buildings := GetBuildings(username)
	armoury := 0
	for i := 0; i < len(buildings); i++ {
		if buildings[i].Type == "Armoury" {
			armoury++
		}
	}

	if defenseType.ID == 0 {
		return false
	}
	defense := VillageDefMapping{
		VillageID:  village.ID,
		Amount:     payloadAmount,
		DefensesID: defenseType.ID,
	}

	if defenseType.Cost > village.Money || (deftot+payloadAmount) > armoury*10 {
		return false
	}

	village.Money = village.Money - defenseType.Cost*payloadAmount
	result := DB.Create(&defense)
	if result.Error != nil {
		return false
	}
	DB.Save(&village)
	return true
}

func ExtractResources(username string, buildingID uint) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	var building Building
	DB.Where(&Building{ID: buildingID, VillageID: village.ID}).First(&building)
	if building.ID == 0 {
		return false
	}

	var resColl ResourceColl
	DB.Where(&ResourceColl{Level: building.Level, Type: building.BuildingType}).First(&resColl)
	if resColl.Level == 0 || time.Since(building.ResourceCollected) < time.Duration(resColl.TimeRefill)*time.Second {
		return false
	}

	if building.BuildingType == "Gold" {
		village.Gold = village.Gold + resColl.ResourceYield
	} else if building.BuildingType == "Oil" {
		village.Oil = village.Oil + resColl.ResourceYield
	} else if building.BuildingType == "Farm" {
		village.Money = village.Money + int(float32(resColl.ResourceYield)*0.1)
	}

	building.ResourceCollected = time.Now()
	DB.Save(&building)
	DB.Save(&village)
	return true
}

func UpgradeTroop(username string, payloadType string) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	var mappings []TroopVillageMapping
	DB.Where(&TroopVillageMapping{VillageID: village.ID}).Find(&mappings)

	for i := 0; i < len(mappings); i++ {
		var troop Troops
		DB.Where(&Troops{ID: mappings[i].TroopsID}).First(&troop)
		if troop.Type == payloadType {
			var nextTroop Troops
			DB.Where(&Troops{Type: payloadType, Level: troop.Level + 1}).First(&nextTroop)
			if nextTroop.ID == 0 {
				return false
			}
			if nextTroop.Cost > village.Money {
				return false
			}
			village.Money -= nextTroop.Cost
			mappings[i].TroopsID = nextTroop.ID
			DB.Save(&mappings[i])
			DB.Save(&village)
			return true
		}
	}
	return false
}

func UpgradeBuilding(username string, buildingID uint) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	var building Building
	DB.Where(&Building{ID: buildingID}).First(&building)
	if building.ID == 0 {
		return false
	}

	var levelSpec LevelSpecific
	DB.Where(&LevelSpecific{Level: building.Level + 1}).First(&levelSpec)

	if levelSpec.MinVillLevel > village.VillageLevel || levelSpec.UpgCost > village.Money {
		return false
	}

	village.Money -= levelSpec.UpgCost
	building.Level++
	DB.Save(&village)
	DB.Save(&building)
	return true
}

func CreateBuilding(username string, payloadType string, payloadX int, payloadY int) bool {
	var user User
	var village Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&village)

	if payloadType != "Gold" && payloadType != "Oil" && payloadType != "Farm" && payloadType != "Armoury" && payloadType != "Barrack" {
		return false
	}

	buildings_ex := GetBuildings(username)
	var w1 int
	var h1 int
	var LevData LevelSpecific
	DB.Where(&LevelSpecific{Level: 1}).First(&LevData)
	w1 = LevData.W
	h1 = LevData.H

	fmt.Println(w1, h1)
	for _, b := range buildings_ex {
		if (payloadX < b.X+b.Width) && (b.X < payloadX+w1) &&
			(payloadY < b.Y+b.Height) && (b.Y < payloadY+h1) {
			return false
		}
	}
	if len(buildings_ex) > LCons[village.VillageLevel].MaxBuildings {
		return false
	}
	village.Money = village.Money - LevData.UpgCost
	if village.Money < 0 {
		return false
	}
	buildingSave := Building{
		X:            payloadX,
		Y:            payloadY,
		BuildingType: payloadType,
		Level:        1,
		VillageID:    village.ID,
	}
	DB.Save(&village)
	DB.Save(&buildingSave)
	return true
}

type RetScoutVillage struct {
	ID           uint
	Username     string
	VillageLevel int
}

func GetAllVillages(username string) []RetScoutVillage {
	var user User
	DB.Where(&User{Username: username}).First(&user)

	var villages []Village
	_ = DB.Find(&villages)
	var returnVillages []RetScoutVillage
	for i := 0; i < len(villages); i++ {
		var owner User
		DB.Where(&User{ID: villages[i].UserID}).First(&owner)
		returnVillages = append(returnVillages, RetScoutVillage{
			ID:           villages[i].ID,
			Username:     owner.Username,
			VillageLevel: villages[i].VillageLevel,
		})
	}
	return returnVillages
}

type FightConf struct {
	TypeTroop  string
	LevelTroop int
	AmountDep  int
	CapType    string
}

type Capability struct {
	CapType        string  `json:"CapType"`
	RelativeDamage float32 `json:"RelativeDamage"`
}

func CreateFight(username string, defenderID uint, config []FightConf) (bool, int) {
	var user User
	var attackerVillage Village
	var defenderVillage Village

	DB.Where(&User{Username: username}).First(&user)
	DB.Where(&Village{UserID: user.ID}).First(&attackerVillage)
	DB.Where(&Village{ID: defenderID}).First(&defenderVillage)

	attackTroop := 0
	heaTroop := 0

	var troopsUsed []Troops
	for i := 0; i < len(config); i++ {
		var troop Troops
		DB.Where(&Troops{Type: config[i].TypeTroop, Level: config[i].LevelTroop}).First(&troop)

		var caps []Capability
		json.Unmarshal(troop.Capabilities, &caps)
		var relDamage float32
		for i := 0; i < len(caps); i++ {
			if caps[i].CapType == config[i].CapType {
				relDamage = caps[i].RelativeDamage
				break
			}
		}

		attackTroop += int(float32(troop.OffensivePower) * relDamage * float32(config[i].AmountDep))
		heaTroop += troop.Health * config[i].AmountDep
		troopsUsed = append(troopsUsed, troop)
	}

	heaDefender := 0
	attackDef := 0

	var defMappings []VillageDefMapping
	DB.Where(&VillageDefMapping{VillageID: defenderVillage.ID}).Find(&defMappings)
	var defensesUsed []Defenses
	for i := 0; i < len(defMappings); i++ {
		var defense Defenses
		DB.Where(&Defenses{ID: defMappings[i].DefensesID}).First(&defense)
		heaDefender += defense.DefensivePower * defMappings[i].Amount
		attackDef += defense.AttackPower * defMappings[i].Amount
		defensesUsed = append(defensesUsed, defense)
	}

	winner := attackTroop > heaDefender

	loot := 0
	if winner == true {
		goldLoot := defenderVillage.Gold / 10
		oilLoot := defenderVillage.Oil / 10
		loot = goldLoot + oilLoot
		attackerVillage.Gold += goldLoot
		attackerVillage.Oil += oilLoot
		defenderVillage.Gold -= goldLoot
		defenderVillage.Oil -= oilLoot
		DB.Save(&attackerVillage)
		DB.Save(&defenderVillage)
		user.FightsWon++
		DB.Save(&user)
	}

	return winner, loot
}
