// An exercise that takes ever n seconds from Time.tick then

// do time.Tick millisecond then for each result multiply it by a random seed, you'll ha a unique result, then put them into a channel,
// then see if you can select on them by even and odd.

package numbernoise

import (
	"math/rand"
	"time"
)

const (
	BufferLimit = 100000
)

// p is the precision of the tick - e.g. a unique number every p
// l is the lifetime of the ticker, after which the returned channel will be closed
func RandomNSeconds(p, l time.Duration) <-chan int {
	c := time.NewTicker(p)
	w := time.NewTimer(l)

	nums := make(chan int, BufferLimit)

	go func() {
		ok := true
		for ok {
			select {
			case tick := <-c.C:
				num := int(int64(tick.UnixMilli())) * rand.Int()
				nums <- num
			case <-w.C:
				c.Stop()
				w.Stop()
				close(nums)
				ok = false
			}
		}
	}()

	return nums
}

func EvenOdds(p, l time.Duration) (<-chan int, <-chan int) {
	evens := make(chan int, BufferLimit)
	odds := make(chan int, BufferLimit)

	go func() {
		for rn := range RandomNSeconds(p, l) {
			if rn%2 == 0 {
				evens <- rn
			} else {
				odds <- rn
			}

			if rn == 0 { // zero value received
				close(evens)
				close(odds)
			}
		}
	}()

	return evens, odds
}
