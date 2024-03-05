package paytable

import (
	"embed"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
	"github.com/releaseband/golang-developer-test/internal/configs/symbols"
)

//go:embed pay_table.txt
var payTable embed.FS

func parsePayouts(data [][]string) (map[symbols.Symbol]Payout, error) {
	payouts := make(map[symbols.Symbol]Payout)
	for i, line := range data {
		payout := Payout{}
		for _, str := range line {
			p, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing payout for symbol %d: %v", i, err)
			}
			payout = append(payout, p)
		}
		payouts[i] = payout
	}
	return payouts, nil
}

func ReadPayTable() (*PayTable, error) {
	data, err := reader.Read(payTable, "pay_table.txt")
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	payouts, err := parsePayouts(data)
	if err != nil {
		return nil, fmt.Errorf("parsePayouts(): %w", err)
	}

	return NewPayTable(payouts), nil
}
