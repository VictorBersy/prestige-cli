package generators

import (
	"github.com/VictorBersy/prestige-cli/internal/ui/context"
	"github.com/VictorBersy/prestige-cli/internal/ui/layers"
	tea "github.com/charmbracelet/bubbletea"
)

type Generators struct {
	layer *layers.Model
}

func NewModel(id int, ctx *context.ProgramContext) Generators {
	g := Generators{
		layer: &layers.Model{
			Id:   id,
			Tier: 2,
			Ctx:  ctx,
			Name: "Generators",
		},
	}

	return g
}

func (g *Generators) Id() int {
	return g.layer.Id
}

func (g *Generators) Name() string {
	return g.layer.Name
}

func (g *Generators) Tier() int {
	return g.layer.Tier
}

func (g *Generators) UpdateProgramContext(ctx *context.ProgramContext) {
	g.layer.UpdateProgramContext(ctx)
}

func (g *Generators) Tick() {
}

func (g *Generators) Prestige() {
}

func (g *Generators) NextPrestigeAt() float64 {
	return 10
}

func (g Generators) Update(msg tea.Msg) (layers.Layer, tea.Cmd) {
	var cmd tea.Cmd
	return &g, cmd
}

func Fetch(id int, ctx context.ProgramContext) (layer layers.Layer) {
	layerModel := NewModel(id, &ctx)
	return &layerModel
}