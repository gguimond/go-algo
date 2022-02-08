package boggle

import "fmt"

var N int = 3
var M int = 3

func isWord(s string) bool {
	for _, word := range []string{"GEEKS", "FOR", "QUIZ", "GO"} {
		if s == word {
			return true
		}
	}
	return false
}

func findWords(boggle [3][3]rune, visited *[3][3]bool, i int, j int, currentStr *string, words *[]string) {
	(*visited)[i][j] = true
	*currentStr = *currentStr + string(boggle[i][j])

	if isWord(*currentStr) {
		*words = append(*words, *currentStr)
	}

	row := i - 1
	for row <= i+1 && row < M {
		col := j - 1
		for col <= j+1 && col < N {
			if row >= 0 && col >= 0 && !(*visited)[row][col] {
				findWords(boggle, visited, row, col, currentStr, words)
			}
			col += 1
		}
		row += 1
	}
	*currentStr = "" + (*currentStr)[:len(*currentStr)-1]
	(*visited)[i][j] = false
}

func FindWords(boggle [3][3]rune) []string {
	visited := [3][3]bool{{false, false, false}, {false, false, false}, {false, false, false}}
	currentStr := ""
	words := []string{}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			findWords(boggle, &visited, i, j, &currentStr, &words)
		}
	}
	return words
}

func main() {
	boggle := [3][3]rune{{'G', 'I', 'Z'}, {'U', 'E', 'K'}, {'Q', 'S', 'E'}}
	fmt.Print(FindWords(boggle))
}
