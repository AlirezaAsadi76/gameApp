package entity

import "time"

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	Players     []uint
	StartTime   time.Time
}

type Player struct {
	ID      uint
	UserID  uint
	GameID  uint
	Score   int
	Answers []PlayerAnswer
}

type PlayerAnswer struct {
	ID         uint
	QuestionID uint
	PlayerID   uint
	Choice     PossibleAnswerChoice
}
