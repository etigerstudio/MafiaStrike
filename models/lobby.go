package models

import (
	"errors"
	"mafia-strike/consts"
	"mafia-strike/util"
	"math/rand"
)

var Lobbies = map[int]*Lobby{}
var lobbyCount = 0

type Lobby struct{
	Players map[string]*Player
	Keywords []string
	CurrentWord string
	PrevWord string
	PrevMafias []string
	Round int
}

func (l *Lobby) PlayerCount() int {
	return len(l.Players)
}

func (l *Lobby) NextMafiaCount() int {
	return (len(l.Players) - 1)/2 + (len(l.Players) - 1)%2 * rand.Intn(2)
}

func (l *Lobby) CurrentMafiaCount() int {
	count := 0
	for _, p := range l.Players {
		if p.IsActive && p.IsMafia {
			count++
		}
	}
	return count
}

func (l *Lobby) AddPlayer(nickname string, creator bool) string {
	playerID := util.GenerateRandomShortID()
	l.Players[playerID] = &Player{
		Nickname:  nickname,
		IsCreator: creator,
	}
	return playerID
}

func (l *Lobby) StartNewRound() error {
	if len(l.Keywords) == 0 {
		return errors.New(consts.ResponseErrorDescNoKeyword)
	}

	if l.CurrentWord != "" {
		l.PrevWord = l.CurrentWord
	}

	l.CurrentWord, l.Keywords = l.Keywords[0], l.Keywords[1:]

	mafiaCount := l.NextMafiaCount()

	mafias := []string{}
	playersCopy := []string{}
	for k := range l.Players {
		playersCopy = append(playersCopy, k)
	}

	for i := 0; i < mafiaCount; i++ {
		mafia := rand.Intn(mafiaCount + 1)
		mafias = append(mafias, playersCopy[mafia])
		playersCopy = util.RemoveSliceElement(playersCopy, mafia)
	}

	l.PrevMafias = []string{}

	for id, p := range l.Players {
		if p.IsActive && p.IsMafia {
			l.PrevMafias = append(l.PrevMafias, p.Nickname)
		}

		mafiaMatched := false
		for _, m := range mafias {
			if m == id {
				mafiaMatched = true
			}
		}
		if mafiaMatched {
			p.IsMafia = true
		} else {
			p.IsMafia = false
		}
		p.IsActive = true
	}

	l.Round++
	return nil
}

type Player struct {
	 Nickname  string
	 IsMafia   bool
	 IsActive  bool
	 IsCreator bool
}

func NewLobby() (*Lobby, int) {
	lobby := Lobby{}
	lobby.Players = map[string]*Player{}

	lobbyID := lobbyCount + 1
	Lobbies[lobbyID] = &lobby
	lobbyCount++

	return &lobby, lobbyID
}