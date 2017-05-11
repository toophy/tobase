package skills

import (
	"math/rand"
	"testing"
	"time"
)

func TestSingleSkill(t *testing.T) {
	rA := &Role{"A", 100, 5, 3, false}
	rB := &Role{"B", 100, 3, 5, false}
	sKA := &SkillInfo{"神龙摆尾", 35, 30}
	sKB := &SkillInfo{"打狗棒法", 28, 40}

	rand.Seed(time.Now().Unix())

	for i := 0; i < 40; i++ {

		hit := myRand(1, 100)
		if hit < 50 {
			hit = myRand(1, 100)
			if hit < 50 {
				ret := rA.UseSkill(sKA, rB)
				rB.ApplySkillResult(ret)
			} else {
				ret := rA.UseSkill(sKB, rB)
				rB.ApplySkillResult(ret)
			}
		} else {
			hit = myRand(1, 100)
			if hit < 50 {
				ret := rB.UseSkill(sKA, rA)
				rA.ApplySkillResult(ret)
			} else {
				ret := rB.UseSkill(sKB, rA)
				rA.ApplySkillResult(ret)
			}
		}

		if rA.Dead || rB.Dead {
			break
		}
	}
}
