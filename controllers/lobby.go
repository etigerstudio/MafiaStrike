package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mafia-strike/consts"
	"mafia-strike/models"
	"mafia-strike/util"
	"net/http"
	"strings"
)

func PostLobbyEntry(c *gin.Context) {
	nickname := util.MustGetPostForm(consts.RequestPostFormNickname ,c)

	lobby, lobbyID := models.NewLobby()
	playerID := lobby.AddPlayer(nickname, true)
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/lobbies/%d?player_id=%s", lobbyID, playerID))
}

func PatchLobbyEntry(c *gin.Context) {
	lobbyID := util.MustGetParamInt(consts.RequestParamLobby ,c)
	action := util.MustGetPostForm(consts.RequestPostFormAction ,c)

	lobby, ok := models.Lobbies[lobbyID]
	if !ok {
		util.ErrorMessageUniversal(c)
		return
	}

	switch action {
	case consts.LobbyPatchActionAddPlayer:
		nickname := util.MustGetPostForm(consts.RequestPostFormNickname, c)

		playerID := lobby.AddPlayer(nickname, false)
		c.JSON(http.StatusOK, gin.H{"player_id": playerID})
		return
	case consts.LobbyPatchActionNextRound:
		playerID := util.MustGetPostForm(consts.RequestPostFormPlayerID, c)

		player, ok := lobby.Players[playerID]
		if !ok {
			util.ErrorMessageUniversal(c)
			return
		}

		if !player.IsCreator {
			util.ErrorMessageUniversal(c)
			return
		}

		err := lobby.StartNewRound()
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{consts.ResponseKeyError: consts.ResponseErrorDescNoKeyword})
			return
		}
		c.String(http.StatusOK, "")
		return
	case consts.LobbyPatchActionUpdateKeywords:
		keywords := util.MustGetPostForm(consts.RequestPostFormKeywords, c)

		lobby.Keywords = strings.Split(keywords, consts.LobbyKeywordSeparator)
		c.String(http.StatusOK, "")
		return
	}

	util.ErrorMessageUniversal(c)
}

func GetLobbyEntry(c *gin.Context) {
	lobbyID := util.MustGetParamInt(consts.RequestParamLobby ,c)
	playerID := util.MustGetQuery(consts.RequestQueryPlayerID,c)

	lobby, ok := models.Lobbies[lobbyID]
	if !ok {
		util.ErrorMessageUniversal(c)
		return
	}

	player, ok := lobby.Players[playerID]
	if !ok {
		util.ErrorMessageUniversal(c)
		return
	}

	params := gin.H{}
	if player.IsActive {
		params["waiting_class"] = classForShouldHide(true)
		params["innocent_class"] = classForShouldHide(player.IsMafia)
		params["mafia_class"] = classForShouldHide(!player.IsMafia)
		params["innocent_count"] = lobby.PlayerCount() - lobby.CurrentMafiaCount()
		params["mafia_count"] = lobby.CurrentMafiaCount()
		if player.IsMafia {
			params["mafia_keyword"] = lobby.CurrentWord
		}
	} else {
		params["waiting_class"] = classForShouldHide(false)
		params["innocent_class"] = classForShouldHide(true)
		params["mafia_class"] = classForShouldHide(true)
		params["count_class"] = classForShouldHide(true)
	}
	params["round_number"] = lobby.Round
	params["player_count"] = lobby.PlayerCount()
	params["creator_class"] = classForShouldHide(!player.IsCreator)
	params["player_id"] = playerID
	params["lobby_id"] = lobbyID
	params["keyword_list"] = strings.Join(lobby.Keywords, consts.LobbyKeywordSeparator)
	c.HTML(http.StatusOK, "lobby.tmpl", params)
}

func classForShouldHide(hidden bool) string {
	if hidden {
		return "hide"
	} else {
		return ""
	}
}