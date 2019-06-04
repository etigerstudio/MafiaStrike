package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mafia-strike/consts"
	"mafia-strike/models"
	"mafia-strike/util"
	"net/http"
)

func PostLobbyEntry(c *gin.Context) {
	nickname := util.MustGetPostForm(consts.RequestPostFormNickname ,c)

	lobby, lobbyID := models.NewLobby()
	playerID := lobby.AddPlayer(nickname)
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/lobbies/%d?player=%s", lobbyID, playerID))
}

func GetLobbyEntry(c *gin.Context) {
	lobbyID := util.MustGetParamInt(consts.RequestParamLobby ,c)
	playerID := util.MustGetQuery(consts.RequestQueryPlayer ,c)

	lobby, ok := models.Lobbies[lobbyID]
	if !ok {
		util.ErrorMessageUniversal(c)
	}

	_, ok = lobby.Players[playerID]
	if !ok {
		util.ErrorMessageUniversal(c)
	}

	c.Status(http.StatusOK)
}