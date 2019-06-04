package models

import (
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

func (l *Lobby) StartNewRound() {
	mafiaCount := l.NextMafiaCount()
	util.Infoln("mafia count: ", mafiaCount)

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
	util.Infoln("mafias: ", mafias)

	for id, p := range l.Players {
		mafiaMatched := false
		for _, m := range mafias {
			if m == id {
				mafiaMatched = true
				util.Infoln("new mafia: ", p.Nickname)
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
}

type Player struct {
	 Nickname  string
	 IsMafia   bool
	 IsActive  bool
	 IsCreator bool
}

func NewLobby() (*Lobby, int) {
	lobby := Lobby{}
	lobby.Players = map[string]*Player{
		"dsad": {Nickname: "aa"},
		"nnn": {Nickname: "b"},
		"bbb": {Nickname: "c"},
		"sd": {Nickname: "d"},
	}

	lobbyID := lobbyCount + 1
	Lobbies[lobbyID] = &lobby
	lobbyCount++

	return &lobby, lobbyID
}