package ui

import (
	"fmt"

	helpdesk "freeipa-tui/internal/helpdesk/backend"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Run(client *helpdesk.IPAClient) error {
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
	searchInput.SetFieldWidth(40)
	searchInput.SetLabelColor(tcell.ColorYellow)
	searchInput.SetFieldBackgroundColor(tcell.ColorBlack)
	searchInput.SetFieldTextColor(tcell.ColorWhite)

	searchHint := tview.NewTextView()
	searchHint.SetText("Введите логин, имя или фамилию и нажмите Enter. Esc — назад.")

	searchResults := tview.NewList()
	searchResults.SetBorder(true)
	searchResults.SetTitle(" Результат ")
	searchResults.ShowSecondaryText(false)
	searchResults.AddItem("Введите запрос и нажмите Enter", "", 0, nil)

	userCard := tview.NewTextView()
	userCard.SetBorder(true)
	userCard.SetTitle(" Карточка пользователя ")

	showUserCard := func(user helpdesk.IPAUser) {
		userCard.SetText(fmt.Sprintf(
			"Login: %s\nИмя: %s\nФамилия: %s\n\nДействия появятся тут позже.\n\nEsc — назад к результатам поиска",
			user.Username,
			user.FirstName,
			user.LastName,
		))

		pages.SwitchToPage("user_card")
		app.SetFocus(userCard)
	}

	searchInput.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}

		query := searchInput.GetText()
		if query == "" {
			searchResults.Clear()
			searchResults.AddItem("Введите логин, имя или фамилию пользователя.", "", 0, nil)
			return
		}

		searchResults.Clear()
		searchResults.AddItem("Ищу пользователей...", "", 0, nil)

		users, err := client.FindUsers(query)
		if err != nil {
			searchResults.Clear()
			searchResults.AddItem("Ошибка поиска пользователя: "+err.Error(), "", 0, nil)
			return
		}

		if len(users) == 0 {
			searchResults.Clear()
			searchResults.AddItem("Пользователи не найдены.", "", 0, nil)
			return
		}

		searchResults.Clear()
		searchResults.SetTitle(fmt.Sprintf(" Результат: %d ", len(users)))
		for _, user := range users {
			user := user
			title := fmt.Sprintf("%s — %s %s", user.Username, user.FirstName, user.LastName)
			searchResults.AddItem(title, "", 0, func() {
				showUserCard(user)
			})
		}

		app.SetFocus(searchResults)
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
	searchPanel.SetDirection(tview.FlexRow)
	searchPanel.AddItem(searchInput, 1, 0, true)
	searchPanel.AddItem(searchHint, 1, 0, false)

	userFindScreen := tview.NewFlex()
	userFindScreen.SetDirection(tview.FlexRow)
	userFindScreen.AddItem(searchPanel, 4, 0, true)
	userFindScreen.AddItem(searchResults, 0, 1, false)

	pages.AddPage("main", layout, true, true)
	pages.AddPage("user_find", userFindScreen, true, false)
	pages.AddPage("user_card", userCard, true, false)

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

			if pageName == "user_card" {
				pages.SwitchToPage("user_find")
				app.SetFocus(searchResults)
				return nil
			}

			pages.SwitchToPage("main")
			app.SetFocus(menu)
			return nil
		}

		return event
	})

	return app.SetRoot(pages, true).Run()
}
