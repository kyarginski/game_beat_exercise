package paytable

import (
	"testing"

	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
)

func TestPayTable_Get(t *testing.T) {
	testSymbolPayouts := map[symbols.Symbol]Payout{
		1: {0, 50, 200, 500},
		2: {0, 20, 100, 200},
	}
	payTable := NewPayTable(testSymbolPayouts)

	tests := []struct {
		name    string
		symbol  symbols.Symbol
		index   int
		want    uint64
		wantErr bool
	}{
		{"Valid symbol and index", 1, 2, 200, false},
		{"Valid symbol, index out of bounds", 1, 5, 0, true},
		{"Invalid symbol", 3, 0, 0, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := payTable.Get(tt.symbol, tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("PayTable.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PayTable.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
