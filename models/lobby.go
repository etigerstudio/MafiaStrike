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
	IsRoundEnded bool
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
	l.IsRoundEnded = false
	return nil
}

func (l *Lobby) SubmitResult(winner string) {
	switch winner {
	case consts.GameWinnerMen:
		for _, p := range l.Players {
			if !p.IsMafia {
				p.Score++
			}
		}
	case consts.GameWinnerMafias:
		for _, p := range l.Players {
			if p.IsMafia {
				p.Score++
			}
		}
	case consts.GameWinnerDraw:
		for _, p := range l.Players {
			p.Score++
		}
	}
	l.IsRoundEnded = true
}

type Player struct {
	 Nickname  string
	 IsMafia   bool
	 IsActive  bool
	 IsCreator bool
	 Score int
}

func NewLobby() (*Lobby, int) {
	lobby := Lobby{}
	lobby.Players = map[string]*Player{}
	//lobby.Players = map[string]*Player{
	//	"1": {Nickname: "a"},
	//	"2": {Nickname: "b"},
	//	"3": {Nickname: "c"},
	//	"4": {Nickname: "d"},
	//	"5": {Nickname: "e"},
	//	"6": {Nickname: "f"},
	//}
	lobby.IsRoundEnded = true

	lobbyID := lobbyCount + 1
	Lobbies[lobbyID] = &lobby
	lobbyCount++

	return &lobby, lobbyID
}