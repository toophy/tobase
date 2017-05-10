package tobase

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

type Role struct {
	Name   string
	Hp     int
	PhyAtt int
	PhyDef int
	Dead   bool
}

type SkillResult struct {
	Bingo     bool
	Attacker  *Role
	Defender  *Role
	PhyDamage int
	Skill     *SkillInfo
}

type SkillInfo struct {
	Name    string // 技能名
	PhyAtt  int    // 物理攻击力
	HitRate int    // 命中率
}

func (r *Role) UseSkill(s *SkillInfo, t *Role) *SkillResult {

	ret := &SkillResult{}
	ret.Skill = s
	ret.Attacker = r
	ret.Defender = t

	if r.Dead {
		return ret
	}

	if r.Name != t.Name {
		if !t.Dead {
			hit := rand.Intn(100)
			log.Printf("hit = %d", hit)
			if hit <= s.HitRate {
				if r.PhyAtt+s.PhyAtt > t.PhyDef {
					ret.Bingo = true
					ret.PhyDamage = r.PhyAtt + s.PhyAtt - t.PhyDef
				}
			}
		}
	} else {
	}

	return ret
}

func (r *Role) ApplySkillResult(s *SkillResult) {
	if !s.Bingo {
		log.Printf("%s施展一招%s,%s成功闪避", s.Attacker.Name, s.Skill.Name, s.Defender.Name)
		return
	}

	if r.Hp > s.PhyDamage {
		r.Hp -= s.PhyDamage
		log.Printf("%s施展一招%s,打的%s掉了%d血", s.Attacker.Name, s.Skill.Name, s.Defender.Name, s.PhyDamage)
	} else {
		r.Hp = 0
	}

	if r.Hp == 0 {
		r.Dead = true
		log.Printf("%s施展一招%s,%s体力不支,倒下了", s.Attacker.Name, s.Skill.Name, s.Defender.Name)
	}
}

func TestRand(t *testing.T) {
	randx := func(min, max int32) int32 {
		o := rand.Int31()
		x := float64(o) / float64(2147483647)
		y := (float64(max) - float64(min)) * x
		return min + int32(y)
	}

	for i := 0; i < 100; i++ {
		hit := randx(1, 10)
		log.Printf("hit =\t%d", hit)
	}
}

func TestSingleSkill(t *testing.T) {
	rA := &Role{"A", 100, 5, 3, false}
	rB := &Role{"B", 100, 3, 5, false}
	sKA := &SkillInfo{"神龙摆尾", 35, 60}
	sKB := &SkillInfo{"打狗棒法", 28, 70}

	rand.Seed(time.Now().Unix())

	for i := 0; i < 20; i++ {

		hit := rand.Intn(2)
		log.Printf("2hit = %d", hit)
		if hit < 1 {
			hit = rand.Intn(2)
			log.Printf("2hit = %d", hit)
			if hit < 1 {
				ret := rA.UseSkill(sKA, rB)
				rB.ApplySkillResult(ret)
			} else {
				ret := rA.UseSkill(sKB, rB)
				rB.ApplySkillResult(ret)
			}
		} else {
			hit = rand.Intn(2)
			log.Printf("2hit = %d", hit)
			if hit < 1 {
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
