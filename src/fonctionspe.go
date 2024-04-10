package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

// Tâche 2 : Création du personnage
func createPlayer() player {

	go playAudioFiles()
	delaysup()
	fmt.Println("Welcome to our RPG game on Go !")

	var nom string
	for {
		fmt.Print("Enter the name of your character : ")
		fmt.Scanln(&nom)
		if isValidName(nom) {
			break
		}

		fmt.Println("Invalid name. Please enter a name without numbers or special characters.")
	}

	classe := chooseClass()
	return newPlayer(nom, classe)
}

// Fonction pour choisir la classe
func chooseClass() string {
	var choixClasse int

	for {
		fmt.Println("Choose your class :")
		fmt.Println("1. Human")
		fmt.Println("2. Elf")
		fmt.Println("3. Dwarf")
		fmt.Print("Your choice (1/2/3): ")
		_, err := fmt.Scan(&choixClasse)
		if err != nil || (choixClasse != 1 && choixClasse != 2 && choixClasse != 3) {
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
			continue
		}
		break
	}

	switch choixClasse {
	case 1:
		return "Human"
	case 2:
		return "Elf"
	case 3:
		return "Dwarf"
	default:
		return "Unknown"
	}
}

// met au norme le nom
func formatNom(nom string) string {
	nom = strings.ToLower(nom) // Convertit le nom en minuscules
	nom = strings.Title(nom)   // Met la première lettre en majuscule
	return nom
}

// verif c'est bin des letrte
func isValidName(name string) bool {
	for _, char := range name {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

// fonction pour jouer les fichiers audio en boucle
// /a un probleme avec cet fonction car de base sa dois lire plusieur fichier audio de base mais gitea accepte pas les fichier de plus de 1Mo
func playAudioFiles() {
	// initialise le son
	err := speaker.Init(44100, 44100/10)
	if err != nil {
		log.Fatal(err)
	}

	// boucle infinie
	for {
		//for _, filename := range audioFiles {
		// ouvre le fichier audio

		f, err := os.Open("serieux.mp3")
		if err != nil {
			log.Fatal(err)
		}
		streamer, _, err := mp3.Decode(f)

		if err != nil {
			log.Fatal(err)
		}

		// joue le fichier audio
		done := make(chan struct{})
		speaker.Play(beep.Seq(streamer))

		// attend que la lecture soit terminée
		<-done

		// ferme le fichier audio
		f.Close()
		//}
	}
}

// donne acces a inventaire
func accessInventory(joueur player) {
	clearScreen()
	fmt.Println("Here is your inventory:")
	for i, item := range joueur.inventory {
		fmt.Printf("%d. %s\n", i+1, item.name)
	}
	fmt.Println("Press enter to go back to the menu")
	var choice int
	fmt.Scanf("%d", &choice)
	clearScreen()
}

func add(joueur *player, StufToAdd item) {
	switch StufToAdd.tipe {
	case "Resource":
		fmt.Println("this item was an ressource and can't equipe in stuff")
		delaysup()
		addInventory(joueur, StufToAdd)
	case "Tete":
		addStufTete(joueur, StufToAdd)
	case "Torse":
		addStufTorse(joueur, StufToAdd)
	case "Pieds":
		addStufPieds(joueur, StufToAdd)
	case "Arme":
		addStufArme(joueur, StufToAdd)
	case "Bouclier":
		addStufBouclier(joueur, StufToAdd)
	case "ArmeSecondair":
		addStufArmeSecondair(joueur, StufToAdd)
	default:
		fmt.Printf("Unknown item type: %s\n", StufToAdd.tipe)
	}
	fmt.Println("you have equiped", StufToAdd.name)
	delaysup()
}

//fonction pour revendre des item mais sa marche pas encore
/*
	func SellItem(joueurs []player, itemToSell item) {
		clearScreen()
		fmt.Println("Here is your inventory:")
		for i, item := range joueurs[0].inventory {
			fmt.Printf("%d. %s\n", i+1, item.name)
		}
		fmt.Println("Enter the name of the Item you want to sell :")
		fmt.Scanf("%s", itemToSell.name, itemToSell.Quantity)
		removeInventory(&joueurs[0], itemToSell.name, itemToSell.Quantity)

}
*/
//fonction pour ajuoter de l'equipend
func addInventory(joueur *player, itemToAdd item) {

	// parcourez l'inventaire existant pour voir si l'item est déjà présent
	for i, existingItem := range joueur.inventory {
		if existingItem.name == itemToAdd.name {
			//si est déjà, augmentez la quantité
			joueur.inventory[i].Quantity += itemToAdd.Quantity
			return
		}
	}

	//si l'item n'a pas été trouvé sa l'ajoute a l'inventaire
	joueur.inventory = append(joueur.inventory, itemToAdd)
}

// fonctionnement par equipment qui evite que totu soit dans inventaires
func addEquipment(joueur *player, StufToAdd item, equipmentSlice *[]item) {
	StufToAdd.Quantity = 1
	if len(joueur.equipement) == 0 {
		joueur.equipement = append(joueur.equipement, equipment{})
	}
	if len(*equipmentSlice) == 0 {
		*equipmentSlice = append(*equipmentSlice, StufToAdd)
		return // Élément ajouté
	}

	// parcourir l'équipement pour vérifier si ya deja un item à la même position
	for i := 0; i < len(*equipmentSlice); i++ {
		if (*equipmentSlice)[i].tipe == StufToAdd.tipe {
			// Remplacer l'item existant
			(*equipmentSlice)[i] = StufToAdd
			return
		}
	}

	//si pas item du même type n'a été trouvé, ajouter le nouvel item à la fin
	*equipmentSlice = append(*equipmentSlice, StufToAdd)
}

// /////////////////sussetion de add pour chaque truc tete torse pieds arme ...///////////////////////

//Tâche 15 : Mamma Mia

func addStufTete(joueur *player, StufToAdd item) {
	equipmentSlot := &joueur.equipement[0].Tete // Accédez à la tranche d'équipement pour la tête

	// Vérifiez si la tranche d'équipement pour la tête est vide
	if len(*equipmentSlot) > 0 {
		// Si déjà élément dans équipement,
		// déplacez-le dans l'inventaire
		joueur.pvmax -= joueur.equipement[0].Tete[0].point
		addInventory(joueur, (*equipmentSlot)[0])

	}
	// Équipez le nouvel élément et enlever
	removeInventory(joueur, StufToAdd.name, 1)
	addEquipment(joueur, StufToAdd, &joueur.equipement[0].Tete)
	joueur.pvmax += joueur.equipement[0].Tete[0].point
}

func addStufTorse(joueur *player, StufToAdd item) {
	equipmentSlot := &joueur.equipement[0].Torse

	if len(*equipmentSlot) > 0 {
		joueur.pvmax -= joueur.equipement[0].Torse[0].point
		addInventory(joueur, (*equipmentSlot)[0])

	}
	removeInventory(joueur, StufToAdd.name, 1)
	addEquipment(joueur, StufToAdd, &joueur.equipement[0].Torse)
	joueur.pvmax += joueur.equipement[0].Torse[0].point

}

func addStufPieds(joueur *player, StufToAdd item) {
	equipmentSlot := &joueur.equipement[0].Pieds

	if len(*equipmentSlot) > 0 {
		joueur.pvmax -= joueur.equipement[0].Pieds[0].point
		addInventory(joueur, (*equipmentSlot)[0])

	}
	removeInventory(joueur, StufToAdd.name, 1)
	addEquipment(joueur, StufToAdd, &joueur.equipement[0].Pieds)
	joueur.pvmax += joueur.equipement[0].Pieds[0].point
}

func addStufArme(joueur *player, StufToAdd item) {
	equipmentSlot := joueur.equipement[0].Arme

	if len(equipmentSlot) > 0 {
		addInventory(joueur, equipmentSlot[0])

	}
	removeInventory(joueur, StufToAdd.name, 1)
	addEquipment(joueur, StufToAdd, &joueur.equipement[0].Arme)
}

func addStufBouclier(joueur *player, StufToAdd item) {
	equipmentSlot := &joueur.equipement[0].Bouclier

	if len(*equipmentSlot) > 0 {
		joueur.pvmax -= joueur.equipement[0].Bouclier[0].point
		addInventory(joueur, (*equipmentSlot)[0])
	}
	removeInventory(joueur, StufToAdd.name, 1)
	addEquipment(joueur, StufToAdd, &joueur.equipement[0].Bouclier)
	joueur.pvmax += joueur.equipement[0].Bouclier[0].point
}

func addStufArmeSecondair(joueur *player, StufToAdd item) {
	equipmentSlot := &joueur.equipement[0].ArmeSecondair

	if len(*equipmentSlot) > 0 {
		addInventory(joueur, (*equipmentSlot)[0])
	}
	removeInventory(joueur, StufToAdd.name, 1)
	addEquipment(joueur, StufToAdd, &joueur.equipement[0].ArmeSecondair)
}

///////////////////////////////////////////////////////////////////////////////////////////

func removeInventory(joueur *player, itemNameToRemove string, q int) {
	for i, inventoryItem := range joueur.inventory {
		if inventoryItem.name == itemNameToRemove {

			if inventoryItem.Quantity > 1 {

				joueur.inventory[i].Quantity -= q // Décrémenter la quantité de l'élément réel
				if joueur.inventory[i].Quantity == 0 {
					joueur.inventory = append(joueur.inventory[:i], joueur.inventory[i+1:]...)

					return
				}
				return
			} else {

				// Retirer l'élément de l'inventaire
				joueur.inventory = append(joueur.inventory[:i], joueur.inventory[i+1:]...)

				return
			}
		}
	}
}

// Tâche 6 : Potion de vie
func takePot(joueur *player) {
	for _, item := range joueur.inventory {
		if item.name == "Potion of life" {
			joueur.pvac += 50
			if joueur.pvac > joueur.pvmax {
				joueur.pvac = joueur.pvmax
			}
			fmt.Printf("Vous avez utilisé une Potion of life. Points de vie actuels : %d/%d\n", joueur.pvac, joueur.pvmax)
			removeInventory(joueur, "Potion of life", 1)
			return
		}

	}
	fmt.Println("Aucune Potion of life trouvée dans l'inventaire.")
}

// utilise popo et appel enlevement
func takeMana(joueur *player) {
	for _, item := range joueur.inventory {
		if item.name == "Potion of mana" {
			joueur.Mana += 40
			if joueur.Mana > joueur.ManaMax {
				joueur.Mana = joueur.ManaMax
			}
			fmt.Printf("Vous avez utilisé une Potion of mana. mana : %d/%d\n", joueur.Mana, joueur.ManaMax)
			removeInventory(joueur, "Potion of mana", 1)
			return
		}
	}
	fmt.Println("Aucune Potion of mana trouvée dans l'inventaire.")
}

// Tâche 9 : Potion de poison
// Fonction pour infliger des dégâts de poison
func poisonPot(joueur *player, ennemi *Monstre) {
	fmt.Println("Vous avez empoisonné ", ennemi.Nom, "!")
	for _, item := range joueur.inventory {
		if item.name == "Potion of poison" {
			for i := 0; i < 3; i++ {
				ennemi.pvac -= 10
				if ennemi.pvac < 0 {
					ennemi.pvac = 0
				}
				fmt.Printf("Points de vie actuels: %d/%d\n", ennemi.pvac, ennemi.pvmax)
				time.Sleep(1 * time.Second)
			}
			removeInventory(joueur, "Potion of poison", 1)
			return
		}
	}
	fmt.Println("Aucune Potion of poison trouvée dans l'inventaire.")
}

// Tâche 8 : Wasted !
// fonction quand tes mort revive 50%
func dead(joueur *player) {
	if joueur.pvac <= 0 {
		delaysup()
		fmt.Println("Vous êtes mort ! Vous êtes ressuscité avec 50% de vos points de vie maximum.")
		joueur.pvac = joueur.pvmax / 2
	}
}

// Tâche 4 : Affichage des informations du personnage
func displayInfo(joueur player) {
	fmt.Printf("Nom: %s, Classe: %s, Niveau: %d, PV Actuels: %d/%d, mana: %d/%d \n",
		joueur.name, joueur.class, joueur.level, joueur.pvac, joueur.pvmax, joueur.Mana, joueur.ManaMax)
	for _, Skills := range joueur.Skills {
		fmt.Printf("Skill: %s, Level: %d, Degat: %d\n", Skills.name, Skills.level, Skills.degats)
	}
}

// Tâche 5 : Accès à l’inventaire
func displayInventory(joueur player) {
	fmt.Printf("Inventory of %s:\n", joueur.name)
	a := 1
	for _, item := range joueur.inventory {
		fmt.Printf("%d - Item: %s, Level: %d, Point: %d, Value: %d, Quantity :%d\n", a, item.name, item.level, item.point, item.value, item.Quantity)
		a++
	}
}

// affiche equipement
func displayStuf(joueur player) {

	fmt.Printf("Head ")
	displayItem(joueur.equipement[0].Tete)

	fmt.Printf("Chest ")
	displayItem(joueur.equipement[0].Torse)

	fmt.Printf("Feet ")
	displayItem(joueur.equipement[0].Pieds)

	fmt.Printf("Main Weapon ")
	displayItem(joueur.equipement[0].Arme)

	fmt.Printf("Secondary Weapon ")
	displayItem(joueur.equipement[0].ArmeSecondair)

	fmt.Printf("Shield ")
	displayItem(joueur.equipement[0].Bouclier)

	fmt.Println("Press enter to go back to the menu")
	var choice int
	fmt.Scanf("%d", &choice)
	clearScreen()
}

// affiche item dans terminal et propose de l'utilise
func displayItem(items []item) {
	if len(items) > 0 {
		fmt.Printf("Item: %s, Type: %s, Level: %d, Point: %d, Value: %d\n", items[0].name, items[0].tipe, items[0].level, items[0].point, items[0].value)
	} else {
		fmt.Println("No item equipped")
	}
}

// affiche un item spe dans terminal et propose de l'utilise
func displayItem1(joueurs []player) {
	for _, joueur := range joueurs {
		if len(joueur.inventory) > 0 {
			fmt.Printf("Nom joueur: %s, Nom du premier item: %s, Point du premier item: %d\n", joueur.name, joueur.inventory[0].name, joueur.inventory[1].point)
		} else {
			fmt.Printf("Nom joueur: %s, Inventaire vide\n", joueur.name)
		}
	}
}

// Mission 5.1 : Combat magique
func attaque(joueur *player, ennemi *Monstre) {
	fmt.Println(" ")
	fmt.Println("Choisissez une attaque:")
	// Afficher les compétences du joueur
	for i, skill := range joueur.Skills {
		fmt.Printf("%d. %s (Level %d)\n", i+1, skill.name, skill.level)
	}

	fmt.Printf("%d. Retour\n", len(joueur.Skills)+1)
	fmt.Println(" ")
	var choix int
	fmt.Print("Choisissez une compétence : ")
	fmt.Scan(&choix)

	if choix < 1 || choix > len(joueur.Skills)+1 {
		fmt.Println("Choix invalide.")
		return
	}

	// Si le joueur choisit "Retour", sortir de la fonction
	if choix == len(joueur.Skills)+1 {
		clearScreen()
		charTurn(joueur, ennemi)
	} else {
		// Le joueur a choisi une compétence et l'utilise
		selectedSkill := joueur.Skills[choix-1]
		attaqueD(joueur, ennemi, selectedSkill)
	}
}

// Mission 5.2 : Ressource de mana
func attaqueD(joueur *player, ennemi *Monstre, selectedSkill skills) {
	if joueur.Mana < selectedSkill.manaDecresse {
		fmt.Printf("attaque imposible mana insufisant")
		return
	}
	fmt.Printf("vous attaque avec %s.\n", selectedSkill.name)
	joueur.Mana -= selectedSkill.manaDecresse
	fmt.Printf("%s  mana : %d/%d\n", joueur.name, joueur.Mana, joueur.ManaMax)
	fmt.Printf("Vous infligez  %d dégats au %s", selectedSkill.degats, ennemi.Nom)
	ennemi.pvac -= selectedSkill.degats
	if ennemi.pvac < 0 {
		ennemi.pvac = 0
	}
	fmt.Printf("%s  PV : %d/%d\n", ennemi.Nom, ennemi.pvac, ennemi.pvmax)
	time.Sleep(2 * time.Second)
}

func inventoryFight(joueur *player, ennemi *Monstre) {
	// Affiche l'inventaire du joueur
	fmt.Println("Inventaire du joueur:")
	if len(joueur.inventory) >= 1 {
		for i, item := range joueur.inventory {
			fmt.Printf("%d. %s, Points: %d, Valeur: %d)\n", i+1, item.name, item.point, item.value)
		}
		fmt.Println(len(joueur.inventory)+1, "retour")

		fmt.Println("Choisissez un objet dans votre inventaire (1 - ", len(joueur.inventory)+1, ") : ")
		var choixObjet int
		fmt.Scan(&choixObjet)

		if choixObjet == len(joueur.inventory)+1 {
			clearScreen()
			charTurn(joueur, ennemi)

		} else if choixObjet >= 1 && choixObjet <= len(joueur.inventory) {
			objet := joueur.inventory[choixObjet-1]

			if objet.name == "Potion of life" {
				takePot(joueur)
			} else if objet.name == "Potion of poison" {
				poisonPot(joueur, ennemi)
			} else if objet.name == "Potion of mana" {
				takeMana(joueur)
			} else {
				fmt.Println("Vous ne pouvez pas utiliser cet objet en combat.")
			}

		} else {
			fmt.Println("Choix d'objet invalide.")
		}
	} else {
		fmt.Println("inventaire vide")
	}
}

func contains(skills []skills, skill string) bool {
	for _, s := range skills {
		if s.name == skill {
			return true
		}
	}
	return false
}

func addMoney(joueur player, a int) {
	joueur.money += a
	fmt.Println("csolde:", joueur.money)
	delaysup()
}

func checkInventory(joueur *player, requiredItems []item) bool {
	for _, requiredItem := range requiredItems {
		itemFound := false
		for _, inventoryItem := range joueur.inventory {
			if inventoryItem.name == requiredItem.name {
				if inventoryItem.Quantity >= requiredItem.Quantity {
					itemFound = true
					break
				}
			}
		}
		// Si l'objet requis n'est pas trouvé ou la quantité est insuffisante, retournez false
		if !itemFound {
			return false
		}
	}
	// Si tous les éléments requis sont présents en quantité suffisante, retournez true
	return true
}

// Tâche 18 : A.I. Intelligence artificielle
// Fonction pour simuler le modèle de combat
func ennemiPattern(ennemi Monstre, joueur *player) {
	fmt.Println("Tour de", ennemi.Nom)
	degats := ennemi.PointsAttaque
	joueur.pvac -= degats
	fmt.Printf("%s inflige %d dégats à %s.\n", ennemi.Nom, degats, joueur.name)
	fmt.Printf("%s PV : %d/%d\n", joueur.name, joueur.pvac, joueur.pvmax)
}

func manaTours(joueur *player) {
	joueur.Mana += 5

	if joueur.Mana > joueur.ManaMax {
		joueur.Mana = joueur.ManaMax
		fmt.Printf("mana :%d/%d.\n", joueur.Mana, joueur.ManaMax)
		return
	} else {
		fmt.Printf("mana + 5\n")
		fmt.Printf("mana :%d/%d.\n", joueur.Mana, joueur.ManaMax)
	}
}

// Bonus 2 : Qui sont-ils ?
func creators() {
	word := "motsdepasse"
	attempts := 10
	guessedLetters := make([]bool, len(word))

	// Boucle principale du jeu
	for attempts > 0 {
		displayWord(word, guessedLetters)
		fmt.Print("\npropose :")
		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')
		guess = strings.ToLower(strings.TrimSpace(guess))
		if len(guess) != 1 && guess != word {
			fmt.Println("Veuillez entrer une seul lettre.")
			continue
		}
		if guess == word {
			fmt.Println("bien jouer vous avez fini le jeu")
			delaysup()
			fmt.Println("les createur du jeu sont")
			fmt.Println("Jessy Piquerel\n")
			fmt.Println("et\n")
			fmt.Println("Mamoune Benouna\n")
			time.Sleep(5 * time.Second)
			break
		}
		if !isValidName(guess) {
			fmt.Println("Attention : seules les lettres alphabétiques sont autorisées.")
			continue
		}
		// Vérifier si la lettre propose est das mots
		found := false
		for i, char := range word {
			if guessedLetters[i] {
				continue
			}
			if guess == string(char) {
				guessedLetters[i] = true
				found = true
			}
		}
		if !found {
			fmt.Println("Lettre incorrecte. Vous avez", attempts-1, "essais restants.")
			attempts--
		} else if allLettersGuessed(guessedLetters) {
			fmt.Print("Mot : motsdepasse")

			fmt.Println("bien jouer vous avez fini le jeu")
			delaysup()
			fmt.Println("les createur du jeu sont")
			fmt.Println("Jessy Piquerel\n")
			fmt.Println("et\n")
			fmt.Println("Mamoune Benouna\n")
			time.Sleep(5 * time.Second)
			break
		}
	}

	if attempts == 0 {
		fmt.Println("Perdu le mots étais: ", word)
	}

}

func displayWord(word string, guessedLetters []bool) {
	// Afficher le mot avec des underscores pour les lettres non devinées
	fmt.Print("Mot : ")

	for i, char := range word {
		if guessedLetters[i] {
			fmt.Print(string(char))
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func allLettersGuessed(guessedLetters []bool) bool {
	// Vérifi si toute les lettres trouve
	for _, guessed := range guessedLetters {
		if !guessed {
			return false
		}
	}
	return true
}

// Mission 2.2 : On and on and on
func upgradeInventorySlot(joueur *player) {

	if joueur.charge <= 20 {
		joueur.charge += 10
		joueur.money -= 15

		fmt.Println("you have incressed your inventory by 10 slot")
		fmt.Println("slot: ", joueur.charge)
		fmt.Printf("\nCurrent solde : %d\n", joueur.money)
		time.Sleep(3 * time.Second)

	} else if joueur.charge > 20 && joueur.charge < 30 {
		joueur.charge = 30
		joueur.money -= 15

		fmt.Println("you have incressed your inventory slot to the max")
		fmt.Println("slot: ", joueur.charge)
		fmt.Printf("\nCurrent solde : %d\n", joueur.money)
		time.Sleep(3 * time.Second)

	} else if joueur.charge >= 30 {
		fmt.Println("you have already the max slot")
		time.Sleep(3 * time.Second)
	}

}

func levelUp(joueur *player) {
	joueur.Experience -= joueur.ExperienceMax
	joueur.ExperienceMax += 20
	joueur.level++
	joueur.ManaMax += joueur.level * 10
	joueur.pvmax += 10
	joueur.pvac = joueur.pvmax
	fmt.Printf("Félicitations ! Vous êtes maintenant niveau %d.\n", joueur.level)
	delaysup()
}

// Mission 4 : Expérience
// Fonction pour gagner de l'expérience
func gainExperience(joueur *player, exp int) {
	clearScreen()
	fmt.Printf("Vous avez gagné %d points d'expérience.\n", exp)
	joueur.Experience += exp
	if joueur.Experience >= joueur.ExperienceMax {
		levelUp(joueur)
	}
}

func buy(joueur *player, StufToAdd item) {
	if joueur.money >= StufToAdd.value {
		// vérifiez d'abord si l'ajout de l'article dépasse la charge maximale pour pas rajouter alors que cet desja plein
		TotalQuantity := 0
		for _, existingItem := range joueur.inventory {
			TotalQuantity += existingItem.Quantity
			//d'abord sa compte
		}
		if TotalQuantity+StufToAdd.Quantity <= joueur.charge {
			fmt.Println("You bought :", StufToAdd.name)
			joueur.money -= StufToAdd.value
			fmt.Printf("\nCurrent solde : %d\n", joueur.money)
			addInventory(joueur, StufToAdd)
			delaysup()
		} else {
			fmt.Println("Transaction refused, Your inventory is full !")
			delaysup()
			return
		}

	} else {
		fmt.Printf("\nyour solde is not sufficient")
		delaysup()
	}
}

// //////////////////////////////////////////////////////
// /a changer en  delay 4s
// fonction qui fait un delay et supprime tout de l'ecran
func delaysup() {
	time.Sleep(3 * time.Second)
	fmt.Print("\033[H\033[2J")
}

// /////////////////////////////////////////////////////////////////////// //
