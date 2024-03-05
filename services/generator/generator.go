package generator

import (
	"fmt"

	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

type Symbols struct {
	rowsCount int
	gameTapes []symbols.Symbols
}

func NewSymbols(rowsCount int, gameTapes []symbols.Symbols) *Symbols {
	return &Symbols{rowsCount: rowsCount, gameTapes: gameTapes}
}

func (s *Symbols) Generate(rng rng.RNG) (symbols.Reels, error) {
	reels := make(symbols.Reels, len(s.gameTapes))
	for i, tape := range s.gameTapes {
		if len(tape) < s.rowsCount {
			return nil, fmt.Errorf("tape %d is shorter than rowsCount", i)
		}
		startIndex := rng.Random(0, uint32(len(tape)-s.rowsCount))
		reels[i] = tape[startIndex : startIndex+uint32(s.rowsCount)]
	}
	return reels, nil
}

func (s *Symbols) GetReelSymbols(reelIndex int, rowIndex int) symbols.Symbols {
	if reelIndex < 0 || reelIndex >= len(s.gameTapes) {
		return nil
	}

	reel := s.gameTapes[reelIndex]
	symbolsCount := len(reel)
	if symbolsCount == 0 {
		return nil
	}

	result := make(symbols.Symbols, s.rowsCount)
	for i := 0; i < s.rowsCount; i++ {
		// Вычисляем индекс символа, учитывая возможный циклический переход через конец списка
		symbolIndex := (rowIndex + i) % symbolsCount
		result[i] = reel[symbolIndex]
	}

	return result
}
