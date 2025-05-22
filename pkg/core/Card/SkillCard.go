package Card

type SkillCard struct {
	ID          int    //卡片唯一識別碼（ID
	Name        string //卡片名稱
	Description string //技能或效果描述
	Cost        int    //消耗體力（優先扣除護盾）
	CostHealth  int    //消耗體力（直接扣除體力）
	Score       int    //分數
}
