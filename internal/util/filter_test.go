package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xc5216/circle-console-go/internal/util"
)

func TestFilter(t *testing.T) {
	t.Run("Filter ints", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{2, 4}

		result := util.Filter(input, func(n int) bool {
			return n%2 == 0
		})

		assert.Equal(t, expected, result)
	})

	t.Run("Filter strings", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry", "date"}
		expected := []string{"apple"}

		result := util.Filter(input, func(s string) bool {
			return len(s) == 5
		})

		assert.Equal(t, expected, result)
	})

	t.Run("Filter structs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		input := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}
		expected := []Person{
			{"Alice", 30},
			{"Charlie", 35},
		}

		result := util.Filter(input, func(p Person) bool {
			return p.Age > 25
		})

		assert.Equal(t, expected, result)
	})
}
