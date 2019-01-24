// https://atcoder.jp/contests/pakencamp-2018-day3/tasks/pakencamp_2018_day3_d
package main

import (
	"fmt"
	"os"
)

type MovingRewards struct {
	acc        []int64
	numOfRooms int
}

type RoomNumber int

func NewMovingRewards(C []int64, N int) *MovingRewards {
	acc := make([]int64, N) //累計利得．acc[0] = 0, acc[1] = C1, acc[2] = acc[1] + C2, ...
	acc[0] = 0
	for i := 1; i < N; i++ {
		// acc[0] = 1 <-> 1への移動利得 = 0
		// acc[1] = 1 <-> 2への移動利得 = C[0]
		// acc[2] = 1 <-> 3への移動利得 = C[0] + C[1] = acc[1] + C[1]
		// acc[3] = 1 <-> 4への移動利得 = C[0] + C[1] + C[2] = acc[2] + C[2]
		// ...
		// acc[N-1] = 1 <-> Nへの移動利得 = C[0] + ... + C[N-2] = acc[N-2] + C[N-2]
		acc[i] = acc[i-1] + C[i-1]
	}
	return &MovingRewards{
		acc:        acc,
		numOfRooms: N,
	}
}

func (m *MovingRewards) fromTo(from RoomNumber, to RoomNumber) int64 {
	roomNumberToACCRewardsIndex := func(number RoomNumber) int {
		return int(number - 1)
	}

	f, t := from, to
	if f > t {
		f, t = to, from
	}
	fIdx, tIdx := roomNumberToACCRewardsIndex(f), roomNumberToACCRewardsIndex(t)
	return m.acc[tIdx] - m.acc[fIdx]
}

func (m *MovingRewards) CalcMaxRewardOf(from RoomNumber) int64 {
	max := int64(0)
	for i := from; i <= RoomNumber(m.numOfRooms); i++ {
		r := m.fromTo(from, RoomNumber(i))
		if r > max {
			max = r
		}
	}
	for i := from; i >= RoomNumber(1); i-- {
		r := m.fromTo(from, RoomNumber(i))
		if r > max {
			max = r
		}
	}
	return max
}
func main() {
	N := 0 //部屋数
	fmt.Fscanf(os.Stdin, "%d", &N)

	// C[0] = 部屋1<->部屋2の移動利得, C[1] = 部屋2<->部屋3の移動利得, ..., C[N-2] = 部屋N-1<->部屋Nの移動利得
	C := make([]int64, N-1)
	for i := 0; i < N-1; i++ {
		fmt.Fscanf(os.Stdin, "%d", &(C[i]))
	}

	// X := make([]int64, N)
	movingRewards := NewMovingRewards(C, N)
	for i := 1; i <= N; i++ {
		max := movingRewards.CalcMaxRewardOf(RoomNumber(i))
		fmt.Println(max)
	}
}
