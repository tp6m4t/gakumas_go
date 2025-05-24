package core

/*
	技能
	符合條件時,技能效果 tips:條件為空時視為符合條件
*/
type Skill struct {
	Condition []Condition //發動條件描述(須完全符合)
	Effect    []Effect    //效果描述
}

/*
	條件
	ex.
 		Label    := "Health"
 		Operator := "="
 		Value    := "10"
	此時 Health = 10 即為符合條件
*/
type Condition struct {
	Label    string //比較目標
	Operator string //比較運算符
	Value    string //比較值
}

// 效果
// ex.
// Label := "Impression"
// Operator := "+="
// Value := "10"
// 效果為 好印象 += 10
type Effect struct {
	Label    string //效果目標
	Operator string //附值運算符
	Value    string //效果值
}
