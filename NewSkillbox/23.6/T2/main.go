package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sentences := [4]string{"Hello world", "Hello skillSbox", "Привет Мир", "Привет skillSbox"}
	chars := [5]rune{'H', 'E', 'L', 'S', 'М'}

	result, err := parseTest(sentences, chars)

	if err != nil {
		fmt.Print(err)
		return
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			if result[i][j] < strconv.Itoa(0) {
				continue
			}
			fmt.Printf("<%s> индекс в последнем слове предложения <%s> = <%v>\n", strconv.QuoteRune(chars[j]), sentences[i], result[i][j])
		}
	}
}

func parseTest(sentences [4]string, chars [5]rune) ([][]string, error) {
	var indexes [][]string

	if len(sentences) == 0 || len(chars) == 0 {
		return nil, errors.New("Один или оба массива во входящих параметров пустые!")
	}

	for i := 0; i < len(sentences); i++ {
		var temp = make([]string, 0)
		var index int
		for j := 0; j < len(chars); j++ {
			var currentSentenceWords = strings.FieldsFunc(sentences[i], func(r rune) bool {
				return r == ',' || r == '.' || r == ' '
			})

			var currentSentenceLastWord = currentSentenceWords[len(currentSentenceWords)-1]
			index = strings.IndexRune(currentSentenceLastWord, chars[j])
			temp = append(temp, strconv.Itoa(index))
		}

		indexes = append(
			indexes,
			temp,
		)
	}
	return indexes, nil
}
