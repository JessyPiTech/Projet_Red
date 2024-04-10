package main

import (
	"fmt"
)

type item struct {
	name     string
	tipe     string
	level    int
	point    int
	value    int
	Quantity int
}

// Tâche 14 : I saw it in the mirror
type equipment struct {
	Tete          []item
	Torse         []item
	Pieds         []item
	Arme          []item
	Bouclier      []item
	ArmeSecondair []item
}

type skills struct {
	name         string
	level        int
	degats       int
	manaDecresse int
}
type player struct {
	name          string
	class         string
	level         int
	pvmax         int
	pvac          int
	equipement    []equipment
	inventory     []item
	Mana          int
	ManaMax       int
	Experience    int
	ExperienceMax int
	money         int
	charge        int
	Skills        []skills
	initiative    int
}

type Monstre struct {
	Nom           string
	pvmax         int
	pvac          int
	PointsAttaque int
	initiative    int
}

// Tâche 10 : Wingardium leviosa
func newSkills(name string, level int, degats int, manaDecresse int) skills {
	return skills{
		name:         name,
		level:        level,
		degats:       degats,
		manaDecresse: manaDecresse,
	}
}

func newItem(name string, tipe string, level int, point int, value int, quantity int) item {
	return item{
		name:     name,
		tipe:     tipe,
		level:    level,
		point:    point,
		value:    value,
		Quantity: quantity,
	}
}

// Tâche 3 : Initialisation du personnage
// Tâche 11 : Money, money, money

func newPlayer(name string, class string) player {
	pvmax, pvac, initiative, charge := initializePlayerStats(class)

	// Initialize equipment slices for each type
	equipement := []equipment{
		{[]item{}, []item{}, []item{}, []item{}, []item{}, []item{}},
	}

	return player{
		name:          formatNom(name),
		class:         class,
		level:         1,
		pvmax:         pvmax,
		pvac:          pvac,
		equipement:    equipement,
		inventory:     []item{newItem("Potion of life", "Resource", 1, 40, 10, 1), newItem("Potion of poison", "Resource", 1, 40, 10, 1), newItem("Potion of mana", "Resource", 1, 40, 10, 1)},
		money:         100,
		ManaMax:       50,
		Mana:          50,
		Experience:    0,
		ExperienceMax: 100,
		charge:        charge,
		Skills:        []skills{newSkills("Coup de poing", 1, 8, 5)},
		initiative:    initiative,
	}
}

// Mission 1 : Amélioration de la création de personnage
// Mission 2.1 : Limite d’inventair (charge)
func initializePlayerStats(class string) (int, int, int, int) {
	var pvmax, pvac, initiative, charge int

	switch class {
	case "Human":
		pvmax = 100
		pvac = 50
		charge = 10
		initiative = 3
	case "Elf":
		pvmax = 80
		pvac = 40
		charge = 15
		initiative = 5
	case "Dwarf":
		pvmax = 120
		pvac = 60
		charge = 5
		initiative = 1
	default:
		pvmax = 100
		pvac = 50
		charge = 10
		initiative = 3
	}

	return pvmax, pvac, initiative, charge
}

// Tâche 16 : La Chose
func InitGoblin() Monstre {
	gobelin := Monstre{
		Nom:           "Gobelin",
		pvmax:         40,
		pvac:          40,
		PointsAttaque: 5,
		initiative:    5,
	}
	return gobelin
}

func InitChevalier() Monstre {
	chevalier := Monstre{
		Nom:           "chevalier",
		pvmax:         100,
		pvac:          100,
		PointsAttaque: 15,
		initiative:    2,
	}
	return chevalier
}

func main() {
	joueurs := []player{createPlayer()}
	fmt.Printf("Name : %s, Life : %d/%d, class : %s\n", joueurs[0].name, joueurs[0].pvac, joueurs[0].pvmax, joueurs[0].class)
	delaysup()

	menu(joueurs)
	a := 0
	menu3(joueurs, a)
	select {}
}
