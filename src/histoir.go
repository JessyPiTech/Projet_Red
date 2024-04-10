package main

import (
	"fmt"
	"os"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
func combatRound(joueur *player, gobelin *Monstre) {
	charTurn(joueur, gobelin)
	fmt.Println(" ")
	if gobelin.pvac <= 0 {
		fmt.Println("Vous avez vaincu le Gobelin d'entraînement !")
		gainExperience(joueur, 100)
		return
	}
	ennemiPattern(*gobelin, joueur)
	if joueur.pvac <= 0 {
		fmt.Println("Le Gobelin d'entraînement vous a vaincu.")
		return
	}
}

// Tâche 17 : Fighter Squad
// Tâche 20 : Duel
// Mission 3 : Initiative
func trainingFight(joueur player) {
	clearScreen()
	fmt.Println("Combat d'entraînement contre un Gobelin d'entraînement !")
	gobelin := InitGoblin()
	tour := 1

	for {
		clearScreen()
		fmt.Printf("\nTour %d\n", tour)
		manaTours(&joueur)
		if joueur.initiative > gobelin.initiative {
			combatRound(&joueur, &gobelin)
		} else {
			ennemiPattern(gobelin, &joueur)
			fmt.Println(" ")
			if joueur.pvac <= 0 {
				fmt.Println("Le Gobelin d'entraînement vous a vaincu.")
				return
			}
			charTurn(&joueur, &gobelin)
			if gobelin.pvac <= 0 {
				fmt.Println("Vous avez vaincu le Gobelin d'entraînement !")
				gainExperience(&joueur, 100)
				return
			}
		}
		tour++
	}
}

// Tâche 19 : Ready Player One
// Fonction pour simuler le tour du joueur
// Fonction pour gérer le tour du personnage
func charTurn(joueur *player, ennemi *Monstre) {
	fmt.Println("Tour de ", joueur.name, ":")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	var choix int
	fmt.Println(" ")
	fmt.Print("Choisissez une action : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		attaque(joueur, ennemi)
	case 2:
		inventoryFight(joueur, ennemi)
	default:
		fmt.Println("Choix invalide.")
	}

}

// Bonus 1 : Améliorer le jeu
func debut(joueurs []player) {
	delaysup()
	histoirp1(joueurs[0])
	delaysup()
	choix1(&joueurs[0])
	delaysup()
	histoirp2()
	delaysup()
	choix2(&joueurs[0])
	delaysup()
	histoirp3(&joueurs[0])
	delaysup()
	Fight1(joueurs[0])

	a := 0
	menu3(joueurs, a)

}

func histoirp1(joueurs player) {
	fmt.Println("depuis votre plus jeune age on vous a raconte des histoirs")
	delaysup()
	fmt.Println("racontant les avantures de heros partent vers l'inconnu")
	delaysup()
	fmt.Println("mais aujourd'hui c'est a vous d'ecrir votre histoire")
	delaysup()
	fmt.Println("apres avoir preparer vostre sac a dos")
	delaysup()
	fmt.Println("vous sésice une carte")
	delaysup()
	fmt.Println("et place votre doigt en plein sur sont centre")
	delaysup()
	fmt.Println("sur la terre d'ou sont originaire toute les histoire")
	delaysup()
	fmt.Println("Les plaine d'opales")
	delaysup()
	fmt.Println("vous marchez pendant deux jours")
	delaysup()
	anim3()
	delaysup()
	fmt.Println("jusqu'a un petit vilage du nom Arbat")
	delaysup()
	fmt.Println("vous tombez sur un chevalier qui vous interpelle")
	delaysup()
	fmt.Println("-jeune homme")
	delaysup()
	fmt.Println("-tu a l'aire fort peu tu m'aide ?")
	delaysup()
}

func choix1(joueur *player) {
	clearScreen()
	choice := makeChoice()

	switch choice {
	case 1:
		fmt.Println("vous aider l'homme")
		return
	case 2:
		fmt.Println("c'est pas serieux c'est logique qui va passer un truc en plus si vous l'aider donc go fais 1")
		handleSecondChoice()
	default:
		fmt.Println("Choix invalide.")
	}
}
func makeChoice() int {
	for {
		fmt.Println("1. l'aide")
		fmt.Println("2. continué votre chemin")
		var choice int
		fmt.Print("Choisissez une option : ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erreur de saisie.")
			continue
		}
		return choice
	}
}

func handleSecondChoice() {
	for {
		choice := makeChoice()

		switch choice {
		case 1:
			fmt.Println("bien bon choix")
			delaysup()
			fmt.Println("vous aider l'homme")
			delaysup()
			return
		case 2:
			fmt.Println("non mais t'est serieux ?")
			delaysup()
			fmt.Println("bas tu c'est quoi")
			delaysup()
			fmt.Println("Vous êtes mort spontanement d'une crise cardiaque")
			delaysup()
			fmt.Println("DEAD")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func histoirp2() {
	fmt.Println("-merci beaucoup mon garçon")
	delaysup()
	fmt.Println("-viens je vais te payer un coup")
	delaysup()
	fmt.Println("-bon pour te remercier je te propose 2 chose")
	delaysup()
}

func choix2(joueur *player) {
	for {
		fmt.Println("1. soit je t'entraine au combat d'épée")
		fmt.Println("2. soit je t'offre 5 d'or")
		var choice int
		fmt.Print("Choisissez une option : ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erreur de saisie.")
			continue
		}
		switch choice {
		case 1:
			fmt.Println("d'accord suit moi")
			return
		case 2:
			fmt.Println("bon et bien voila ton solde")
			a := 5
			addMoney(*joueur, a)
			delaysup()
			fmt.Println("bon en vrai vue que le but en vrai c'est d'avancer dans l'histoir on va faire autrement")
			delaysup()
			fmt.Println("Alors que vous vous appretiez a accepter l'argent")
			delaysup()
			fmt.Println("une voit s'ecrillat,", joueur.name, "!!!!!")
			delaysup()
			fmt.Println("tu dois apprendre a combatre !!!")
			delaysup()
			fmt.Println("vous decidâme alors de vous entrenez avec le chevalier")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func histoirp3(joueur *player) {
	fmt.Println("prend cet épee, et affronte moi")
	StufToAdd := newItem("épee d'entrainemnt", "Arme", 1, 10, 15, 1)
	addEquipment(joueur, StufToAdd, &joueur.equipement[0].Arme)
	joueur.Skills = append(joueur.Skills, newSkills("Arme principale", 1, joueur.equipement[0].Arme[0].point, 10))
}
func prepareCombat(joueur *player, chevalier *Monstre) {
	clearScreen()
	fmt.Println("Préparez-vous")
	joueur.Mana = joueur.ManaMax
	fmt.Printf("Mana : %d/%d\n", joueur.Mana, joueur.ManaMax)
}

func handleCombatOutcome(joueur *player, chevalier *Monstre, tour int) {
	if chevalier.pvac <= 0 {
		clearScreen()
		fmt.Println("- Bien joué jeune padawan, vous êtes désormais un maître d'arme.")
		delaysup()
		gainExperience(joueur, 120)
		delaysup()
		fmt.Println("Vous avez appris l'attaque avec l'arme principale.")
		delaysup()
		fmt.Println("- Gardez cette épée, je vous l'offre.")
		delaysup()
		fmt.Println("Vous avez reçu une arme d'entraînement.")
		delaysup()
		fmt.Println("Si vous souhaitez infliger davantage de dégâts, équipez une nouvelle arme depuis votre inventaire.")
		delaysup()
		return
		fmt.Println("1111111")
	}
	if joueur.pvac <= 10 {
		fmt.Println("Alors que vous vous rapprochez de plus en plus du sol...")
		delaysup()
		fmt.Println("Le chevalier s'exclama :")
		delaysup()
		fmt.Println("- Ma cheville !")
		delaysup()
		fmt.Println("Dans votre chute, votre épée tomba au sol.")
		delaysup()
		fmt.Println("Et alors qu'il préparait son coup final...")
		delaysup()
		fmt.Println("Il se rétamait comme une merde.")
		delaysup()
		chevalier.pvac = 0
	}

}

func Fight1(joueur player) {
	chevalier := InitChevalier()
	prepareCombat(&joueur, &chevalier)
	tour := 1

	for {
		clearScreen()
		fmt.Printf("\nTour %d\n", tour)
		manaTours(&joueur)
		if joueur.initiative > chevalier.initiative {
			fmt.Println(" ")
			charTurn(&joueur, &chevalier)
			handleCombatOutcome(&joueur, &chevalier, tour)
			ennemiPattern(chevalier, &joueur)
			dead(&joueur)
			handleCombatOutcome(&joueur, &chevalier, tour)
			fmt.Println(" ")

		} else {
			fmt.Println(" ")
			ennemiPattern(chevalier, &joueur)
			fmt.Println(" ")
			dead(&joueur)
			handleCombatOutcome(&joueur, &chevalier, tour)

			fmt.Println(" ")
			charTurn(&joueur, &chevalier)
			handleCombatOutcome(&joueur, &chevalier, tour)
		}
		if chevalier.pvac <= 0 {
			break
		}
		tour++
	}
	return
}
