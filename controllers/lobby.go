package controllers

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"mafia-strike/consts"
	"mafia-strike/models"
	"mafia-strike/util"
	"net/http"
	"strconv"
	"strings"
)

func PostLobbyEntry(c *gin.Context) {
	nickname := util.MustGetPostForm(consts.RequestPostFormNickname ,c)

	lobby, lobbyID := models.NewLobby()
	playerID := lobby.AddPlayer(nickname, true)

	c.JSON(http.StatusOK, gin.H{
		consts.ResponseKeyLobbyID: lobbyID,
		consts.ResponseKeyPlayerID: playerID,
	})
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
		c.JSON(http.StatusOK, gin.H{consts.ResponseKeyPlayerID: playerID})
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

		lobby.Keywords = strings.Split(keywords, consts.StringKeywordsSeparator)
		c.String(http.StatusOK, "")
		return
	case consts.LobbyPatchActionSubmitResult:
		playerID := util.MustGetPostForm(consts.RequestPostFormPlayerID, c)
		winner := util.MustGetPostForm(consts.RequestPostFormWinner, c)

		player, ok := lobby.Players[playerID]
		if !ok {
			util.ErrorMessageUniversal(c)
			return
		}

		if !player.IsCreator {
			util.ErrorMessageUniversal(c)
			return
		}

		lobby.SubmitResult(winner)
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
	if player.IsActive && !lobby.IsRoundEnded {
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
	params["player_nickname"] = player.Nickname
	playerListStr := ""
	for _, p := range lobby.Players {
		playerListStr += `{"nickname":"` + p.Nickname +
			`","score":` + strconv.Itoa(p.Score) + `},`
	}
	params["player_list"] = template.JS("[" + playerListStr + "]")
	if lobby.Round == 0 || lobby.IsRoundEnded {
		params["end_round_class"] = classForShouldHide(true)
	}
	if !lobby.IsRoundEnded {
		params["start_new_round_class"] = classForShouldHide(true)
	}
	if lobby.Round > 1 {
		params["last_keyword"] = lobby.PrevWord
		params["last_mafias"] = template.HTML(strings.Join(lobby.PrevMafias, consts.StringHTMLLineBreak))
	} else {
		params["last_round_class"] = classForShouldHide(true)
	}
	params["lobby_id"] = lobbyID
	params["keyword_list"] = strings.Join(lobby.Keywords, consts.StringKeywordsSeparator)
	c.HTML(http.StatusOK, "lobby.tmpl", params)
}

func classForShouldHide(hidden bool) string {
	if hidden {
		return "hide"
	} else {
		return ""
	}
}