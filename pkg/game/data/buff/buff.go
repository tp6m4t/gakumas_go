package buff

type Buff interface {
	IsDebuff() bool
}

type BaseBuff struct {
	id           int
	name         string
	descriptions string
}

func (BaseBuff) IsDebuff() bool {
	return false
}

type DeBuff struct {
}

func (DeBuff) IsDebuff() bool {
	return false
}

/*
const (
	//職業相關buff
	GoodCondition   // 好調(分數上升量提升50%)
	PerfectForm     // 絕好調(分數上升量提升 好調*10%)
	Concentration   // 集中(數值上升時而外增加數值)
	Impression      // 好印象(回合結束依數值上升分數)
	Motivation      // 幹勁(能量上升時而外增加能量)
	FullPowerPoints // 全力
	Strength        // 堅決
	Preservation    // 溫存
	//悠閒
	//其他bufff
	//體力消費減少
	//卡牌可使用次數數增加
	//免消耗體力
	//應援
	//debuff
	//體力消費增加
	//活力增加無效
	//狀態不佳
	//使用體力追加
	//禁用部分手牌
	//手牌減少
	//心情浮躁 手牌消耗體力隨機變化
	//膽怯 能量增加量-1
	//精神萎靡 數值無法上升
	//不安 能量增加量-33%
	Other // 其他
)

func (b Buff) String() string {
	return [...]string{
		"好調",
		"絕好調",
		"集中",
		"好印象",
		"幹勁",
		"全力",
		"堅決",
		"溫存",
		"其他",
	}[b]
}
*/
