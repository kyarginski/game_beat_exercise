package lines

import (
	"embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/lines_err.txt
var testFS embed.FS

//go:embed testdata/lines_no_len.txt
var testFS2 embed.FS

func TestReadLines(t *testing.T) {
	// good case
	res, err := ReadLines()
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 5, len(res))

	// bad case (error element)
	lines = testFS
	_, err = ReadLinesFromFile("testdata/lines_err.txt")
	require.Error(t, err, "strconv.Atoi(): strconv.Atoi: parsing \"a\": invalid syntax")

	// bad case (error length)
	lines = testFS2
	res, err = ReadLinesFromFile("testdata/lines_no_len.txt")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.NotEqual(t, 5, len(res))

}
