package main

import (
	"fmt"
	"gogogadget/pkg/tools"
	"strconv"

	"github.com/rivo/tview"
)

var app *tview.Application
var menu *tview.List

func main() {
	app = tview.NewApplication()

	menu = tview.NewList().AddItem("Prettify JSON", "Prettify a JSON string", '1', func() {
		app.SetFocus(nil) // Remove focus from the menu to avoid input interference
		prettifyJSON()
	}).
		AddItem("Password", "Generate a password", '2', func() {
			app.SetFocus(nil) // Remove focus from the menu to avoid input interference
			generatePassword()
		}).
		AddItem("Exit", "Exit the application", 'q', func() {
			app.Stop()
		})

	menu.SetTitle("Menu").SetTitleAlign(tview.AlignCenter).
		SetBorder(true).SetBorderPadding(1, 1, 1, 1)

	if err := app.SetRoot(menu, true).Run(); err != nil {
		panic(err)
	}
}

func prettifyJSON() {
	form := tview.NewForm()
	form.SetTitle("Prettify JSON").SetTitleAlign(tview.AlignCenter).
		SetBorder(true).SetBorderPadding(1, 1, 1, 1)

	inputField := tview.NewInputField().
		SetLabel("Enter JSON: ").
		SetFieldWidth(0)

	form.AddFormItem(inputField)

	form.AddButton("Prettify", func() {
		prettyJSON, err := tools.IndentJSON(inputField.GetText(), 4)
		if err != nil {
			displayResult(fmt.Sprintf("Invalid JSON: %v", err))
		} else {
			displayResult(string(prettyJSON))
		}
	})

	form.AddButton("Back", func() {
		app.SetRoot(menu, true)
	})

	app.SetRoot(form, true)
}

func generatePassword() {
	form := tview.NewForm()
	form.SetTitle("Generate Password").SetTitleAlign(tview.AlignCenter).
		SetBorder(true).SetBorderPadding(1, 1, 1, 1)

	lengthInput := tview.NewInputField().
		SetLabel("Password Length: ").
		SetText("16") // Default length

	digitInput := tview.NewInputField().
		SetLabel("Number of Digits: ").
		SetText("4") // Default number of digits

	symbolInput := tview.NewInputField().
		SetLabel("Number of Symbols: ").
		SetText("2") // Default number of symbols$

	quantityInput := tview.NewInputField().
		SetLabel("Number of passwords: ").
		SetText("32") // Default number of symbols

	form.AddFormItem(lengthInput)
	form.AddFormItem(digitInput)
	form.AddFormItem(symbolInput)
	form.AddFormItem(quantityInput)

	form.AddButton("Generate", func() {
		length, _ := strconv.Atoi(lengthInput.GetText())
		numDigits, _ := strconv.Atoi(digitInput.GetText())
		numSymbols, _ := strconv.Atoi(symbolInput.GetText())
		numQuantity, _ := strconv.Atoi(quantityInput.GetText())

		options := tools.PasswordOptions{
			Length:     length,
			NumDigits:  numDigits,
			NumSymbols: numSymbols,
		}
		var passwords []string
		for i := 0; i < numQuantity; i++ {

			password, err := tools.GeneratePassword(options)
			if err != nil {
				displayResult(fmt.Sprintf("Unable to generate passwords n°%d : %v", i, err))
				return
			}
			passwords = append(passwords, password)
		}

		displayPasswordResults(passwords[:])
	})

	form.AddButton("Back", func() {
		app.SetRoot(menu, true)
	})

	app.SetRoot(form, true)
}

func displayResult(result string) {
	resultView := tview.NewTextView().SetText(result).
		SetTextAlign(tview.AlignLeft)

	backButton := tview.NewButton("Back").SetSelectedFunc(func() {
		app.SetRoot(menu, true)
	})

	flex := tview.NewFlex().
		AddItem(resultView, 0, 1, true).
		AddItem(backButton, 1, 0, false)

	app.SetRoot(flex, true)
}

func displayPasswordResults(passwords []string) {
	resultView := tview.NewTextView().
		SetTextAlign(tview.AlignLeft)

	for _, password := range passwords {
		resultView.SetText(resultView.GetText(true) + password + "\n")
	}

	backButton := tview.NewButton("Back").SetSelectedFunc(func() {
		app.SetRoot(menu, true)
	})

	flex := tview.NewFlex().
		AddItem(resultView, 0, 1, true).
		AddItem(backButton, 1, 1, false)

	app.SetRoot(flex, true)
}
