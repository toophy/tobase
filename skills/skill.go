package skills

// SkillResult 技能使用结果集
type SkillResult struct {
	Bingo     bool
	Attacker  *Role
	Defender  *Role
	PhyDamage int32
	Skill     *SkillInfo
}

// SkillInfo 技能信息
type SkillInfo struct {
	Name    string // 技能名
	PhyAtt  int32  // 物理攻击力
	HitRate int32  // 命中率
}
