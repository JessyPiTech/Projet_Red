package main

import (
	"fmt"
	"time"
)

func lancement(joueurs []player, c int) {
	if c == 0 {
		c++
		anim()
	}
	if c == 1 {
		c++
		trainingFight(joueurs[0])
	}
	if c == 2 {
		c++
		a := 0
		menu2(joueurs, a)
	}
	a := 0
	menu2(joueurs, a)
}

func aventure(joueurs []player, a int) {
	if a == 0 {
		a++
		anim2()
	}
	if a == 1 {
		a++
		debut(joueurs)
	}
	if a == 2 {
		a++
		anim2()
	}
}

// /////////////////////////pour la propreter du code, fonction animation avant lancement de la game
const (
	barLength      = 20
	animationSpeed = 100 * time.Millisecond
	totalDuration  = 3 * time.Second
)

func drawBar(length int) {
	for i := 0; i < length; i++ {
		fmt.Print("0")
	}
	fmt.Println()
}

func anim() {
	clearScreen()
	startTime := time.Now()
	for {
		for i := barLength; i >= 0; i-- {
			clearScreen()
			drawBar(i)
			time.Sleep(animationSpeed)
		}
		if time.Since(startTime) >= totalDuration {
			break
		}
	}
}

/////////////////////

func anim2() {
	// Nombre total d'étapes de l'animation
	numSteps := 50

	// Boucle pour chaque étape
	for step := 1; step <= numSteps; step++ {
		clearScreen()
		drawDiagonals(step, 3)
		time.Sleep(100 * time.Millisecond) // Délai entre chaque étape
	}
	clearScreen()
}

func drawDiagonals(step, maxDiagonals int) {
	for i := 0; i < maxDiagonals; i++ {
		if step-i > 0 {
			// Dessine l'espace à gauche
			for j := 0; j < step-i-1; j++ {
				fmt.Print(" ")
			}
			// Dessine la diagonale "/"
			fmt.Println("/")
		}
	}
}

// /////////////////////function anim 3 decoupe en morso de moins de 30 ligne
func anim3() {
	p1()
	p2()
	p3()
	p4()
	p5()
	p6()
	p7()
	p8()
	p9()
	p10()
	p11()
	p12()
	p13()
	p14()
	p15()
	p16()
	p17()
	p18()
}
func p1() {
	timing()
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          //                        \n")
	fmt.Printf("-----ooo------------------ooo----------------ooo-----\n")
	timing()
}
func p2() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                            \\\\                        \n")
	fmt.Printf("                             ))                        \n")
	fmt.Printf("-ooo----------------ooo------------------ooo----------\n")
	timing()
}

func p3() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          ((                        \n")
	fmt.Printf("------------ooo----------------ooo----------------ooo\n")
	timing()
}

func p4() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          //                        \n")
	fmt.Printf("------ooo---------------ooo---------------ooo--------\n")
	timing()
}

func p5() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                            \\\\                        \n")
	fmt.Printf("                             ))                        \n")
	fmt.Printf("-ooo---------------ooo---------------ooo-------------o\n")
	timing()
}

func p6() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          ((                        \n")
	fmt.Printf("---------------ooo---------------ooo-------------ooo-\n")
	timing()
}
func p7() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          //                        \n")
	fmt.Printf("-----------ooo-----------------ooo------------ooo----\n")
	timing()
}
func p8() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                            \\\\                        \n")
	fmt.Printf("                             ))                        \n")
	fmt.Printf("-------ooo-------------ooo---------------ooo----------\n")
	timing()
}

func p9() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          ((                        \n")
	fmt.Printf("-----------ooo----------------ooo----------------ooo--\n")
	timing()
}

func p10() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          //                        \n")
	fmt.Printf("----ooo----------------ooo------------------ooo-----\n")
	timing()
}

func p11() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                            \\\\                        \n")
	fmt.Printf("                             ))                        \n")
	fmt.Printf("ooo----------------ooo-------------------ooo----------\n")
	timing()
}

func p12() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          ((                        \n")
	fmt.Printf("------------ooo---------------ooo--------------ooo--\n")
	timing()
}

func p13() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          //                        \n")
	fmt.Printf("------ooo---------------ooo-------------ooo----------\n")
	timing()
}

func p14() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                            \\\\                        \n")
	fmt.Printf("                             ))                        \n")
	fmt.Printf("ooo----------------ooo---------------ooo--------------\n")
	timing()
}

func p15() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          ((                        \n")
	fmt.Printf("---------------ooo-------------ooo------------ooo----\n")
	timing()
}

func p16() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          //                        \n")
	fmt.Printf("-----------ooo-----------ooo----------ooo----------\n")
	timing()
}

func p17() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                            \\\\                        \n")
	fmt.Printf("                             ))                        \n")
	fmt.Printf("-------ooo--------------------------------------------\n")
	timing()
}

func p18() {
	fmt.Printf("                             0                        \n")
	fmt.Printf("                          ------                        \n")
	fmt.Printf("                            /                      \n")
	fmt.Printf("                           ))                        \n")
	fmt.Printf("                          ((                        \n")
	fmt.Printf("--ooo------------------------------------------------\n")
	timing()
}

////////////////////

func timing() {
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\033[H\033[2J")
}
