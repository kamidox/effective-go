// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math/rand"
)

// 掷骰子游戏，其规则如下
//   -> 总共有 2 个玩家，轮流掷骰子，谁的得分先超过 100，就获胜
//   -> 游戏开始时，由一个玩家先开始掷骰子
//      -> 如果掷到的骰子是 1 点，则清空当前的得分，并把骰子控制权转给另外一个玩家
//      -> 如果掷到的骰子不是 1 点，则把当前掷出来的骰子点数累加上去，并且骰子控制权依然在当前玩家手上。此时，玩家可以有 2 个策略可以选择：
//         -> 继续掷骰子，则按照上述规则。如果掷到 1 点，分数归零，控制权转移，否则掷出来的点数继续累加。
//         -> 拿走分数，骰子控制权交给另外一个玩家。
const (
	win            = 100 // The winning score in a game of Pig
	gamesPerSeries = 100 // The number of games per series to simulate
)

// A score includes scores accumulated in previous turns for each player,
// as well as the points scored by the current player in this turn.
// player 是当前玩家的分数
// opponent 是对手的分数
// thisTurn 是当前玩家控制的，骰子掷出的分数总和。比如第一次掷了 2 点，再次掷了 5 点后，就是 7 点，此时玩家可以决定继续掷还是拿走分数。
type score struct {
	player, opponent, thisTurn int
}

// An action transitions stochastically to a resulting score.
// 类似 C 语言里的函数指针。这是个策略执行函数，输入是当前的分数，输出是结果分数，以及是否切换控制权的 turnIsOver 标识。
type action func(current score) (result score, turnIsOver bool)

// roll returns the (result, turnIsOver) outcome of simulating a die roll.
// If the roll value is 1, then thisTurn score is abandoned, and the players'
// roles swap.  Otherwise, the roll value is added to thisTurn.
// roll 函数是一个策略，它掷骰子：
//   -> 如果掷到 1 点，则丢失当前分数，并且骰子控制权转移给对方；
//   -> 否则，就可以得到这个分数，并且依然有决定权，是继续掷骰子还是拿走分数。
func roll(s score) (score, bool) {
	outcome := rand.Intn(6) + 1 // A random int in [1, 6]
	if outcome == 1 {
		return score{s.opponent, s.player, 0}, true
	}
	return score{s.player, s.opponent, outcome + s.thisTurn}, false
}

// stay returns the (result, turnIsOver) outcome of staying.
// thisTurn score is added to the player's score, and the players' roles swap.
// stay 函数是一个策略，它拿走当前累积的分数，然后开始让对手掷骰子。
func stay(s score) (score, bool) {
	return score{s.opponent, s.player + s.thisTurn, 0}, true
}

// A strategy chooses an action for any given score.
// strategy 是策略函数，根据输入的 score 返回 action 函数
type strategy func(score) action

// stayAtK returns a strategy that rolls until thisTurn is at least k, then stays.
// 这就是高阶函数 higher-order function。
// stayAtK 函数根据根据输入的参数 k 选择策略函数，它返回的是一个匿名函数。
// 这个匿名函数本身，把 thisTurn 的点数和 k 相比较，
//   -> 如果 thisTurn 点数 >= k，则返回 stay 函数，stay 函数会执行拿走分数的策略。
//   -> 如果 thisTurn < k，则返回 roll 函数，roll 函数会转动骰子，只要没有转到 1 点，就有机会再转动骰子或拿走分数。
// 总结: stayAtK 返回一个函数，这个函数根据当前分数和 k 值比较，再返回策略函数。这里面，涉及了三阶的函数。
func stayAtK(k int) strategy {
	return func(s score) action {
		if s.thisTurn >= k {
			return stay
		}
		return roll
	}
}

// play simulates a Pig game and returns the winner (0 or 1).
// play 函数模拟两个玩家按照规则玩掷骰子游戏，并返回获胜的玩家。两个玩家可以采用不同的策略，由输入参数 strategy0/1 指定。
func play(strategy0, strategy1 strategy) int {
	strategies := []strategy{strategy0, strategy1}
	var s score
	var turnIsOver bool
	currentPlayer := rand.Intn(2) // Randomly decide who plays first
	for s.player+s.thisTurn < win {
		action := strategies[currentPlayer](s)
		s, turnIsOver = action(s)
		if turnIsOver {
			currentPlayer = (currentPlayer + 1) % 2
		}
	}
	return currentPlayer
}

// roundRobin simulates a series of games between every pair of strategies.
// roundRobin 函数从输入的 strategies 列表里，枚举出两两对弈的策略，每种策略执行 10 次对战，统计胜负。
// 返回的 wins 是一种数组，里面放着每个策略获胜的次数，另外一个返回值 gamesPerStrategy 表示每个策略执行的所有对弈次数
func roundRobin(strategies []strategy) ([]int, int) {
	fmt.Printf("start to simulates %d strategies ...\n", len(strategies))
	wins := make([]int, len(strategies))
	for i := 0; i < len(strategies); i++ {
		for j := i + 1; j < len(strategies); j++ {
			for k := 0; k < gamesPerSeries; k++ {
				winner := play(strategies[i], strategies[j])
				if winner == 0 {
					wins[i]++
				} else {
					wins[j]++
				}
			}
		}
	}
	gamesPerStrategy := gamesPerSeries * (len(strategies) - 1) // no self play
	return wins, gamesPerStrategy
}

// ratioString takes a list of integer values and returns a string that lists
// each value and its percentage of the sum of all values.
// e.g., ratios(1, 2, 3) = "1/6 (16.7%), 2/6 (33.3%), 3/6 (50.0%)"
func ratioString(vals ...int) string {
	total := 0
	for _, val := range vals {
		total += val
	}
	s := ""
	for _, val := range vals {
		if s != "" {
			s += ", "
		}
		pct := 100 * float64(val) / float64(total)
		s += fmt.Sprintf("%d/%d (%0.1f%%)", val, total, pct)
	}
	return s
}

func main() {
	// 总共 100 个策略，即最极端的策略，累计超过 100 分时，才拿走分数，一次性就赢得比赛，这种策略当然也有胜率
	strategies := make([]strategy, win)
	for k := range strategies {
		strategies[k] = stayAtK(k + 1)
	}
	wins, games := roundRobin(strategies)

	for k := range strategies {
		fmt.Printf("Wins, losses staying at k = %3d: %s\n",
			k+1, ratioString(wins[k], games-wins[k]))
	}
}
