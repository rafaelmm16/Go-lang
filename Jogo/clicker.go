package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/rivo/tview"
)

func main() {
	// Crie uma aplicação tview
	app := tview.NewApplication()

	// Variáveis para controlar os pontos e atualizações de pontos
	var (
		points     int
		pointsLock sync.Mutex
	)

	// componente Textview para mostrar os pontos
	pointsTextView := tview.NewTextView().
		SetText("Pontos: 0").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Função para atualizar a exibição de pontos
	updatePoints := func() {
		pointsLock.Lock()
		pointsTextView.SetText("Pontos: " + strconv.Itoa(points))
		pointsLock.Unlock()
	}

	//botão para clicar e ganhar pontos
	clickButton := tview.NewButton("Clique para ganhar pontos").
		SetSelectedFunc(func() {
			pointsLock.Lock()
			points++
			pointsLock.Unlock()
			updatePoints()
		})

	//função de atualização para ganhar pontos automaticamente
	go func() {
		for {
			time.Sleep(1 * time.Second)
			pointsLock.Lock()
			points++
			pointsLock.Unlock()
			app.QueueUpdateDraw(func() {
				updatePoints()
			})
		}
	}()

	// Layout da aplicação
	flex := tview.NewFlex().
		AddItem(pointsTextView, 0, 1, false).
		AddItem(clickButton, 0, 1, true)

	// Configure o aplicativo e inicie
	if err := app.SetRoot(flex, true).Run(); err != nil {
		fmt.Println(err)
	}
}
