package csv_test

import (
	"testing"

	"github.com/antonio-muniz/gophercises/pkg/csv"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	rows, err := csv.Read("C:\\Users\\ahnm_\\go\\src\\github.com\\antonio-muniz\\gophercises\\test\\data\\pkg\\csv\\file.csv")
	require.NoError(t, err)
	require.Equal(t, [][]string{
		{"name", "city_of_birth", "date_of_birth"},
		{"Toninho", "Poá", "1994-06-14"},
		{"Moarah", "Suzano", "1995-01-04"},
	}, rows)
}
