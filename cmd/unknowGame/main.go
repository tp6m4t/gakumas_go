package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Game state
type model struct {
	location        string
	inventory       []string
	health          int
	world           map[string]string
	items           map[string]string
	err             error
	waitingForInput string // Added this field to track input state
}

// Initialize model
func initialModel() model {
	return model{
		location:  "forest",
		inventory: []string{},
		health:    100,
		world: map[string]string{
			"forest":     "你身處一片茂密的森林，陽光透過樹葉灑下。前方有一條小路，東邊似乎有什麼東西。",
			"path":       "你沿著小路前行，周圍的樹木更加稀疏。西邊是森林，東邊是一個破舊的小屋。",
			"hut":        "你來到一個破舊的小屋前，門半掩著。西邊是小路。",
			"hut_inside": "小屋內部昏暗而簡陋。角落裡有一張桌子和一個箱子。",
			"table":      "桌子上放著一把生鏽的鑰匙。",
			"locked_box": "一個上了鎖的箱子。也許附近有什麼東西可以打開它。",
			"open_box":   "箱子是空的。",
		},
		items: map[string]string{
			"key": "一把看起來很舊的鐵鑰匙。",
		},
		err:             nil,
		waitingForInput: "",
	}
}

// Initialize Bubble Tea program
func (m model) Init() tea.Cmd {
	return nil
}

// Update game state
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "l":
			return m.look(), nil
		case "go":
			// Wait for next input as direction
			m.waitingForInput = "go"
			return m, nil
		case "take":
			// Wait for next input as item name
			m.waitingForInput = "take"
			return m, nil
		case "use":
			// Wait for next input as item to use
			m.waitingForInput = "use"
			return m, nil
		case "inventory", "i":
			return m.showInventory(), nil
		case "status", "s":
			return m.showStatus(), nil
		default:
			if m.waitingForInput == "go" {
				return m.moveTo(msg.String()), nil
			} else if m.waitingForInput == "take" {
				return m.takeItem(msg.String()), nil
			} else if m.waitingForInput == "use" {
				return m.useItem(msg.String()), nil
			} else {
				m.err = fmt.Errorf("我不明白 '%s'。你可以嘗試 'look', 'go [方向]', 'take [物品]', 'use [物品]', 'inventory', 'status'。", msg.String())
				return m, nil
			}
		}
	}
	return m, nil
}

// Show current location description
func (m model) look() model {
	m.err = nil
	fmt.Println(m.world[m.location])
	if m.location == "hut_inside" && !m.hasItem("key") {
		fmt.Println("你看到角落的桌子上有一把 [鑰匙]。")
	}
	if m.location == "hut_inside" && m.hasItem("key") && m.location != "open_box" {
		fmt.Println("你看到角落裡有一個 [箱子]。")
	}
	if m.location == "locked_box" && m.hasItem("key") {
		fmt.Println("這個 [箱子] 看起來可以用你找到的鑰匙打開。")
	}
	return m
}

// Move to new location
func (m model) moveTo(direction string) model {
	m.err = nil
	m.waitingForInput = ""
	switch m.location {
	case "forest":
		if direction == "path" || direction == "前方" {
			m.location = "path"
			fmt.Println("你沿著小路前行。")
		} else if direction == "east" || direction == "東邊" {
			m.location = "hut"
			fmt.Println("你朝東邊走去，看到一個破舊的小屋。")
		} else {
			m.err = fmt.Errorf("你無法朝那個方向移動。")
		}
	case "path":
		if direction == "west" || direction == "西邊" {
			m.location = "forest"
			fmt.Println("你回到了森林。")
		} else if direction == "east" || direction == "東邊" {
			m.location = "hut"
			fmt.Println("你繼續沿著小路，來到了小屋前。")
		} else {
			m.err = fmt.Errorf("你無法朝那個方向移動。")
		}
	case "hut":
		if direction == "west" || direction == "西邊" {
			m.location = "path"
			fmt.Println("你離開小屋，回到了小路。")
		} else if direction == "enter" || direction == "進入" {
			m.location = "hut_inside"
			fmt.Println("你走進了小屋。裡面很昏暗。")
		} else {
			m.err = fmt.Errorf("你無法朝那個方向移動。")
		}
	case "hut_inside":
		if direction == "exit" || direction == "離開" {
			m.location = "hut"
			fmt.Println("你走出了小屋。")
		} else if (direction == "table" || direction == "桌子") && !m.hasItem("key") {
			fmt.Println(m.world["table"])
		} else if direction == "box" || direction == "箱子" {
			if m.hasItem("key") {
				m.location = "locked_box"
				fmt.Println(m.world["locked_box"])
			} else {
				fmt.Println(m.world["locked_box"])
			}
		} else {
			m.err = fmt.Errorf("你無法在那裡移動。")
		}
	case "locked_box":
		if direction == "hut" || direction == "小屋" || direction == "離開" {
			m.location = "hut_inside"
			fmt.Println("你回到了小屋內部。")
		}
	case "open_box":
		if direction == "hut" || direction == "小屋" || direction == "離開" {
			m.location = "hut_inside"
			fmt.Println("你回到了小屋內部。")
		}
	default:
		m.err = fmt.Errorf("你不知道該往哪裡去。")
	}
	return m
}

// Pick up item
func (m model) takeItem(itemName string) model {
	m.err = nil
	m.waitingForInput = ""
	if m.location == "hut_inside" && itemName == "key" && !m.hasItem("key") {
		m.inventory = append(m.inventory, "key")
		fmt.Println("你撿起了鑰匙。")
	} else {
		m.err = fmt.Errorf("你在這裡找不到 [%s]。", itemName)
	}
	return m
}

// Use item
func (m model) useItem(itemName string) model {
	m.err = nil
	m.waitingForInput = ""
	if m.location == "locked_box" && itemName == "key" && m.hasItem("key") {
		m.location = "open_box"
		fmt.Println("你用鑰匙打開了箱子。裡面是空的。")
	} else if !m.hasItem(itemName) {
		m.err = fmt.Errorf("你沒有 [%s] 這個物品。", itemName)
	} else {
		m.err = fmt.Errorf("你不知道如何在這裡使用 [%s]。", itemName)
	}
	return m
}

// Show inventory
func (m model) showInventory() model {
	m.err = nil
	if len(m.inventory) == 0 {
		fmt.Println("你的背包是空的。")
	} else {
		fmt.Println("你的背包裡有:")
		for _, item := range m.inventory {
			fmt.Printf("- %s (%s)\n", item, m.items[item])
		}
	}
	return m
}

// Show character status
func (m model) showStatus() model {
	m.err = nil
	fmt.Printf("你的生命值: %d\n", m.health)
	fmt.Printf("你目前在: %s\n", m.location)
	return m
}

// Check if player has an item
func (m model) hasItem(itemName string) bool {
	for _, item := range m.inventory {
		if item == itemName {
			return true
		}
	}
	return false
}

// View (UI)
func (m model) View() string {
	s := "歡迎來到文字冒險！\n\n"
	s += m.world[m.location] + "\n\n"

	if m.location == "hut_inside" && !m.hasItem("key") {
		s += "你看到角落的桌子上有一把 [鑰匙]。\n"
	}
	if m.location == "hut_inside" && m.hasItem("key") && m.location != "open_box" {
		s += "你看到角落裡有一個 [箱子]。\n"
	}
	if m.location == "locked_box" && m.hasItem("key") {
		s += "這個 [箱子] 看起來可以用你找到的鑰匙打開。\n"
	}

	s += "\n你可以做什麼？ (輸入命令)\n"
	s += "- look (l)\n"
	s += "- go [方向]\n"
	s += "- take [物品]\n"
	s += "- use [物品]\n"
	s += "- inventory (i)\n"
	s += "- status (s)\n"
	s += "- quit (q)\n\n"

	if m.err != nil {
		s += fmt.Sprintf("錯誤: %v\n", m.err)
	}

	if m.waitingForInput != "" {
		switch m.waitingForInput {
		case "go":
			s += "請輸入方向: "
		case "take":
			s += "請輸入要撿起的物品: "
		case "use":
			s += "請輸入要使用的物品: "
		}
	}

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
