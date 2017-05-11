package skills

import "math/rand"

// myRand 略加修改的随机数
func myRand(min, max int32) int32 {
	if max == min {
		return min
	} else if max < min {
		temp := max
		max = min
		min = temp
	}

	return min + int32((float64(max)-float64(min)+1)*float64(rand.Int63())/(1<<63+1))
}
