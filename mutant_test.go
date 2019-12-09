package mutant

import (
	"testing"
)

var (
	human = []string{"ATGCGA", "AAGTGC", "ATATTT", "AGACGG", "GCGTCA", "TCACTG"}
	mutantBase = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	mutantFastDetection = []string{"ATGCGA", "AAGTGC", "ATAGTC", "AGGAGC", "GGCTCC", "TCACTG"}
	mutantHorizontal = []string{"AAAAGA", "CAGTGC", "TTTTGT", "AGAAGG", "CCCCTA", "TCACTG"}
	mutantVertical = []string{"ATGCGA", "AAGTGC", "ATATTC", "AGACGC", "GCGTCC", "TCACTG"}
	mutantDiagonal = []string{"ATGCGA", "AAGTGC", "ATAGTC", "GGGAGC", "GGCTCC", "TCACTG"}
	mutantBigMatrix = []string{"ATGCGAA", "CAGTGCC", "TTATGTT", "AGAAGGG", "CCCCTAA", "TCACTGG", "ATGCGAA"}
	invalidMatrix = []string{"AGA", "GC"}
)

func BenchmarkIsMutant(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsMutant(mutantBigMatrix)
	}
}

func TestIsMutant(t *testing.T) {
	testCases := []struct {
		dna []string
		isMutant bool
	}{
		{human, false},
		{mutantBase, true},
		{mutantFastDetection, true},
		{mutantHorizontal, true},
		{mutantVertical, true},
		{mutantDiagonal, true},
		{mutantBigMatrix, true},
	}

	for _, testCase := range testCases {
		result := IsMutant(testCase.dna)
		if result != testCase.isMutant {
			t.Errorf("DNA %+v expected to be mutant? %t, got: %t",
				testCase.dna,
				testCase.isMutant,
				result,
			)
		}
	}
}

func BenchmarkIsSquareMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsMutant(mutantBigMatrix)
	}
}

func TestIsSquareMatrix(t *testing.T) {
	testCases := []struct {
		dna []string
		isValid bool
	}{
		{human, true},
		{invalidMatrix, false},
	}
	for _, testCase := range testCases {
		result := IsSquareMatrix(testCase.dna)
		if result != testCase.isValid {
			t.Errorf("In DNA %+v should be valid? %t got: %t",
				testCase.dna,
				testCase.isValid,
				result,
			)
		}
	}
}

func TestSearchMutantSubSequence(t *testing.T) {
	testCases := []struct {
		dna []string
		startPosition Point
		found int
	}{
		{human, Point{0, 0}, 1},
		{mutantBase, Point{0, 0}, 1},
		{mutantFastDetection, Point{0, 0}, 2},
	}

	for _, testCase := range testCases {
		// reset subsequences
		mutantSubsequences = 0
		searchMutantSubSequence(testCase.dna, testCase.startPosition)
		if mutantSubsequences != testCase.found {
			t.Errorf("In DNA %+v from position (%d, %d) found: %d, expected: %d",
				testCase.dna,
				testCase.startPosition.x,
				testCase.startPosition.y,
				mutantSubsequences,
				testCase.found,
			)
		}
	}
}

func TestGetCharsLeftCount(t *testing.T) {
	testCases := []struct {
		limit    int
		position Point
		dir      Point
		result   int
	}{
		// test direction 1,0
		{5, Point{0, 0}, Point{1, 0}, 6},
		// test direction 1,1
		{5, Point{1, 1}, Point{1, 1}, 5},
		// test direction 0,1
		{5, Point{0, 3}, Point{0, 1}, 3},
		// test direction -1, 1
		{5, Point{2, 0}, Point{-1, 1}, 3},
	}

	for _, testCase := range testCases {
		result := getCharsLeftCount(testCase.limit, testCase.position, testCase.dir)
		if result != testCase.result {
			t.Errorf("Direction (%d, %d) with limit %d, starting from: (%d, %d), expect %d characters left, has %d",
				testCase.position.x,
				testCase.position.y,
				testCase.limit,
				testCase.dir.x,
				testCase.dir.y,
				testCase.result,
				result,
			)
		}
	}
}

func TestOutOfBound(t *testing.T) {
	testCases := []struct {
		point  Point
		limit  int
		result bool
	}{
		// valid case
		{Point{5, 1}, 5, false},
		// bellow boundaries X
		{Point{-4, 1}, 5, true},
		// exceed boundaries X
		{Point{7, 1}, 5, true},
		// bellow boundaries Y
		{Point{0, -2}, 5, true},
		// exceed boundaries Y
		{Point{0, 7}, 5, true},
	}

	for _, testCase := range testCases {
		result := outOfBound(testCase.point, testCase.limit)
		if result != testCase.result {
			var msg string
			if testCase.result {
				msg = "should"
			} else {
				msg = "should not"
			}
			t.Errorf("With Limit %d, Case: (%d, %d) %s exceed boundaries", testCase.limit, testCase.point.x, testCase.point.y, msg)
		}
	}
}
