package structs

import "time"

type MatchDetail struct {
	SeasonId      interface{} `json:"seasonId"`
	MatchResult   string      `json:"matchResult"`
	MatchEndType  int         `json:"matchEndType"`
	SystemPause   int         `json:"systemPause"`
	Foul          int         `json:"foul"`
	Injury        int         `json:"injury"`
	RedCards      int         `json:"redCards"`
	YellowCards   int         `json:"yellowCards"`
	Dribble       int         `json:"dribble"`
	CornerKick    int         `json:"cornerKick"`
	Possession    int         `json:"possession"`
	OffsideCount  int         `json:"offsideCount"`
	AverageRating float64     `json:"averageRating"`
	Controller    string      `json:"controller"`
}

type Shoot struct {
	ShootTotal          int `json:"shootTotal"`
	EffectiveShootTotal int `json:"effectiveShootTotal"`
	ShootOutScore       int `json:"shootOutScore"`
	GoalTotal           int `json:"goalTotal"`
	GoalTotalDisplay    int `json:"goalTotalDisplay"`
	OwnGoal             int `json:"ownGoal"`
	ShootHeading        int `json:"shootHeading"`
	GoalHeading         int `json:"goalHeading"`
	ShootFreekick       int `json:"shootFreekick"`
	GoalFreekick        int `json:"goalFreekick"`
	ShootInPenalty      int `json:"shootInPenalty"`
	GoalInPenalty       int `json:"goalInPenalty"`
	ShootOutPenalty     int `json:"shootOutPenalty"`
	GoalOutPenalty      int `json:"goalOutPenalty"`
	ShootPenaltyKick    int `json:"shootPenaltyKick"`
	GoalPenaltyKick     int `json:"goalPenaltyKick"`
}

type ShootDetail struct {
	GoalTime   int     `json:"goalTime"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	Type       int     `json:"type"`
	Result     int     `json:"result"`
	SpId       int     `json:"spId"`
	SpGrade    int     `json:"spGrade"`
	SpLevel    int     `json:"spLevel"`
	SpIdType   bool    `json:"spIdType"`
	Assist     bool    `json:"assist"`
	AssistSpId int     `json:"assistSpId"`
	AssistX    float64 `json:"assistX"`
	AssistY    float64 `json:"assistY"`
	HitPost    bool    `json:"hitPost"`
	InPenalty  bool    `json:"inPenalty"`
}

type Pass struct {
	PassTry                  int `json:"passTry"`
	PassSuccess              int `json:"passSuccess"`
	ShortPassTry             int `json:"shortPassTry"`
	ShortPassSuccess         int `json:"shortPassSuccess"`
	LongPassTry              int `json:"longPassTry"`
	LongPassSuccess          int `json:"longPassSuccess"`
	BouncingLobPassTry       int `json:"bouncingLobPassTry"`
	BouncingLobPassSuccess   int `json:"bouncingLobPassSuccess"`
	DrivenGroundPassTry      int `json:"drivenGroundPassTry"`
	DrivenGroundPassSuccess  int `json:"drivenGroundPassSuccess"`
	ThroughPassTry           int `json:"throughPassTry"`
	ThroughPassSuccess       int `json:"throughPassSuccess"`
	LobbedThroughPassTry     int `json:"lobbedThroughPassTry"`
	LobbedThroughPassSuccess int `json:"lobbedThroughPassSuccess"`
}

type Defence struct {
	BlockTry      int `json:"blockTry"`
	BlockSuccess  int `json:"blockSuccess"`
	TackleTry     int `json:"tackleTry"`
	TackleSuccess int `json:"tackleSuccess"`
}

type Status struct {
	Shoot                int     `json:"shoot"`
	EffectiveShoot       int     `json:"effectiveShoot"`
	Assist               int     `json:"assist"`
	Goal                 int     `json:"goal"`
	Dribble              int     `json:"dribble"`
	Intercept            int     `json:"intercept"`
	Defending            int     `json:"defending"`
	PassTry              int     `json:"passTry"`
	PassSuccess          int     `json:"passSuccess"`
	DribbleTry           int     `json:"dribbleTry"`
	DribbleSuccess       int     `json:"dribbleSuccess"`
	BallPossesionTry     int     `json:"ballPossesionTry"`
	BallPossesionSuccess int     `json:"ballPossesionSuccess"`
	AerialTry            int     `json:"aerialTry"`
	AerialSuccess        int     `json:"aerialSuccess"`
	BlockTry             int     `json:"blockTry"`
	Block                int     `json:"block"`
	TackleTry            int     `json:"tackleTry"`
	Tackle               int     `json:"tackle"`
	YellowCards          int     `json:"yellowCards"`
	RedCards             int     `json:"redCards"`
	SpRating             float64 `json:"spRating"`
}

type Player struct {
	SpId       int    `json:"spId"`
	SpPosition int    `json:"spPosition"`
	SpGrade    int    `json:"spGrade"`
	Status     Status `json:"status"`
}

type MatchInfo struct {
	Ouid        string        `json:"ouid"`
	Nickname    string        `json:"nickname"`
	MatchDetail MatchDetail   `json:"matchDetail"`
	Shoot       Shoot         `json:"shoot"`
	ShootDetail []ShootDetail `json:"shootDetail"`
	Pass        Pass          `json:"pass"`
	Defence     Defence       `json:"defence"`
	Player      []Player      `json:"player"`
}

type MatchDetailResponse struct {
	MatchId   string      `json:"matchId"`
	MatchDate CustomTime  `json:"matchDate"`
	MatchType int         `json:"matchType"`
	MatchInfo []MatchInfo `json:"matchInfo"`
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(`"2006-01-02T15:04:05"`, s) // TODO-FINDOUT: parsing work like this? wow
	return
}
