package skills

import (
	crand "crypto/rand"
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func TestRand(t *testing.T) {

	randx := func(min, max int32) int32 {
		if max == min {
			return min
		} else if max < min {
			temp := max
			max = min
			min = temp
		}

		return min + int32((float64(max)-float64(min)+1)*float64(rand.Int63())/(1<<63+1))
	}

	for i := 0; i < 200; i++ {
		hit := randx(1, 10)
		log.Printf("hit =\t%d", hit)
	}
}

//真随机 -- 用标准库封装好的
func a3() {
	b := make([]byte, 16)
	// On Unix-like systems, Reader reads from /dev/urandom.
	// On Windows systems, Reader uses the CryptGenRandom API.
	_, err := crand.Read(b) //返回长度为0 - 32 的值
	if err != nil {
		fmt.Println("[a3] ", err)
		return
	}
	fmt.Println("[a3] b:\t", b)
}

func TestRandNormal(t *testing.T) {
	x := rand.Perm(200)
	for i := 0; i < 200; i++ {
		//hit := rand.Intn(10) + 1
		//log.Printf("hit =\t%d", hit)
		//a3()
		log.Printf("hit =\t%d", x[i])
	}
}
