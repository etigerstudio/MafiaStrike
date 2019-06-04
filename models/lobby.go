package models

import "mafia-strike/util"

var Lobbies = map[int]Lobby{}
var lobbyCount = 0

type Lobby struct{
	Players map[string]Player
	Keywords []string
	CurrentWord string
	PrevWord string
	Round int
}

func (l *Lobby) PlayerCount() int {
	return len(l.Players)
}

func (l *Lobby) AddPlayer(nickname string) string {
	playerID := util.GenerateRandomShortID()
	l.Players[playerID] = Player{
		Nickname: nickname,
	}
	return playerID
}

type Player struct {
	 Nickname string
	 IsMafia bool
	 Active bool
}

func NewLobby() (*Lobby, int) {
	lobby := Lobby{}
	lobby.Players = map[string]Player{}

	lobbyID := lobbyCount + 1
	Lobbies[lobbyID] = lobby
	lobbyCount++

	return &lobby, lobbyID
}