package state

import (
	"game/assets"
	"game/ui"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Start struct {
	widgets []ui.Widget
	set     SetState
}

func NewStart(set SetState) *Start {
	s := &Start{set: set}
	textStart := "Старт"
	s.widgets = []ui.Widget{
		ui.NewButton(
			0, 0, 1, 1,
			assets.Images["background"],
			assets.Images["background"],
			func() {
				s.set(NewScore(s.set))
			},
		),
		ui.NewLabel(0.5, 0.35, 0.1, colornames.Red, &textStart),
	}
	return s
}

func (g *Start) Update(screen image.Rectangle) error {
	for _, widget := range g.widgets {
		widget.Update(screen)
	}
	return nil
}

func (g *Start) Draw(screen *ebiten.Image) {
	for _, widget := range g.widgets {
		widget.Draw(screen)
	}
}
