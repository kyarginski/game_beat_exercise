package paytable

import (
	"errors"
	"fmt"

	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
)

// Payout - таблица выплаты определенного символа
type Payout []uint64

// PayTable - таблица выплат всех символов
type PayTable struct {
	symbolPayouts map[symbols.Symbol]Payout
}

func NewPayTable(symbolPayouts map[symbols.Symbol]Payout) *PayTable {
	return &PayTable{symbolPayouts: symbolPayouts}
}

func (p *PayTable) Get(s symbols.Symbol, index int) (uint64, error) {
	if payout, ok := p.symbolPayouts[s]; ok {
		if index >= 0 && index < len(payout) {
			return payout[index], nil
		} else {
			return 0, fmt.Errorf("invalid index: %d for symbol: %v", index, s)
		}
	}
	return 0, errors.New("symbol not found")
}
