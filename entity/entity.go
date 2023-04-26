package entity

import "fmt"

type EegDto struct {
	Id     int
	Input  EegHistoryModel
	System SystemInfo
}

type EegHistoryModel struct {
	Attention     int
	Meditation    int
	Signal        int
	Delta         int
	Theta         int
	LowAlpha      int
	HighAlpha     int
	LowBeta       int
	HighBeta      int
	LowGamma      int
	HighGamma     int
	SystemMouseId *int64
}

type SystemInfo struct {
	Id   int64
	X    int
	Y    int
	ToX  int
	ToY  int
	EndX int
	EndY int
}

func (b EegHistoryModel) ToString() string {
	return fmt.Sprintf("Attention: %d Meditation: %d Signal: %d Delta: %d Theta: %d LowAlpha: %d HighAlpha: %d LowBeta: %d HighBeta: %d LowGamma: %d HighGamma: %d", b.Attention, b.Meditation, b.Signal, b.Delta, b.Theta, b.LowAlpha, b.HighAlpha, b.LowBeta, b.HighBeta, b.LowGamma, b.HighGamma)
}
