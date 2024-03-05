package game

import (
	"github.com/releaseband/golang-developer-test/internal/game/result"
	"github.com/releaseband/golang-developer-test/internal/rng"
)

// RoundCost - функция, которая возвращает стоимость одного раунда
func RoundCost(linesCount int) uint64 {
	return uint64(linesCount)
}

type Slot struct {
	generator  Generator
	calculator Calculator
	roundCost  uint64
}

func newSlot(generator Generator, calculator Calculator, roundCost uint64) *Slot {
	return &Slot{generator: generator, calculator: calculator, roundCost: roundCost}
}

func (s *Slot) Spin(rng rng.RNG) (*result.Round, error) {
	symbols, err := s.generator.Generate(rng)
	if err != nil {
		return nil, err // Возвращаем ошибку, если генерация не удалась.
	}

	// Вычисляем выигрыши на основе сгенерированных символов.
	wins, err := s.calculator.Calculate(symbols)
	if err != nil {
		return nil, err // Возвращаем ошибку, если расчет не удался.
	}

	// Создаем результат раунда.
	roundResult := result.NewRound(symbols, wins, s.roundCost)

	return roundResult, nil
}

func (s *Slot) RoundCost() uint64 {
	return s.roundCost
}
