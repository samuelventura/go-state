package state

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestReversedOrder(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mn := 1000
	mmn := mn * mn
	r := NewSorted()
	ar := make([]string, 0, mn)
	for i := 0; i < mn; i++ {
		n := rand.Intn(mmn)
		s := fmt.Sprintf("%08d", n)
		for {
			if r.Get(s) == nil {
				break
			}
			n = rand.Intn(mmn)
			s = fmt.Sprintf("%08d", n)
		}
		ar = append(ar, s)
		r.Set(s, s)
	}
	assert.Equal(t, mn, r.Count())
	for i, s := range r.Names() {
		assert.Equal(t, s, ar[i])
	}
	for i, v := range r.Values() {
		assert.Equal(t, v.(string), ar[i])
	}
}
