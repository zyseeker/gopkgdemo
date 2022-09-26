package naming

import (
	"math/rand"
)

var names = []string{"Mike", "Jack", "Tim"}

func GiveMeAName() string {
	return names[rand.Int()%3]
}
