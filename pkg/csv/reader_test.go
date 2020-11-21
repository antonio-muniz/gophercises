package csv_test

import (
	"strings"
	"testing"

	"github.com/antonio-muniz/gophercises/pkg/csv"
	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	reader := strings.NewReader(strings.Join([]string{
		"name,city_of_birth,date_of_birth",
		"Toninho,Poá,1994-06-14",
		"Moarah,Suzano,1995-01-04",
		"",
	}, "\n"))
	rows, err := csv.Read(reader)
	require.NoError(t, err)
	require.Equal(t, [][]string{
		{"name", "city_of_birth", "date_of_birth"},
		{"Toninho", "Poá", "1994-06-14"},
		{"Moarah", "Suzano", "1995-01-04"},
	}, rows)
}
