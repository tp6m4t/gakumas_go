package models

import tea "github.com/charmbracelet/bubbletea"

// PageMsg 用於頁面間導航
type PageMsg int

const (
	HomePage PageMsg = iota
	Page1
	Page2
	Page3
)

// BaseModel 是所有頁面的基礎模型
type BaseModel struct {
	Width  int
	Height int
}

// 所有頁面都需要實現的接口
type Page interface {
	tea.Model
	View() string
}
