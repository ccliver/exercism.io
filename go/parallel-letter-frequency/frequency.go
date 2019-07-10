package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calls Frequency in a goroutine for each wordset in words, then combines and returns the result.
func ConcurrentFrequency(words []string) FreqMap {
	maps := FreqMap{}
	mapBuf := make(chan FreqMap, len(words))
	var wg sync.WaitGroup
	wg.Add(len(words))

	callFrequency := func(str string) {
		mapBuf <- Frequency(str)
		wg.Done()
	}

	for _, s := range words {
		go callFrequency(s)
	}
	wg.Wait()

	close(mapBuf)
	for m := range mapBuf {
		for k, v := range m {
			if _, ok := maps[k]; ok {
				maps[k] += v
			} else {
				maps[k] = v
			}
		}
	}

	return maps
}
