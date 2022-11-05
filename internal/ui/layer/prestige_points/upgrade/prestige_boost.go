package upgrade

import (
	"fmt"
	"math"

	"github.com/VictorBersy/prestige-cli/internal/ui/layer"
)

type prestigeBoost struct {
	Upgrade *layer.ModelUpgrade
}

func FetchPrestigeBoost(layers *layer.Layers) (upgrade layer.Upgrade) {
	model := prestigeBoost{
		Upgrade: &layer.ModelUpgrade{
			Name:        "Prestige Boost",
			Description: "Prestige Points boost Point generation.",
			Layers:      layers,
			Cost:        1,
		},
	}
	return &model
}

func (m *prestigeBoost) Buy() {
	m.Upgrade.Layers.PrestigePoints.Model().Amount = m.Upgrade.Buy(m.Upgrade.Layers.PrestigePoints.Model().Amount)
}

func (m *prestigeBoost) Tick() {
	m.Upgrade.Layers.Points.Amount = m.Upgrade.Layers.Points.Amount + m.TickAmount()/100
}

func (m *prestigeBoost) Effect() string {
	return fmt.Sprintf("%.2fx", m.TickAmount())
}

func (m *prestigeBoost) Unlocked() bool {
	return m.Upgrade.Layers.PrestigePoints.Model().Upgrades[0].GetModel().Enabled
}

func (m *prestigeBoost) TickAmount() float64 {
	return math.Pow(m.Upgrade.Layers.PrestigePoints.Model().Amount+2, 0.5)
}

func (m *prestigeBoost) GetModel() *layer.ModelUpgrade {
	return m.Upgrade
}