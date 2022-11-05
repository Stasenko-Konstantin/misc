package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/handlers"
	"github.com/hegedustibor/htgo-tts/voices"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	wait = 10
)

var (
	parts      = []string{"левая рука", "левая нога", "правая рука", "правая нога"}
	colors     = []string{"красный", "желтый", "синий", "зеленый"}
	players    int
	currPlayer = 1
)

func checkErr(err error, w *fyne.Window) {
	if err != nil {
		dialog.ShowError(err, *w)
		log.Println(err.Error())
	}
}

func strPlayer(i int) string {
	switch i {
	case 1:
		return "один"
	case 2:
		return "два"
	case 3:
		return "три"
	case 4:
		return "четыре"
	}
	return ""
}

func main() {
	a := app.New()
	w := a.NewWindow("twister")
	rand.Seed(time.Now().Unix())

	choosePlayers := widget.NewRadioGroup([]string{"2", "3", "4"}, func(p string) {
		r, _ := strconv.Atoi(p)
		players = r
	})
	choosePlayers.SetSelected("2")

	var (
		newGame   *widget.Button
		stopGame  *widget.Button
		startGame *widget.Button
		speech    = htgotts.Speech{Folder: "/storage/emulated/0/Android/data/twister.mobile/files", Language: voices.Russian, Handler: &handlers.Native{}}
	)

	newGame = widget.NewButton("Новая игра", func() {
		dialog.ShowCustom("Выберите кол-во игроков:", "ok", choosePlayers, w)
		newGame.Hide()
		startGame.Show()
	})

	startGame = widget.NewButton("Начать игру", func() {
		startGame.Hide()
		stopGame.Show()
		part := parts[rand.Intn(len(parts))]
		color := colors[rand.Intn(len(colors))]
		text := "игрок номер " + strPlayer(currPlayer) + ", " + part + " на " + color
		for i := 0; ; i++ {
			//err := speech.Speak(text)
			f, err := speech.CreateSpeechFile(text, strconv.Itoa(i)+".mp3")
			checkErr(err, &w)
			err = speech.PlaySpeechFile(f)
			checkErr(err, &w)

			if currPlayer == players {
				currPlayer = 1
			} else {
				currPlayer++
			}
			//time.Sleep(wait * time.Second)
		}
	})
	startGame.Hide()

	stopGame = widget.NewButton("Остановить игру", func() {
		stopGame.Hide()
		newGame.Show()
	})
	stopGame.Hide()

	content := container.NewVBox(
		widget.NewLabel("\n\n\n\n"),
		newGame,
		startGame,
		stopGame,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
