package skills

import "log"

// Role 角色数据
type Role struct {
	Name     string
	Hp       int32
	PhyAtt   int32
	PhyDef   int32
	HitRate  int32
	MissRate int32
	Dead     bool
}

// UseSkill 对单体目标使用技能
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
			hit := myRand(1, 100)
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

// ApplySkillResult 应用技能结果集到角色上
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
