package chess

import (
	"log"

	"github.com/notnil/chess"
)

func (m *Match) Forfeit() {
	log.Printf("The master has forfeited")
	m.Forfeited = true
	m.Game.Resign(chess.White)
	m.sendUpdate()
	m.PostGame()
	m.end()
}

func (m *Match) bestTwitterMove() (move *string) {
	err := m.FetchTweets()
	if err != nil {
		log.Printf("Could not get tweets replies: %v", err)
		return
	}
	if len(m.Moves) != 0 {
		var mostValued uint
		for mv, val := range m.Moves {
			if val > mostValued {
				newmv := mv
				move = &newmv
				mostValued = val
			}
		}
	} else {
		move = m.randomMove()
	}
	return
}

func (m *Match) onTurnEnd() {
	var move *string
	if m.Game.Position().Turn() == playerColor {
		// Play a random move if the palyer didn't pick in time
		move = m.randomMove()
	} else {
		move = m.bestTwitterMove()
	}
	if move != nil {
		if err := m.move(*move); err != nil {
			log.Printf("WARN: move %s failed:  %v", *move, err)
		}
	} else {
		// should never be reached
		m.end()
	}
}
