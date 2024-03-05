package symbols

import (
	"embed"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed symbols.txt
var symbols embed.FS

const skipSymbol = -1

func parseReels(data [][]string) ([]Symbols, error) {
	reels := make([]Symbols, 5) // Предполагаем, что у нас 5 барабанов
	for _, row := range data {
		for reelIndex, symbolStr := range row {
			if reelIndex >= len(reels) {
				continue
			}
			if symbolStr == "" || symbolStr == fmt.Sprintf("%d", skipSymbol) {
				continue // Пропускаем пустые строки и строки со значением skipSymbol
			}
			symbol, err := strconv.Atoi(symbolStr)
			if err != nil {
				return nil, fmt.Errorf("error converting string to int: %v", err)
			}
			reels[reelIndex] = append(reels[reelIndex], symbol)
		}
	}
	return reels, nil
}

// ReadReels - read symbols from file
func ReadReels() ([]Symbols, error) {
	// обрати внимание, что в файле symbols.txt символы разделены через \t
	// и что в конце каждой строки есть \n
	// символ -1 нужен только для выравнивания таблицы
	data, err := reader.Read(symbols, "symbols.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	symbols, err := parseReels(data)
	if err != nil {
		return nil, fmt.Errorf("parseReels(): %w", err)
	}

	return symbols, nil
}
