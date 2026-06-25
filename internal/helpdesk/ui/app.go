package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run() error {
	app := tview.NewApplication()
	pages := tview.NewPages()

	footer := tview.NewTextView()
	footer.SetText(" ↑↓ выбор | Enter открыть | Esc/Ctrl+Q выход ")
	footer.SetTextAlign(tview.AlignCenter)

	menu := tview.NewList()
	menu.SetBorder(true)
	menu.SetTitle(" FreeIPA Helpdesk ")
	menu.ShowSecondaryText(false)

	description := tview.NewTextView()
	description.SetBorder(true)
	description.SetTitle(" Описание ")
	description.SetText("Выберите действие слева")

	searchInput := tview.NewInputField()
	searchInput.SetLabel("Поиск: ")
	searchInput.SetPlaceholder("login, имя или фамилия")

	searchResults := tview.NewTextView()
	searchResults.SetBorder(true)
	searchResults.SetTitle(" Результат ")
	searchResults.SetText("Введите запрос и нажмите Enter")

	searchInput.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}

		query := searchInput.GetText()
		if query == "" {
			searchResults.SetText("Введите логин, имя или фамилию пользователя.")
			return
		}

		searchResults.SetText("Пока это заглушка.\n\nЗапрос: " + query)
	})

	menu.AddItem("Найти пользователя", "", '1', func() {
		pages.SwitchToPage("user_find")
		app.SetFocus(searchInput)
	})
	menu.AddItem("Создать пользователя", "", '2', nil)
	menu.AddItem("Сбросить пароль", "", '3', nil)
	menu.AddItem("Разблокировать пользователя", "", '4', nil)
	menu.AddItem("Включить / отключить пользователя", "", '5', nil)
	menu.AddItem("Массовые операции", "", '6', nil)

	menu.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		switch index {
		case 0:
			description.SetText("Поиск пользователя по логину, имени или фамилии.")
		case 1:
			description.SetText("Создание нового пользователя FreeIPA.")
		case 2:
			description.SetText("Сброс пароля пользователя и выдача временного пароля.")
		case 3:
			description.SetText("Разблокировка пользователя после блокировки или ошибок входа.")
		case 4:
			description.SetText("Включение или отключение учетной записи пользователя.")
		case 5:
			description.SetText("Массовые операции над списком пользователей.")
		}
	})

	content := tview.NewFlex()
	content.AddItem(menu, 35, 0, true)
	content.AddItem(description, 0, 1, false)

	layout := tview.NewFlex()
	layout.SetDirection(tview.FlexRow)
	layout.AddItem(content, 0, 1, true)
	layout.AddItem(footer, 1, 0, false)

	searchPanel := tview.NewFlex()
	searchPanel.SetBorder(true)
	searchPanel.SetTitle(" Поиск пользователя ")
	searchPanel.AddItem(searchInput, 0, 1, true)

	userFindScreen := tview.NewFlex()
	userFindScreen.SetDirection(tview.FlexRow)
	userFindScreen.AddItem(searchPanel, 3, 0, true)
	userFindScreen.AddItem(searchResults, 0, 1, false)

	pages.AddPage("main", layout, true, true)
	pages.AddPage("user_find", userFindScreen, true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlQ:
			app.Stop()
			return nil
		case tcell.KeyEsc:
			pageName, _ := pages.GetFrontPage()
			if pageName == "main" {
				app.Stop()
				return nil
			}

			pages.SwitchToPage("main")
			return nil
		}

		return event
	})

	return app.SetRoot(pages, true).Run()
}
