package main

import (
	"fmt"
	"os"
	"time"
)

// Tâche 1 : Création du menu
// menu principal avant tuto
func menu(joueurs []player) {
	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("Main Menu")
		fmt.Println("1. Show your character")
		fmt.Println("2. Inventory")
		fmt.Println("3. Go to the training fight")
		fmt.Println("4. Exit")
		var choice int
		fmt.Print("Choose an option : ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}
		switch choice {
		case 1:
			subMenu1(joueurs[0])
		case 2:
			subMenu2(joueurs[0])
		case 3:
			c := 0
			lancement(joueurs, c)
		case 4:
			fmt.Println("See you next time !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}

	}
}

// function pour evite 30 ligne
func subMenu1(joueur player) {
	clearScreen()
	fmt.Println("Here is your character !")
	displayInfo(joueur)
	fmt.Println("Press enter to go back to the menu")
	var choice int
	fmt.Scanf("%d", &choice)
	clearScreen()
}

// function pour evite 30 ligne
func subMenu2(joueur player) {
	clearScreen()
	fmt.Println("Here is your inventory !")
	displayInventory(joueur)
	fmt.Println("Press enter to go back to the menu")
	var choice int
	fmt.Scanf("%d", &choice)
	clearScreen()
}

// Tâche 13 : Gimme! Gimme! Gimme!
// submenu3 pour les 30 lignes
func subMenu3(joueurs []player) {
	for {
		clearScreen()
		Csolde := &joueurs[0].money
		fmt.Println("Welcome to the black-smith, current solde ", *Csolde, "\n")
		fmt.Println("\n1. Hat of the adventurer -> required : 2 feather crow , 1 acacia wood")
		fmt.Println("2. coat of the adventurer -> required : 2 wolf fur, 1 troll skin")
		fmt.Println("3. Boots of the adventurer -> required : 1 wolf fur, 1 boar fur")
		//fmt.Println("4. Sell Items")
		fmt.Println("4. Go back to the main menu")
		fmt.Println("What would you like to build ?")
		var choice int
		fmt.Println("Choose an option")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erreur de saisie.")
			continue
		}
		switch choice {
		case 1:
			requiredItem := []item{
				{name: "Feather crow", tipe: "Resource", level: 1, point: 2, value: 3, Quantity: 2},
				{name: "Acacia wood", tipe: "Resource", level: 1, point: 2, value: 3, Quantity: 1},
			}
			if checkInventory(&joueurs[0], requiredItem) {
				// Supprimez "Feather crow" de l'inventaire
				removeInventory(&joueurs[0], "Feather crow", requiredItem[0].Quantity)
				fmt.Println("You found Feather crow!")

				// Supprimez "Acacia wood" de l'inventaire
				removeInventory(&joueurs[0], "Acacia wood", requiredItem[1].Quantity)
				fmt.Println("You found Acacia wood!")
				delaysup()

				fmt.Println("Great you bought hat of the adventurer")
				*Csolde -= 5
				itemToAdd := newItem("Hat of the adventurer", "Tete", 1, 20, 3, 1)
				addInventory(&joueurs[0], itemToAdd)
			} else {
				fmt.Println("You dont have the required Items")
				time.Sleep(3 * time.Second)
			}
		case 2:
			requiredItem := []item{
				{name: "Wolf fur", tipe: "Resource", level: 1, point: 2, value: 3, Quantity: 2},
				{name: "Troll skin", tipe: "Resource", level: 1, point: 2, value: 3, Quantity: 1},
			}
			if checkInventory(&joueurs[0], requiredItem) {
				removeInventory(&joueurs[0], "Wolf fur", requiredItem[0].Quantity)
				removeInventory(&joueurs[0], "Troll skin", requiredItem[1].Quantity)
				fmt.Println("Great you bought coat of the adventurer")
				*Csolde -= 5
				itemToAdd := newItem("Coat of the adventurer", "Torse", 1, 30, 3, 1)
				addInventory(&joueurs[0], itemToAdd)
			} else {
				fmt.Println("You dont have the required Items")
				time.Sleep(3 * time.Second)
			}
		case 3:
			requiredItem := []item{
				{name: "Wolf fur", tipe: "Resource", level: 1, point: 2, value: 3, Quantity: 1},
				{name: "Boar fur", tipe: "Resource", level: 1, point: 2, value: 3, Quantity: 2},
			}
			if checkInventory(&joueurs[0], requiredItem) {
				removeInventory(&joueurs[0], "Wolf fur", requiredItem[0].Quantity)
				removeInventory(&joueurs[0], "Boar fur", requiredItem[1].Quantity)
				fmt.Println("Great you bought Boots of the adventurer")
				*Csolde -= 5
				itemToAdd := newItem("Boots of the adventurer", "Pieds", 1, 20, 3, 1)
				addInventory(&joueurs[0], itemToAdd)
			} else {
				fmt.Println("You dont have the required Items")
				time.Sleep(3 * time.Second)
			}
		//case 4:
		//SellItem(joueurs, itemToSell)
		case 4:
			a := 0
			menu3(joueurs, a)
		}
	}
}

// Tâche 7 : Marchand
func subMenu4(joueur *player) {
	for {
		clearScreen()
		Csolde := joueur.money
		fmt.Printf("\nMarchand: \nCurrent solde : %d\n", Csolde) // créer variable Csolde qui prend l'adresse de la valeur de joueur.money puis modifier la valeur vers laquelle pointe Csolde(Csolde pointe vers)(Csolde := &joueurs[0].money puis mettre des pointeur a Csolde partout en gros)
		fmt.Println("1. Potion of health, price : 3 gold")
		fmt.Println("2. Potion of poison, price : 6 gold")
		fmt.Println("3. Potion of mana, price : 3 gold")
		fmt.Println("4. Spellbook : Fireball, price : 25 gold")
		fmt.Println("5. Retour au menu principal")
		fmt.Print("What would you like to buy ? : ")
		var choixMarchand int
		_, err := fmt.Scanln(&choixMarchand)
		if err != nil {
			fmt.Println("Erreur de saisie.")
			continue
		}
		switch choixMarchand {
		case 1:
			StufToAdd := newItem("Potion of life", "Resource", 1, 40, 3, 1)
			buy(joueur, StufToAdd)
		case 2:
			StufToAdd := newItem("Potion of poison", "Resource", 1, 40, 6, 1)
			buy(joueur, StufToAdd)
		case 3:
			StufToAdd := newItem("Potion of mana", "Resource", 1, 40, 3, 1)
			buy(joueur, StufToAdd)
		case 4:
			StufToAdd := newSkills("Boule de feu", 1, 25, 10)
			if joueur.money >= 25 {
				if !contains(joueur.Skills, "Boule de feu") {
					joueur.Skills = append(joueur.Skills, StufToAdd)
					fmt.Println("Congrats ! You have learned the ability : Fireball.")
					time.Sleep(3 * time.Second)
					joueur.money -= 25
					fmt.Printf("\nCurrent solde : %d\n", joueur.money)
				} else {
					fmt.Println("You have already learned this ability : Fireball.")
					time.Sleep(3 * time.Second)
				}
			} else {
				fmt.Printf("\nyour solde is not sufficient")
				time.Sleep(3 * time.Second)
			}
		case 5:
			return
		default:
			fmt.Println("Choix invalide.")
			time.Sleep(3 * time.Second)
		}
	}
}

// Tâche 12 : Two for the Price of One

func subMenu5(joueurs []player) {
	for {
		clearScreen()
		Csolde := joueurs[0].money
		fmt.Printf("\nMarchand:\nCurrent solde : %d\n", Csolde) // créer variable Csolde qui prend l'adresse de la valeur de joueur.money puis modifier la valeur vers laquelle pointe Csolde(Csolde pointe vers)(Csolde := &joueurs[0].money puis mettre des pointeur a Csolde partout en gros)
		fmt.Println("1. Potion of life, price : 3 gold")
		fmt.Println("2. Potion of poison, price : 6 gold")
		fmt.Println("3. Potion of mana, price : 3 gold")
		fmt.Println("4. Feather crow, price 1 gold")
		fmt.Println("5. Elfic blade, price : 10 gold")
		fmt.Println("6. Dwarf made Armor, price : 7 gold")
		fmt.Println("7. Wolf fur, price  : 4 gold")
		fmt.Println("8. Troll skin, price : 2 gold")
		fmt.Println("9 Acacia wood, price : 6 gold")
		fmt.Println("10. Spellbook : boule de feu, price : 25 gold")
		fmt.Println("11. uprgrade your inventory, price : 15 gold")
		fmt.Println("12. Retour au menu principal")
		fmt.Print("Que souhaitez-vous acheter ? : ")
		var choixMarchand int
		_, err := fmt.Scanln(&choixMarchand)
		if err != nil {
			fmt.Println("Erreur de saisie.")
			continue
		}
		switch choixMarchand {
		case 1:
			StufToAdd := newItem("Potion of life", "Resource", 1, 40, 3, 1)
			buy(&joueurs[0], StufToAdd)
		case 2:
			StufToAdd := newItem("Potion of poison", "Resource", 1, 40, 6, 1)
			buy(&joueurs[0], StufToAdd)
		case 3:
			StufToAdd := newItem("Potion of mana", "Resource", 1, 40, 3, 1)
			buy(&joueurs[0], StufToAdd)
		case 4:
			StufToAdd := newItem("Feather crow", "Resource", 1, 10, 1, 1)
			buy(&joueurs[0], StufToAdd)
		case 5:
			StufToAdd := newItem("Elfic blade", "Arme", 1, 50, 10, 1)
			buy(&joueurs[0], StufToAdd)
		case 6:
			StufToAdd := newItem("Dwarf made armor", "Torse", 1, 30, 7, 1)
			buy(&joueurs[0], StufToAdd)
		case 7:
			StufToAdd := newItem("Wolf fur", "Resource", 1, 10, 4, 1)
			buy(&joueurs[0], StufToAdd)
		case 8:

			StufToAdd := newItem("Troll skin", "Resource", 1, 10, 7, 1)
			buy(&joueurs[0], StufToAdd)
		case 9:
			StufToAdd := newItem("Acacia wood", "Resource", 1, 10, 1, 1)
			buy(&joueurs[0], StufToAdd)
		case 10:
			StufToAdd := newSkills("Boule de feu", 1, 25, 10)
			if joueurs[0].money >= 25 {
				if !contains(joueurs[0].Skills, "Boule de feu") {
					joueurs[0].Skills = append(joueurs[0].Skills, StufToAdd)
					fmt.Println("Congrats ! You have learned the ability : Fireball.")
					time.Sleep(3 * time.Second)
					joueurs[0].money -= 25
					fmt.Printf("\nCurrent solde : %d\n", joueurs[0].money)
				} else {
					fmt.Println("You have already learned this ability : Fireball.")
					time.Sleep(3 * time.Second)
				}
			} else {
				fmt.Printf("\nyour solde is not sufficient")
			}
		case 11:
			if joueurs[0].money >= 15 {
				upgradeInventorySlot(&joueurs[0])
			} else {
				fmt.Printf("\nyour solde is not sufficient")
			}
		case 12:
			a := 0
			menu3(joueurs, a)
		default:
			fmt.Println("Invalid input.")
		}
	}
}

// function pour evite 30 ligne
func subMenu6(joueur *player) {
	clearScreen()
	fmt.Println("Here is your inventory!")
	displayInventory(*joueur)
	fmt.Println(len(joueur.inventory)+1, "- go back to the menu")
	fmt.Println("Enter the number of the item you want to equip or press ", len(joueur.inventory)+1, " to back:")
	var choice int
	fmt.Scanf("%d", &choice)

	if choice >= 1 && choice <= len(joueur.inventory) {
		selectedItem := joueur.inventory[choice-1]
		add(joueur, selectedItem) // Équiper l'élément
		subMenu6(joueur)
	} else if choice == len(joueur.inventory)+1 {
		return
	} else if choice < 1 && choice > len(joueur.inventory) && choice != len(joueur.inventory)+1 {
		fmt.Println("Invalid choice. Please enter a valid item number.")
		delaysup()
		subMenu6(joueur)
	}
	clearScreen()
	return
}

// menu 2 qui permet d'accede au shop apres le tuto
func menu2(joueurs []player, a int) {
	// Boucle du menu

	for {
		delaysup()
		fmt.Println("\nMenu:")
		fmt.Println("1. Start to play !")
		fmt.Println("2. Show the character information")
		fmt.Println("3. Inventory")
		fmt.Println("4. Take a health potion")
		fmt.Println("5. Shop")
		fmt.Println("6. Show equipement")
		fmt.Println("7. Quitt")

		var choice int
		fmt.Print("Choisissez une option : ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Erreur de saisie.")
			continue
		}
		switch choice {
		case 1:
			a := 0
			aventure(joueurs, a)
		case 2:
			subMenu1(joueurs[0])
		case 3:
			subMenu2(joueurs[0])
		case 4:
			takePot(&joueurs[0])
			delaysup()
		case 5:
			subMenu4(&joueurs[0])
		case 6:
			clearScreen()
			displayStuf(joueurs[0])
		case 7:
			fmt.Println("Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}
		dead(&joueurs[0])
	}
}

// menu 3 qui permet d'accede au shop apres chevalier
func menu3(joueurs []player, a int) {
	// Boucle du menu
	for {
		clearScreen()
		fmt.Println("\nMenu:")
		fmt.Println("1. Start to play !")
		fmt.Println("2. Show the character information")
		fmt.Println("3. Inventory")
		fmt.Println("4. Take a health potion")
		fmt.Println("5. Shop")
		fmt.Println("6. Show Equipement")
		fmt.Println("7. Black-Smith")
		fmt.Println("8. Quitt")
		fmt.Println("10. credits")
		var choice int
		fmt.Print("Choose an option : ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}
		switch choice {
		case 1:
			Fight1(joueurs[0])
		case 2:
			subMenu1(joueurs[0])
		case 3:
			subMenu6(&joueurs[0])
		case 4:
			takePot(&joueurs[0])
			delaysup()
		case 5:
			subMenu5(joueurs)
		case 6:
			clearScreen()
			displayStuf(joueurs[0])
		case 7:
			subMenu3(joueurs)
		case 8:
			fmt.Println("See you next time ! !")
			os.Exit(0)
		case 10:
			creators()
		default:
			fmt.Println("Invalid Input.")
		}
		dead(&joueurs[0])
	}
}
