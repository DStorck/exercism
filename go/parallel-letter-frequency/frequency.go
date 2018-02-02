package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(blobs []string) FreqMap {
	ch := make(chan FreqMap, len(blobs))
	for _, text := range blobs {
		go func(str string) { ch <- Frequency(str) }(text)
	}

	freqs := FreqMap{}
	for range blobs {
		for letter, value := range <-ch {
			freqs[letter] += value
		}
	}
	return freqs
}
