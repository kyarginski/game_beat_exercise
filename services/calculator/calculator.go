package calculator

import (
	"fmt"

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
	var wins []win.Win

	for _, line := range c.lines {
		var matchingSymbols symbols.Symbols
		symbolCount := 0
		var winSymbol symbols.Symbol

		for i, index := range line.GetIndices() {
			reelIndex := index % len(spinSymbols)   // Получаем индекс барабана
			symbolIndex := index / len(spinSymbols) // Получаем индекс символа в барабане

			// Проверяем на выход за границы
			if reelIndex >= len(spinSymbols) || symbolIndex >= len(spinSymbols[reelIndex]) {
				return nil, fmt.Errorf("index out of range")
			}

			currentSymbol := spinSymbols[reelIndex][symbolIndex]

			if i == 0 { // Для первого символа в линии
				winSymbol = currentSymbol                            // Устанавливаем выигрышный символ
				if winSymbol == WILD && len(line.GetIndices()) > 1 { // Если первый символ WILD, ищем реальный символ далее
					continue // Пропускаем дальнейшую логику в этой итерации, чтобы найти следующий непустой символ
				}
			}

			if currentSymbol != winSymbol && currentSymbol != WILD {
				break // Если текущий символ не совпадает с выигрышным и не является WILD, прерываем
			}

			if winSymbol == WILD && currentSymbol != WILD { // Если выигрышный символ был WILD, но мы нашли реальный
				winSymbol = currentSymbol // Обновляем выигрышный символ
			}

			symbolCount++                                            // Увеличиваем количество подходящих символов
			matchingSymbols = append(matchingSymbols, currentSymbol) // Добавляем текущий символ к подходящим
		}

		if symbolCount > 0 {
			// Если количество подходящих символов достаточно для выигрыша
			payout, err := c.payTable.Get(winSymbol, symbolCount)
			if err == nil {
				wins = append(wins, win.NewWin(payout, matchingSymbols, winSymbol))
			}
		}
	}

	return wins, nil
}
