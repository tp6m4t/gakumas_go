# 卡片定義

* Idol Card 角色卡片(要訓練的偶像)
* Support Card 支援卡(養成時可選六張提供支援)
* Memories Card 回憶卡(養成後生成)
* Skill Card 技能卡(戰鬥時使用)
  
| 欄位名稱      | 資料型別 | 說明                     | Other |
| ------------- | -------- | ------------------------ | ----- |
| `ID`          | `string` | 卡片唯一識別碼（ID）     |       |
| `Name`        | `string` | 卡片名稱                 |       |
| `Description` | `string` | 技能或效果描述           |       |
| `Cost`        | `int`    | 消耗體力（優先扣除護盾） |       |
| `CostHealth`  | `int`    | 消耗體力（直接扣除體力） |       |
| `Score`       | `int`    | 分數                     |       |
