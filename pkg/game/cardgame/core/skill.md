# 代碼解釋
``` go
/*
技能
Condition 發動條件描述
Effect 效果描述

條件完全符合時,才發動所有效果 tips:為空時視為符合條件
*/
type Skill struct {
	Condition []Condition
	Effect    []Effect
}

// 條件 
// ex. 
// Label    := "Health"
// Operator := "="
// Value    := "10"
//
// 此時 Health = 10 即為符合條件
type Condition struct {
    Label    string    //比較目標
	Operator string    //比較運算符
	Value    string    //比較值
}

// 效果
// ex.
// Label := "Impression"
// Operator := "+="
// Value := "10"
// 效果為 好印象 += 10
type Effect struct {
    Label    string    //效果目標
	Operator string    //附值運算符
	Value    string    //效果值
}
```

## Condition 成員參數限制
>### Label

>### Operator
* =
* <
* \>
* <=
* \>=
* !=
>### Value
* 通常為數字 也支援標籤
## Effect 成員參數限制
>### Label/Value
和Condition相同

>### Operator
* =
* +=
* -=
* *=
* /=