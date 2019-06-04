package consts

const (
	LogPrefix   = "[MafiaStrike]"
	LogInfo     = "[Info]"
	LogWarning  = "[Warning]"
	LogError    = "[Error]"
)

const (
	RequestQueryPlayerID       = "player_id"
)

const (
	RequestPostFormNickname    = "nickname"
	RequestPostFormAction      = "action"
	RequestPostFormPlayerID    = "player_id"
	RequestPostFormKeywords    = "keywords"
	RequestPostFormWinner      = "winner"
)

const (
	LobbyPatchActionAddPlayer       = "add_player"
	LobbyPatchActionNextRound       = "next_round"
	LobbyPatchActionUpdateKeywords  = "update_keywords"
	LobbyPatchActionSubmitResult    = "submit_result"
)

const (
	GameWinnerMen                   = "men"
	GameWinnerMafias                = "mafias"
	GameWinnerDraw                  = "draw"
)

const (
	RequestParamLobby          = "lobby"
)

const (
	ErrorUniversal             = "出错啦"
)

const (
	ResponseKeyError                = "error"
	ResponseKeyPlayerID             = "player_id"
	ResponseKeyLobbyID              = "lobby_id"
)

const (
	ResponseErrorDescNoKeyword      = "no-keyword"
)

const (
	StringKeywordsSeparator         = " "
	StringHTMLLineBreak             = "<br />"
)