package game

import (
	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/services/calculator"
	"github.com/releaseband/golang-developer-test/services/generator"
)

const rowsCount = 3

func New() (*Slot, error) {
	linesConfig, err := lines.ReadLines()
	if err != nil {
		return nil, err
	}
	payTableConfig, err := paytable.ReadPayTable()
	if err != nil {
		return nil, err
	}
	gameTapes, err := symbols.ReadReels()
	if err != nil {
		return nil, err
	}
	gen := generator.NewSymbols(rowsCount, gameTapes)
	calc := calculator.NewCalculator(linesConfig, payTableConfig)

	roundCost := RoundCost(rowsCount)

	slot := newSlot(gen, calc, roundCost)
	return slot, nil
}
