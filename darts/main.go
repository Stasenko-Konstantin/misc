package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
	"strconv"
	"strings"
)

var (
	players    int
	maxPoints  int
	currPlayer = 1
	points     []int
)

func winText() string {
	b := strings.Builder{}
	for i, p := range points[1:] {
		b.WriteString(strconv.Itoa(i + 1))
		b.WriteString(": ")
		b.WriteString(strconv.Itoa(p))
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	a := app.New()
	w := a.NewWindow("darts")

	choosePlayers := widget.NewSelect(
		[]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		func(p string) {
			i, _ := strconv.Atoi(p)
			players = i
		})
	choosePlayers.SetSelected("2")

	chooseRegime := widget.NewRadioGroup(
		[]string{"501", "701", "1001"},
		func(p string) {
			r, _ := strconv.Atoi(p)
			maxPoints = r
		})
	chooseRegime.SetSelected("501")

	var (
		newGame   *widget.Button
		startGame *widget.Button
		newPoints = widget.NewEntry()
	)

	newGame = widget.NewButton("Новая игра", func() {
		dialog.ShowCustom("Выберите кол-во игроков:", "ok", choosePlayers, w)
		dialog.ShowCustom("Выберите режим игры:", "ok", chooseRegime, w)
		newGame.Hide()
		startGame.Show()
	})

	startGame = widget.NewButton("Начать игру", func() {
		startGame.Hide()
		var (
			loop      func(bool)
			nextRound = func() {
				text := fmt.Sprintf("игрок %d: %dp", currPlayer, points[currPlayer])
				dialog.ShowCustomConfirm(text, "дальше", "стоп", newPoints, loop, w)
			}
		)
		points = make([]int, players+1)
		for p := 1; p <= players; p++ {
			points[p] = maxPoints
		}
		currPlayer = 1
		loop = func(b bool) {
			defer func() {
				newPoints.SetText("")
			}()
			if !b {
				newGame.Show()
				return
			}
			p := newPoints.Text
			i, err := strconv.Atoi(p)
			if err != nil {
				nextRound()
				return
			}
			np := points[currPlayer] - i
			points[currPlayer] = np
			if np <= 0 {
				dialog.ShowCustomConfirm("победа!", "еще!", "хватит...", widget.NewLabel(winText()), func(b bool) {
					if !b {
						os.Exit(1)
					}
					newGame.Show()
				}, w)
				return
			}
			if currPlayer == players {
				currPlayer = 1
			} else {
				currPlayer++
			}
			nextRound()
		}
		nextRound()
	})
	startGame.Hide()

	content := container.NewVBox(
		widget.NewLabel("\n\n\n\n"),
		newGame,
		startGame,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
