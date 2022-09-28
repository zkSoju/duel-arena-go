package main

import (
	"errors"
	"fmt"
)

const ()

type Player struct {
	Number    int
	Health    int
	Inventory []int
}

var PLAYER_1 = Player{1, 100, []int{}}
var PLAYER_2 = Player{2, 100, []int{}}

type Stage int

type Game struct {
	Turn  Player
	Stage int
}

var Players = map[int]Player{
	1: PLAYER_1,
	2: PLAYER_2,
}

var CombatItems = map[int]string{
	1: "Monkfish",        // heals 16
	2: "Strength Potion", // increases attack by 10%
	3: "Dharok's Axe",    // has a higher % chance to hit higher with low hp
}

func New() *Game {
	game := &Game{PLAYER_1, 1}
	return game
}

func (game *Game) SelectCombatEquipment(player Player, items []int) (err error) {
	if game.Stage != 1 {
		return errors.New(fmt.Sprintf("Player has to choose items to bring to duel"))
	}
	if !game.TurnIs(player) {
		return errors.New(fmt.Sprintf("Not the players turn"))
	} else {
		player.Inventory = items
	}

	return
}

func (game *Game) Choose() {}

func (game *Game) TurnIs(player Player) bool {
	return game.Turn.Equals(&player)
}

func (player *Player) Equals(other *Player) bool {
	number_eq := player.Number == other.Number
	health_eq := player.Health == other.Health
	inventory_eq := len(player.Inventory) == len(other.Inventory)

	for index, item := range player.Inventory {
		if item != other.Inventory[index] {
			inventory_eq = false
		}
	}

	return (number_eq && health_eq && inventory_eq)
}
