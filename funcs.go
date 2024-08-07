package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func menu(window fyne.App) *fyne.MainMenu {
	item1 := fyne.NewMenuItem("Change the light theme", func() {
		window.Settings().SetTheme(theme.LightTheme())
	})
	item2 := fyne.NewMenuItem("Change the dark theme", func() {
		window.Settings().SetTheme(theme.DarkTheme())
	})
	menu2 := fyne.NewMenu("Settings", item1, item2)
	return fyne.NewMainMenu(menu2)
}
