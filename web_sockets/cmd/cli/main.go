package main

import (
	"fmt"
	"log"
	poker "main/command-line"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewTexasHoldem(poker.BlindAlertFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
