package main

import (
	"fmt"
	"math/rand"
	"slices"
)

// keys lists keys in order of the Cycle of Fifths.
var keys = []string{"C", "G", "D", "A", "E", "B/Cb", "Gb/F#", "Db/C#", "Ab", "Eb", "Bb", "F"}

// scalePositions lists the scale positions used in movable scales.
var scalePositions = []string{"1", "2", "3", "4", "5"}

var scaleTypes = []string{"major", "minor pentatonic"}

// rhythms lists rhythms used in exercises.
var rhythms = []string{"1", "1/4", "1/3", "1/16"}

// warmups prints a sequence for a left hand fingering exercise with an ascending and descending pattern.
// Exercises start with the first finger with the remaining sequence randomly generated.
// ascending:  [1, 3, 2, 4]
// descending: [4, 2, 3, 1]
func warmups() {
	const patternLength = 4
	startingFinger := "1"
	remainingFingers := []string{"2", "3", "4"}

	rand.Shuffle(len(remainingFingers), func(i, j int) {
		remainingFingers[i], remainingFingers[j] = remainingFingers[j], remainingFingers[i]
	})

	warmupPattern := make([]string, 0, patternLength)
	warmupPattern = append(warmupPattern, startingFinger)
	warmupPattern = append(warmupPattern, remainingFingers...)

	// rhythmCount is used as the upper bound for selecting random rhythms
	rhythmCount := len(rhythms)

	fmt.Printf("warmup pattern ascending %v rhythm: %s (start at 5th fret)", warmupPattern, rhythms[rand.Intn(rhythmCount)])
	fmt.Println()

	slices.Reverse(warmupPattern)
	fmt.Printf("warm pattern descending %v rhythm: %s", warmupPattern, rhythms[rand.Intn(rhythmCount)])
	fmt.Println()
}

// scales prints scales to run across all five positions.
// A single key is used for each scale.
func scales() {
	selectedKey := keys[rand.Intn(len(keys))]
	for _, scaleType := range scaleTypes {
		fmt.Printf("%s %s scale all 5 positions", selectedKey, scaleType)
		fmt.Println()
	}
}

// positions prints a scale position to practice in all scales.
func positions() {
	selectedKey := keys[rand.Intn(len(keys))]
	selectedScalePosition := scalePositions[rand.Intn(len(scalePositions))]

	fmt.Printf("%s major position %s", selectedKey, selectedScalePosition)
	fmt.Println()
	fmt.Printf("%s minor pentatonic position %s", selectedKey, selectedScalePosition)
	fmt.Println()
}

func main() {
	warmups()
	scales()
	positions()
}
