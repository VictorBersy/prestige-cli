package boosters

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	tea "github.com/charmbracelet/bubbletea"
)

type Boosters struct {
	layer *layers.Model
}

func NewModel(id int, ctx *context.ProgramContext) Boosters {
	b := Boosters{
		layer: &layers.Model{
			Id:   id,
			Tier: 2,
			Ctx:  ctx,
			Name: "Boosters",
		},
	}

	return b
}

func (b *Boosters) Id() int {
	return b.layer.Id
}

func (b *Boosters) Name() string {
	return b.layer.Name
}

func (b *Boosters) Tier() int {
	return b.layer.Tier
}

func (b *Boosters) UpdateProgramContext(ctx *context.ProgramContext) {
	b.layer.UpdateProgramContext(ctx)
}

func (b *Boosters) Tick() {
}

func (b *Boosters) Prestige() {
}

func (b *Boosters) NextPrestigeAt() float64 {
	return 10
}

func (m Boosters) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &m, cmd
}

func (b *Boosters) View() string {
	return "Boosters"
}

func Fetch(id int, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, &ctx)
	return &layerModel
}