package mutant

//represent an XY point
type Point struct {
	x int
	y int
}

const (
	// quantity of sequences found to be mutant
	subSequencesToBeMutant int = 2
	// quantity of repeated characters to be a mutant subsequence
	repeatCharsToBeMutant int = 4
)

var (
	// found mutant subsequences in current scan
	mutantSubsequences int
	// available directions
	directions = [4]Point{{1, 0}, {1, 1}, {0, 1}, {-1, 1}}
)

// detect if []string DNA is mutant or human
func IsMutant(dna []string) bool {
	// set found sequences to 0
	mutantSubsequences = 0
	// loop through Y axis
	for row, _ := range dna {
		// loop through X axis
		for col, _ := range dna[0] {
			// search in matrix starting by current Point
			searchMutantSubSequence(dna, Point{col, row})
			// if is mutant return instantly to avoid overprocessing
			if mutantSubsequences == subSequencesToBeMutant {
				return true
			}
		}
	}
	return false
}

// look mutant sequence in every direction
func searchMutantSubSequence(dna []string, startPosition Point) {
	for _, dir := range directions {
		// get characters left in current direction
		var leftCharsCount = getCharsLeftCount(len(dna)-1, startPosition, dir)
		// discard current direction if chars left are less
		// than required for match a mutant subsequence
		if leftCharsCount < repeatCharsToBeMutant {
			continue
		}
		var (
			// current point
			pointer Point
			// characters repetition
			charCount int
			// last tested character
			lastChar uint8
		)
		// loop trough all characters to validate concurrent repeated characters
		for currentPos := 0; currentPos < leftCharsCount; currentPos++ {
			// increase pointer in X direction
			pointer.x = startPosition.x + (dir.x * currentPos)
			// increase pointer in Y direction
			pointer.y = startPosition.y + (dir.y * currentPos)
			// check if last character is equal to current character
			if lastChar == dna[pointer.y][pointer.x] {
				// increase counter
				charCount++
				// validate if reach repeatCharsToBeMutant
				if charCount >= repeatCharsToBeMutant {
					mutantSubsequences++
					break
				}
			} else {
				// if current char is not equal to last char, reset
				charCount = 1
				lastChar = dna[pointer.y][pointer.x]
			}
		}

		// if is mutant return instantly to avoid overprocessing
		if mutantSubsequences == subSequencesToBeMutant {
			return
		}
	}
}

// Get characters left in specific position
func getCharsLeftCount(limit int, position Point, dir Point) int {
	left := 0
	// increase counter while position is not out of boundaries of grid
	for !outOfBound(position, limit) {
		position.x += dir.x
		position.y += dir.y
		left++
	}

	return left
}

// validate if Point is out of limits
func outOfBound(point Point, limit int) bool {
	return point.x > limit || point.x < 0 || point.y > limit || point.y < 0
}

// detect if matrix is NxN
func IsSquareMatrix(dna []string) bool {
	for _, subSequence := range dna {
		if len(subSequence) != len(dna) {
			return false
		}
	}

	return true
}

// detect if []string DNA has invalid characters
func HasInvalidCharacters(dna []string) bool {
	for _, row := range dna {
		// loop through X axis
		for _, char := range row {
			if char != 65 && char != 84 && char != 67 && char != 71 {
				return true
			}
		}
	}

	return false
}