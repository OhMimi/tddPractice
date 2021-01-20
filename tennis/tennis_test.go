package tennis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTennisGame_GetScore(t *testing.T) {

	assertInstance := assert.New(t)

	type args struct {
		firstPlayer  Player
		secondPlayer Player
	}

	tests := []struct {
		name     string
		expected string
		args     args
	}{
		// 0:0
		{"Test Love All", "Love All", args{Player{"Tom", 0}, Player{"Elsa", 0}}},
		// 1:0
		{"Test Fifteen Love", "Fifteen Love", args{Player{"Tom", 1}, Player{"Elsa", 0}}},
		// 2:0
		{"Test Thirty Love", "Thirty Love", args{Player{"Tom", 2}, Player{"Elsa", 0}}},
		// 3:0
		{"Test Forty Love", "Forty Love", args{Player{"Tom", 3}, Player{"Elsa", 0}}},
		// 1:1
		{"Test Fifteen All", "Fifteen All", args{Player{"Tom", 1}, Player{"Elsa", 1}}},
		// 2:2
		{"Test Thirty All", "Thirty All", args{Player{"Tom", 2}, Player{"Elsa", 2}}},
		// 3:3
		{"Test Deuce", "Deuce", args{Player{"Tom", 3}, Player{"Elsa", 3}}},
		// 2:1
		{"Test Thirty Fifteen", "Thirty Fifteen", args{Player{"Tom", 2}, Player{"Elsa", 1}}},
		// 3:1
		{"Test Forty Fifteen", "Forty Fifteen", args{Player{"Tom", 3}, Player{"Elsa", 1}}},
		// 3:2
		{"Test Forty Thirty", "Forty Thirty", args{Player{"Tom", 3}, Player{"Elsa", 2}}},
		// 0:1
		{"Test Love Fifteen", "Love Fifteen", args{Player{"Tom", 0}, Player{"Elsa", 1}}},
		// 0:2
		{"Test Love Thirty", "Love Thirty", args{Player{"Tom", 0}, Player{"Elsa", 2}}},
		// 0:3
		{"Test Love Forty", "Love Forty", args{Player{"Tom", 0}, Player{"Elsa", 3}}},
		// 1:2
		{"Test Fifteen Thirty", "Fifteen Thirty", args{Player{"Tom", 1}, Player{"Elsa", 2}}},
		// 1:3
		{"Test Fifteen Forty", "Fifteen Forty", args{Player{"Tom", 1}, Player{"Elsa", 3}}},
		// 2:3
		{"Test Thirty Forty", "Thirty Forty", args{Player{"Tom", 2}, Player{"Elsa", 3}}},
		// 4:3
		{"Test FirstPlayer Advantage", "Tom Advantage", args{Player{"Tom", 4}, Player{"Elsa", 3}}},
		// 3:4
		{"Test SecondPlayer Advantage", "Elsa Advantage", args{Player{"Tom", 3}, Player{"Elsa", 4}}},
		// 5:3
		{"Test FirstPlayer Win", "Tom Win", args{Player{"Tom", 5}, Player{"Elsa", 3}}},
		// 4:0
		{"Test FirstPlayer Win", "Tom Win", args{Player{"Tom", 4}, Player{"Elsa", 0}}},
		// 3:5
		{"Test SecondPlayer Win", "Elsa Win", args{Player{"Tom", 3}, Player{"Elsa", 5}}},
		// 0:4
		{"Test SecondPlayer Win", "Elsa Win", args{Player{"Tom", 0}, Player{"Elsa", 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tg := &TennisGame{tt.args.firstPlayer, tt.args.secondPlayer}
			assertInstance.Equal(tt.expected, tg.GameScore())
		})
	}
}
