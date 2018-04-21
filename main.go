package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

const (
	XMAX = 10
	YMAX = 10
)

var app = &views.Application{}
var window = &mainWindow{}
var version = "dev"

type model struct {
	endx       int
	endy       int
	x          int
	y          int
	enab       bool
	loc        string
	generation int
	world      world
	stable     bool
	screen     tcell.Screen
}

func (m *model) GetBounds() (int, int) {
	return m.endx, m.endy
}

func (m *model) MoveCursor(offx, offy int) {
	m.x += offx
	m.y += offy
	m.limitCursor()
}

func (m *model) limitCursor() {
	if m.x < 0 {
		m.x = 0
	}
	if m.x > m.endx-1 {
		m.x = m.endx - 1
	}
	if m.y < 0 {
		m.y = 0
	}
	if m.y > m.endy-1 {
		m.y = m.endy - 1
	}
	m.loc = fmt.Sprintf("Cursor is %d,%d", m.x, m.y)
}

func (m *model) GetCursor() (int, int, bool, bool) {
	return m.x, m.y, m.enab, true
}

func (m *model) SetCursor(x int, y int) {
	m.x = x
	m.y = y

	m.limitCursor()
}

func (m *model) GetCell(x, y int) (rune, tcell.Style, []rune, int) {
	var ch rune
	style := tcell.StyleDefault
	if x >= m.endx || y >= m.endy {
		return ch, style, nil, 1
	}

	if m.world.latest[x][y].IsAlive() {
		ch = '*'
	} else {
		ch = ' '
	}

	return ch, style, nil, 1
}

type mainWindow struct {
	main   *views.CellView
	keybar *views.SimpleStyledText
	status *views.SimpleStyledTextBar
	model  *model

	views.Panel
}

func (a *mainWindow) HandleEvent(ev tcell.Event) bool {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyCtrlL:
			app.Refresh()
			return true
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'Q', 'q':
				app.Quit()
				return true
			case 'E', 'e':
				a.model.enab = true
				a.updateKeys()
				return true
			case 'D', 'd':
				a.model.enab = false
				a.updateKeys()
				return true
			case 'N', 'n':
				if a.model.stable {
					return true
				}

				a.model.world.Evaluate()
				a.model.world.NextGeneration()
				a.model.generation++
				a.main.Draw()

				a.status.SetRight(fmt.Sprintf("Generation: %d", a.model.generation))
				a.Panel.Draw()

				if a.model.world.IsStable() {
					a.status.SetCenter("STABLE")
					a.model.stable = true
					a.updateKeys()
				}
				return true
			case ' ':
				x, y, _, _ := a.model.GetCursor()

				if a.model.world.latest[x][y].IsAlive() {
					a.model.world.latest[x][y].Die()
				} else {
					a.model.world.latest[x][y].Revival()
				}
				a.main.Draw()
				return true
			case 'R', 'r':
				a.model.generation = 0
				a.model.stable = false
				a.updateKeys()
				a.status.SetCenter("")
			}
		}
	}
	return a.Panel.HandleEvent(ev)
}

func (a *mainWindow) Draw() {
	a.status.SetLeft(a.model.loc)
	a.Panel.Draw()
}

func (a *mainWindow) updateKeys() {
	m := a.model
	w := "[%AQ%N] Quit"
	w += "  [%ASPACE%N] Revival/Die cell"
	if !m.stable {
		w += "  [%AN%N] Next generation"
	} else {
		w += "  [%AR%N] Restart"
	}
	if !m.enab {
		w += "  [%AE%N] Enable cursor"
	} else {
		w += "  [%AD%N] Disable cursor"
	}
	a.keybar.SetMarkup(w)
	app.Update()
}

func (a *mainWindow) Resize() {
	endx, endy := a.model.screen.Size()

	a.status.SetCenter(fmt.Sprintf("resize x:%d, y:%d", endx, endy))

	a.model.endx = endx
	a.model.endy = endy - 3
	a.model.world = New(endx, endy)
}

func main() {

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	window.model = &model{enab: true, screen: s}

	title := views.NewTextBar()
	title.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorTeal).
		Foreground(tcell.ColorWhite))
	title.SetCenter("Game of life S23/B3", tcell.StyleDefault)
	title.SetRight(version, tcell.StyleDefault)

	window.keybar = views.NewSimpleStyledText()
	window.keybar.RegisterStyle('N', tcell.StyleDefault.
		Background(tcell.ColorSilver).
		Foreground(tcell.ColorBlack))
	window.keybar.RegisterStyle('A', tcell.StyleDefault.
		Background(tcell.ColorSilver).
		Foreground(tcell.ColorRed))

	window.status = views.NewSimpleStyledTextBar()
	window.status.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlue).
		Foreground(tcell.ColorYellow))
	window.status.RegisterLeftStyle('N', tcell.StyleDefault.
		Background(tcell.ColorYellow).
		Foreground(tcell.ColorBlack))

	window.status.SetLeft("My status is here.")
	window.status.SetRight("%UGeneration%N: 0 ")

	window.main = views.NewCellView()
	window.main.SetModel(window.model)
	window.main.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack))

	window.SetMenu(window.keybar)
	window.SetTitle(title)
	window.SetContent(window.main)
	window.SetStatus(window.status)

	window.updateKeys()

	app.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	app.SetRootWidget(window)
	app.SetScreen(s)
	if e := app.Run(); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(1)
	}
}
