package SkillCard

type field interface {
	AddScore(value int)
}

type SkillCard struct {
	ID          int           //卡片唯一識別碼（ID
	Plvl        int           //解鎖等級 -1表示來源為支援卡
	Plan        int           //ALL(0) 感性(1) 理性(2) 非凡(3)
	Rarity      int           //稀有度 N(0) R(1) SR(2) SSR(3)
	Name        string        //卡片名稱
	Description string        //技能或效果描述
	Cost        Cost          //消耗
	Score       int           //分數
	Use         func(f field) //使用卡牌
}

type Cost struct {
	Value int    //消耗多少
	Type  string //消耗甚麼
}

var AppealBasics = SkillCard{
	ID:          0,
	Plvl:        0,
	Plan:        0,
	Rarity:      0,
	Name:        "基础表演",
	Description: "数值+9",
	Score:       9,
	Cost: Cost{
		Type:  "Energy",
		Value: 6,
	},
}

func init() {
	AppealBasics.Use = func(f field) {
		f.AddScore(AppealBasics.Score)
	}
	//var SkillCardList = [...]SkillCard{AppealBasics, PoseBasics}
	/*
	   var PoseBasics = SkillCard{}基础姿势

	   var ExpressionBasics={}基础表现力

	   var Trouble = {} 睡意*/
}
