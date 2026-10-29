package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	delta = t4.Sub(t1) - t3.Sub(t2)
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Microsecond)

	t2 := s_clock.Now()
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	t3 := s_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	minDelta := math.MaxFloat64
	minTheta := math.MaxFloat64
	for i := 0; i < 10; i++ {
		serverClock := &Clock{offset: 0}
		clientClock := &Clock{offset: 350 * time.Millisecond}

		initialOffser := clientClock.offset
		theta, delta := measure(clientClock, serverClock)

		if float64(delta) < minDelta {
			minDelta = float64(delta)
		}
		if float64(theta) < minTheta {
			minTheta = float64(theta)
		}

		fmt.Println(theta, delta)
		clientClock.offset += theta
		finalOffset := clientClock.offset

		initialError := math.Abs(float64(initialOffser))
		finalError := math.Abs(float64(finalOffset))

		if finalError < initialError {
			fmt.Println(time.Duration(initialError), time.Duration(finalError))
		} else {
			fmt.Printf("Offset nu s a redus")
		}
	}

	fmt.Println(minDelta, minTheta)
}
