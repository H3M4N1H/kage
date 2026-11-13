package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"strings"
)

// User-supplied block-letter wordmark; shown only when it fits the menu.
const nameBanner = `██╗  ██╗██████╗ ███╗   ███╗██╗  ██╗███╗   ██╗ ██╗██╗  ██╗
██║  ██║╚════██╗████╗ ████║██║  ██║████╗  ██║███║██║  ██║
███████║ █████╔╝██╔████╔██║███████║██╔██╗ ██║╚██║███████║
██╔══██║ ╚═══██╗██║╚██╔╝██║╚════██║██║╚██╗██║ ██║██╔══██║
██║  ██║██████╔╝██║ ╚═╝ ██║     ██║██║ ╚████║ ██║██║  ██║
╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═╝╚═╝  ╚═╝`

type model struct {
	content                             Content
	cursor, page, offset, width, height int
	accent                              lipgloss.Style
}

func newModel(c Content, style lipgloss.Style) model {
	return model{content: c, page: -1, width: 80, height: 24, accent: style}
}
func (m model) Init() tea.Cmd { return nil }
func clamp(n, lo, hi int) int { return max(lo, min(n, hi)) }
func wrap(s string, width int) string {
	return ansi.Hardwrap(ansi.Wrap(s, max(1, width), ""), max(1, width), false)
}
func (m model) pageLines() []string {
	if m.page < 0 {
		return nil
	}
	return strings.Split(wrap(m.content.Sections[m.page].Body, max(1, m.width-4)), "\n")
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch v := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = clamp(v.Width, 1, 160)
		m.height = clamp(v.Height, 1, 60)
	case tea.KeyMsg:
		switch v.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc", "backspace":
			m.page = -1
			m.offset = 0
		case "up", "k":
			if m.page < 0 {
				m.cursor = (m.cursor + 7) % 8
			} else {
				m.offset--
			}
		case "down", "j":
			if m.page < 0 {
				m.cursor = (m.cursor + 1) % 8
			} else {
				m.offset++
			}
		case "pgdown":
			m.offset += max(1, m.height-8)
		case "pgup":
			m.offset -= max(1, m.height-8)
		case "enter":
			if m.page < 0 {
				m.page = m.cursor
				m.offset = 0
			}
		default:
			if len(v.String()) == 1 && v.String()[0] >= '1' && v.String()[0] <= '8' {
				m.page = int(v.String()[0] - '1')
				m.cursor = m.page
				m.offset = 0
			}
		}
	}
	m.offset = clamp(m.offset, 0, max(0, len(m.pageLines())-max(1, m.height-8)))
	return m, nil
}
func (m model) View() string {
	if m.width < 24 || m.height < 10 {
		return ansi.Truncate("Resize to 24x10 or larger. q exits.", m.width, "")
	}
	title := "h3m4n1h / HEMANTH / Security × Software"
	if m.width < 34 {
		title = "h3m4n1h / HEMANTH"
	}
	lines := []string{}
	if m.page < 0 && m.width >= lipgloss.Width(nameBanner)+2 && m.height >= 23 {
		bannerStyle := m.accent.Foreground(lipgloss.Color("#FF5F56")).Bold(true)
		for _, row := range strings.Split(nameBanner, "\n") {
			lines = append(lines, bannerStyle.Render(row))
		}
	}
	lines = append(lines, m.accent.Render(ansi.Truncate(title, m.width-2, "")), strings.Repeat("─", m.width-2), "")
	available := max(1, m.height-8)
	if m.page < 0 {
		available = max(1, m.height-len(lines)-3)
		start := clamp(m.cursor-available+1, 0, max(0, 8-available))
		for i := start; i < min(8, start+available); i++ {
			prefix := "  "
			if i == m.cursor {
				prefix = "> "
			}
			row := ansi.Truncate(fmt.Sprintf("%s%d  %s", prefix, i+1, m.content.Sections[i].Title), m.width-2, "…")
			if i == m.cursor {
				row = m.accent.Render(row)
			}
			lines = append(lines, row)
		}
	} else {
		lines = append(lines, m.accent.Render(ansi.Truncate(m.content.Sections[m.page].Title, m.width-2, "…")))
		body := m.pageLines()
		end := min(len(body), m.offset+available)
		lines = append(lines, body[m.offset:end]...)
	}
	lines = append(lines, "", ansi.Truncate("1–8 select · ↑↓ move · Esc back · q exit", m.width-2, ""))
	return strings.Join(lines, "\n") + "\n"
}
func plainContent(c Content) string {
	var b strings.Builder
	b.WriteString(nameBanner + "\nh3m4n1h / " + c.Name + " — Read-only portfolio\n\n")
	for i, s := range c.Sections {
		fmt.Fprintf(&b, "%d. %s\n%s\n\n", i+1, s.Title, s.Body)
	}
	return b.String()
}
