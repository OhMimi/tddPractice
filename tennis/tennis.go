package tennis

import (
	"fmt"
	"math"
)

var scoreLookup = map[int]string{
	0: "Love",
	1: "Fifteen",
	2: "Thirty",
	3: "Forty",
}

type Player struct {
	Name  string
	Score int
}

type TennisGame struct {
	FirstPlayer  Player
	SecondPlayer Player
}

func (t *TennisGame) GameScore() string {
	if t.isSameScore() {
		if t.FirstPlayer.Score >= 3 {
			return "Deuce"
		}
		return t.scoreIsAll()
	}

	if t.isAdvantage() {
		return t.playerAdvantage()
	}

	if t.isWin() {
		return t.playerWin()
	}

	return t.lookupScore()
}

func (t *TennisGame) getPlayerName() string {
	if t.FirstPlayer.Score > t.SecondPlayer.Score {
		return t.FirstPlayer.Name
	}
	return t.SecondPlayer.Name
}

func (t *TennisGame) isAdvantage() bool {
	return (t.FirstPlayer.Score >= 3 && t.SecondPlayer.Score >= 3) && t.getScoreDiff() == 1
}

func (t *TennisGame) isWin() bool {
	return (t.FirstPlayer.Score >= 4 || t.SecondPlayer.Score >= 4) && t.getScoreDiff() >= 2
}

func (t *TennisGame) isSameScore() bool {
	return t.FirstPlayer.Score == t.SecondPlayer.Score
}

func (t *TennisGame) lookupScore() string {
	return fmt.Sprintf("%s %s", scoreLookup[t.FirstPlayer.Score], scoreLookup[t.SecondPlayer.Score])
}

func (t *TennisGame) playerAdvantage() string {
	return fmt.Sprintf("%s Advantage", t.getPlayerName())
}

func (t *TennisGame) playerWin() string {
	return fmt.Sprintf("%s Win", t.getPlayerName())
}

func (t *TennisGame) scoreIsAll() string {
	return fmt.Sprintf("%s All", scoreLookup[t.FirstPlayer.Score])
}

func (t *TennisGame) getScoreDiff() int {
	return int(math.Abs(float64(t.FirstPlayer.Score - t.SecondPlayer.Score)))
}
