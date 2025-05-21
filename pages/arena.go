package pages

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/taipi/kakumasu_CLI/models"
)

type arenaModel struct {
	models.BaseModel
	childMenu     []string // 子選單項目
	childCursor   int      // 子選單游標 (垂直導航)
	siblingCursor int      // 兄弟選單游標 (水平導航)
	isSiblingNav  bool     // 是否處於兄弟導航模式
}

func NewArenaModel() tea.Model {
	return &arenaModel{
		childMenu:     []string{"偶像之路", "競技場"},
		childCursor:   0,
		siblingCursor: -1,
		isSiblingNav:  false,
	}
}

func (m arenaModel) Init() tea.Cmd {
	return nil
}

func (m arenaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			m.isSiblingNav = false
			if m.childCursor > 0 {
				m.childCursor--
			}
		case "down", "j":
			m.isSiblingNav = false
			if m.childCursor < len(m.childMenu)-1 {
				m.childCursor++
			}
		case "left", "h":
			m.isSiblingNav = true
			if m.siblingCursor > 0 {
				m.siblingCursor--
			} else {
				m.siblingCursor = len(SharedNavItems) - 1
			}
		case "right", "l":
			m.isSiblingNav = true
			if m.siblingCursor < len(SharedNavItems)-1 {
				m.siblingCursor++
			} else {
				m.siblingCursor = 0
			}
		case "enter":
			if m.isSiblingNav {
				// 確保 siblingCursor 在 SharedNavItems 的有效範圍內
				if m.siblingCursor >= 0 && m.siblingCursor < len(SharedNavItems) {
					return SharedNavItems[m.siblingCursor].Open(), nil // 從全局 SharedNavItems 打開頁面
				}
				return m, nil // 如果游標無效，不執行任何操作
			} else {
				switch m.childCursor {
				//尚未實作功能
				default:
					return m, nil
				}
			}
		}
	case tea.WindowSizeMsg:
		m.BaseModel.Width = msg.Width
		m.BaseModel.Height = msg.Height
	}

	return m, nil
}

func (m arenaModel) View() string {
	s := "競技場\n\n"

	// 顯示子選項
	for i, choice := range m.childMenu {
		cursor := " "
		if !m.isSiblingNav && m.childCursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\n"
	// 顯示兄弟選項
	for i, choice := range SharedNavItems {
		cursor := " "
		if m.isSiblingNav && m.siblingCursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s", cursor, choice.Name)
	}
	s += "\n按 q 或 ctrl+c 退出\n"
	return s
}
