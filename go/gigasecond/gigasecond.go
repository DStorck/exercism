package gigasecond

// A gigasecond is 10^9 (1,000,000,000) seconds.
import "time"

func AddGigasecond(t time.Time) time.Time {
	a := t.Add(time.Second * time.Duration(1000000000))
	return a
}

// fmt.Println(time.Now().Add(time.Second * time.Duration(1000000000)))
