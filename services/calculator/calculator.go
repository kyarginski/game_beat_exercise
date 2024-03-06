package calculator

import (
	"errors"

	"github.com/releaseband/golang-developer-test/internal/configs/lines"
	"github.com/releaseband/golang-developer-test/internal/configs/paytable"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/game/win"
)

// WILD - специальный символ, который может заменить любой другой символ
// он не имеет своего выигрыша, но может увеличить выигрыш за счет замены другого символа
const WILD = symbols.Symbol(0)

type Calculator struct {
	lines    lines.Lines
	payTable *paytable.PayTable
}

func NewCalculator(lines lines.Lines, payTable *paytable.PayTable) *Calculator {
	return &Calculator{lines: lines, payTable: payTable}
}

func (c *Calculator) Calculate(spinSymbols symbols.Reels) ([]win.Win, error) {
	if len(spinSymbols) == 0 {
		return nil, errors.New("no reels")
	}
	var wins []win.Win

	reelsCount := len(spinSymbols)
	symbolsPerReel := len(spinSymbols[0])

	symbolCounts := make(map[symbols.Symbol]int, len(spinSymbols))

	for symbolIndex := 0; symbolIndex < symbolsPerReel; symbolIndex++ {
		var prevSymbol symbols.Symbol
		for reelIndex := 0; reelIndex < reelsCount; reelIndex++ {
			currentSymbol := spinSymbols[reelIndex][symbolIndex]
			if currentSymbol == prevSymbol || currentSymbol == WILD {
				symbolCounts[currentSymbol]++
			} else {
				if prevSymbol != 0 {
					continue
				}
			}
			prevSymbol = currentSymbol
			// fmt.Printf("Symbol at Reel %d, Position %d: %d\n", reelIndex+1, symbolIndex+1, currentSymbol)
		}
		// fmt.Println("--- Line End ---")
	}

	maxCount := 2 // ограничиваем количество символов для выигрыша
	winSymbols := make([]symbols.Symbol, 0)
	for i, counts := range symbolCounts {
		if counts >= maxCount {
			winSymbols = append(winSymbols, i)
		}
		// fmt.Printf("Symbol %d counts: %+v\n", i, counts+1)
	}

	for _, winSymbol := range winSymbols {
		matchingSymbolsCount := symbolCounts[winSymbol]
		payout, err := c.payTable.Get(winSymbol, matchingSymbolsCount)
		symbolsFound := make(symbols.Symbols, 0)
		for i := 0; i <= matchingSymbolsCount; i++ {
			symbolsFound = append(symbolsFound, winSymbol)
		}
		if err == nil && payout > 0 {
			wins = append(wins, win.NewWin(payout, symbolsFound, winSymbol))
		}
	}
	return wins, nil
}
