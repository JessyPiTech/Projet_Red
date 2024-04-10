package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	filesToRun := []string{"jeutest.go", "histoir.go", "menu.go", "lancement.go", "fonctionspe.go"}
	// Combine les noms de fichiers en une seule chaîne avec des espaces
	cmdArgs := append([]string{"go", "run"}, filesToRun...)
	cmd := exec.Command("cmd", "/c", "start", "cmd.exe", "/k", strings.Join(cmdArgs, " "))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Erreur lors de l'exécution de la commande : %v\n", err)
	}
}
