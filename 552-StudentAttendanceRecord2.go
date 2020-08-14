package main

import (
	"fmt"
	"math"
)

func checkRecord2(n int) int {
	states := []byte{'A', 'L', 'P'}
	var recs []string

	total := int(math.Pow(3, float64(n)))

	c := 0
	ci := 0
	for i := 0; i < total; i++ {
		if ci >= total/3 {
			c++
			ci = 0
		}
		ci++
		recs = append(recs, string(states[c]))
	}

	fmt.Println(recs)

	/*for i := 0; i < len(states); i++ {
		for j := 0; j < len(states); j++ {
			recs = append(recs, string(states[i]) + string(states[j]))
		}
	}*/

	count := 0

	for i := 0; i < len(recs); i++ {
		// max answer size
		if count > int(math.Pow(10, 9)+7) {
			break
		}
		// check if rewardable by using our code from #551
		if checkRecord(recs[i]) {
			count++
		}
	}

	return count
}

/*

n   recs  count
1 => 3   => 3
2 => 9   => 8
3 => 27  => 19
4 => 81  => 43
5 => 243 => 94

27/9 = 3

27-3 = 24

*/
