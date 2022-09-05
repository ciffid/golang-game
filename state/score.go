package state

import (
	"fmt"
	"game/assets"
	"game/ui"
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Score struct {
	widgets []ui.Widget
	set     SetState
	ticker  *time.Ticker
	time    int
	score   int
	text    string
}

func NewScore(set SetState) *Score {
	s := &Score{set: set, ticker: time.NewTicker(1 * time.Second)}
	s.time = 5
	s.widgets = []ui.Widget{
		ui.NewBackground(0, 0, 1, 1, assets.Images["background"]),
		ui.NewButton(
			0, 0.75, 1, 0.25,
			assets.Images["button-enabled"],
			assets.Images["button-disabled"],
			func() {
				s.score++
			},
		),
		ui.NewLabel(0.5, 0.35, 0.1, colornames.Red, &s.text),
	}
	return s
}

func (s *Score) Update(screen image.Rectangle) error {
	for _, widget := range s.widgets {
		widget.Update(screen)
	}

	select {
	case <-s.ticker.C:
		s.time--
	default:
	}
	if s.score == 24 {
		s.set(NewSuccess(s.set))
	}
	if s.time == 0 {
		s.set(NewFail(s.set))
	}

	s.text = fmt.Sprintf("Очки: %d\nВремя: %d", s.score, s.time)
	return nil
}

func (s *Score) Draw(screen *ebiten.Image) {
	for _, widget := range s.widgets {
		widget.Draw(screen)
	}
}
