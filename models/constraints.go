package models

type VillageConstraints struct {
	MaxBuildings int
}

var LCons = []VillageConstraints{
	{MaxBuildings: 10},
	{MaxBuildings: 15},
	{MaxBuildings: 20},
	{MaxBuildings: 30},
}
