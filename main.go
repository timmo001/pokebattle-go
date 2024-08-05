package main

import (
	"log"
	"math/rand"

	"github.com/charmbracelet/huh"
)

type MoveType string

const (
	Normal MoveType = "Normal"
	Fire   MoveType = "Fire"
	Grass  MoveType = "Grass"
	Water  MoveType = "Water"
)

type Move struct {
	Name  string
	Type  MoveType
	Power int
}

type Creature struct {
	Name       string
	Type       MoveType
	Health     int
	Weaknesses []MoveType
	Moves      []Move
}

var (
	FireGuy = Creature{
		Name:       "Fire Guy",
		Type:       Fire,
		Health:     100,
		Weaknesses: []MoveType{Water},
		Moves: []Move{
			{
				Name:  "Scratch",
				Type:  Normal,
				Power: 5,
			},
			{
				Name:  "Fireball",
				Type:  Fire,
				Power: 10,
			},
			{
				Name:  "Flamethrower",
				Type:  Fire,
				Power: 15,
			},
		},
	}
	LeafyBoi = Creature{
		Name:       "Leafy Boi",
		Type:       Grass,
		Health:     100,
		Weaknesses: []MoveType{Fire},
		Moves: []Move{
			{
				Name:  "Tackle",
				Type:  Normal,
				Power: 5,
			},
			{
				Name:  "Vine Whip",
				Type:  Grass,
				Power: 10,
			},
			{
				Name:  "Solar Beam",
				Type:  Grass,
				Power: 15,
			},
		},
	}
	WaterDude = Creature{
		Name:       "Water Dude",
		Type:       Water,
		Health:     100,
		Weaknesses: []MoveType{Grass},
		Moves: []Move{
			{
				Name:  "Tackle",
				Type:  Normal,
				Power: 5,
			},
			{
				Name:  "Water Gun",
				Type:  Water,
				Power: 10,
			},
			{
				Name:  "Hydro Pump",
				Type:  Water,
				Power: 15,
			},
		},
	}
)

type Game struct {
	PlayerCreature   *Creature
	OppenentCreature *Creature
}

func main() {
	var (
		playerCreature *Creature
		playerMove     *Move
	)

	formCreature := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[*Creature]().
				Title("Choose your creature").
				Options(
					huh.NewOption(FireGuy.Name, &FireGuy),
					huh.NewOption(LeafyBoi.Name, &LeafyBoi),
					huh.NewOption(WaterDude.Name, &WaterDude),
				).
				Value(&playerCreature),
		),
	)
	ErrCheck(formCreature.Run())

	game := Game{
		PlayerCreature:   playerCreature,
		OppenentCreature: &WaterDude,
	}

	log.Printf("Player chose:   %s", game.PlayerCreature.Name)
	log.Printf("Opponent chose: %s", game.OppenentCreature.Name)

	log.Printf("Begin battle!")

	log.Printf("Player health:   %d", game.PlayerCreature.Health)
	log.Printf("Opponent health: %d", game.OppenentCreature.Health)

	for game.PlayerCreature.Health > 0 && game.OppenentCreature.Health > 0 {
		playerMoveOptions := []huh.Option[*Move]{}
		for _, move := range game.PlayerCreature.Moves {
			playerMoveOptions = append(playerMoveOptions, huh.NewOption(move.Name, &move))
		}

		formMove := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[*Move]().
					Title("Choose your move").
					Options(playerMoveOptions...).
					Value(&playerMove),
			),
		)
		ErrCheck(formMove.Run())

		log.Printf("Player chose move:   %s", playerMove.Name)

		if playerMove.Type == game.OppenentCreature.Type {
			log.Println("Its not very effective...")
			game.OppenentCreature.Health -= playerMove.Power / 2
		} else if playerMove.Type == game.OppenentCreature.Weaknesses[0] {
			log.Println("Its super effective!")
			game.OppenentCreature.Health -= playerMove.Power * 2
		} else {
			// Normal damage
			game.OppenentCreature.Health -= playerMove.Power
		}

		// Choose a random move for the opponent
		randomMoveIndex := rand.Int() % len(game.OppenentCreature.Moves)
		opponentMove := &game.OppenentCreature.Moves[randomMoveIndex]

		log.Printf("Opponent chose move: %s", opponentMove.Name)

		if opponentMove.Type == game.PlayerCreature.Type {
			log.Println("Its not very effective...")
			game.PlayerCreature.Health -= opponentMove.Power / 2
		} else if opponentMove.Type == game.PlayerCreature.Weaknesses[0] {
			log.Println("Its super effective!")
			game.PlayerCreature.Health -= opponentMove.Power * 2
		} else {
			// Normal damage
			game.PlayerCreature.Health -= opponentMove.Power
		}

		log.Printf("Player health:   %d", game.PlayerCreature.Health)
		log.Printf("Opponent health: %d", game.OppenentCreature.Health)
	}

	if game.PlayerCreature.Health <= 0 {
		log.Println("You lose!")
	} else {
		log.Println("You win!")
	}
}

func ErrCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
