package naming

import (
	"math/rand"
)

var names = []string{"Mike", "Jack", "Tim", "Steven"}

func GiveMeAName() string {
	return names[rand.Int()%3]
}
