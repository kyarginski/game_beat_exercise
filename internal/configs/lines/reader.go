package lines

import (
	"embed"
	"fmt"
	"strconv"

	"github.com/releaseband/golang-developer-test/internal/configs/reader"
)

//go:embed lines.txt
var lines embed.FS

func parseLine(data []string) (*Line, error) {
	var indices []int

	for _, part := range data {
		if part == "" {
			continue
		}
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("strconv.Atoi(): %w", err)
		}
		indices = append(indices, num)
	}

	return &Line{indices: indices}, nil
}

func ReadLines() (Lines, error) {
	return ReadLinesFromFile("lines.txt")
}

func ReadLinesFromFile(fileName string) (Lines, error) {
	data, err := reader.Read(lines, fileName)
	if err != nil {
		return nil, fmt.Errorf("reader.Read(): %w", err)
	}

	resp := make([]Line, len(data))
	for i, str := range data {
		line, err := parseLine(str)
		if err != nil {
			return nil, fmt.Errorf("parseLines(): %w", err)
		}

		resp[i] = *line
	}

	return resp, nil
}
