// pages/page4.go
package pages

import (
	"fmt"

	"LocalProject/internal/TUI/models"

	tea "github.com/charmbracelet/bubbletea"
)

type page4Model struct {
	models.BaseModel
	input string
}

func NewPage4Model() tea.Model {
	return &page4Model{
		input: "",
	}
}

func (m page4Model) Init() tea.Cmd {
	return nil
}

func (m page4Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return NewHomeModel(), nil
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			m.input += msg.String()
		}
	case tea.WindowSizeMsg:
		m.BaseModel.Width = msg.Width
		m.BaseModel.Height = msg.Height
	}

	return m, nil
}

func (m page4Model) View() string {
	return fmt.Sprintf(
		"頁面4\n\n輸入內容: %s\n\n按 q 返回主選單\n",
		m.input,
	)
}
