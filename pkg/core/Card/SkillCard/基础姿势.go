package SkillCard

type poseBasics struct {
	SkillCardBase
}

func NewPoseBasics() *poseBasics {
	c := &poseBasics{}
	c.SkillCardBase.Set(
		2, 0, 0, 0,
		"基础姿势",
		"A", []int{2}, 2, "", 0, "数值+2 活力+2", false,
	)
	return c
}

func (c poseBasics) Use(f field) {
	f.AddScore(c.BaseScore[0])
}

func (c poseBasics) Upgrade() {}
