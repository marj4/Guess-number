package main

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"strconv"
)

const (
	score      = "Score:"
	attempLab1 = "Attemp:"
)

var attemp, scor int

func main() {
	a := app.New()
	b := a.NewWindow("                             Guess the number ")
	b.CenterOnScreen()
	b.Resize(fyne.Size{400, 550})
	b.SetFixedSize(true)

	sc := widget.NewLabel(score + "0")
	en := widget.NewEntry()
	res := widget.NewLabel("")
	attempLab2 := widget.NewLabel(attempLab1 + "0")
	g_num := widget.NewLabel("")
	num := rand.Intn(6)
	//cont := container.NewHBox(sc)

	but1 := widget.NewButton("Start", func() {
		attemp++
		attempLab2.SetText(attempLab1 + strconv.Itoa(attemp))
		nump := en.Text
		str, err := strconv.Atoi(nump)
		if err != nil || str < 0 || str > 6 {
			dialog.ShowError(
				errors.New("Please,enter the number [0:6]"), b)
			return
		}
		if str == num {
			res.SetText("Success!!!")
			scor++
			sc.SetText(score + strconv.Itoa(scor))
			attemp = 0
			num2 := rand.Intn(6)
			num = num2
			attempLab2.SetText(attempLab1 + strconv.Itoa(attemp))

		} else {
			res.SetText("No")
		}
	})
	but2 := widget.NewButton("Restart", func() {
		attemp = 0
		num2 := rand.Intn(6)
		num = num2
		attempLab2.SetText(attempLab1 + strconv.Itoa(attemp))
	})

	b.SetMainMenu(menu(a))
	b.SetContent(container.NewVBox(sc, en, but1, but2, attempLab2, res, g_num))
	b.ShowAndRun()
}

func menu(window fyne.App) *fyne.MainMenu {
	item1 := fyne.NewMenuItem("Change theme", func() {
		NewWindow2(window)
	})
	menu2 := fyne.NewMenu("Settings", item1)
	return fyne.NewMainMenu(menu2)
}

func NewWindow2(a fyne.App) {
	b := a.NewWindow("Settings")
	b.Resize(fyne.NewSize(200, 110))
	b.CenterOnScreen()
	b.SetFixedSize(true)

	c := widget.NewRadioGroup([]string{"Set a light theme", "Set a black theme"}, func(s string) {

	})
	but1 := widget.NewButton("Ok", func() {
		if c.Selected == "Set a light theme" {
			a.Settings().SetTheme(theme.LightTheme())
			b.Close()
		} else {
			a.Settings().SetTheme(theme.DarkTheme())
			b.Close()
		}
	})
	b.SetContent(container.NewVBox(c, but1))
	b.Show()
}
