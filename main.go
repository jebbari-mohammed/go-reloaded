package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid input")
		return
	}
	sample := os.Args[1]
	result := os.Args[2]

	if filepath.Ext(sample) != ".txt" || filepath.Ext(result) != ".txt" {
		log.Fatal("The file should be .txt")
	}
	if result == "main.go" {
		fmt.Println("haha nice try")
		return
	}

	lire, err := os.Open(sample)
	if err != nil {
		panic(err)
	}
	defer lire.Close()
	var allLines []string
	scanner := bufio.NewScanner(lire)
	reg := regexp.MustCompile(`^\d+\)$`)

	// Loop through the words and process tags
	voiles := []rune{'a', 'o', 'e', 'h', 'u', 'i', 'A', 'O', 'E', 'H', 'U', 'I'}
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		ponc := regexp.MustCompile(`^[,.?!:;]+$`)
		regBin := regexp.MustCompile(`^[01]+$`)
		regHex := regexp.MustCompile(`^[0-9a-fA-F]+$`)
		regPlus := regexp.MustCompile(`^\+\d+\)$`)
		regMoins := regexp.MustCompile(`^\-\d+\)$`)
		for i := 0; i < len(words); i++ {
			switch words[i] {
			case "(up)":
				if i == 0 {
					words = words[i+1:]
					i--
				}
				if i > 0 {
					for j := i; j > 0; j-- {
						if ponc.MatchString(words[j-1]) {
							continue
						} else {
							words[j-1] = strings.ToUpper(words[j-1])
							break
						}
					}
					words = append(words[:i], words[i+1:]...) // Remove the tag
					i--                                       // Adjust the index after removal
				}

			case "(low)":
				if i == 0 {
					words = words[i+1:]
					i--
				}
				if i > 0 {
					for j := i; j > 0; j-- {
						if ponc.MatchString(words[j-1]) {
							continue
						} else {
							words[j-1] = strings.ToLower(words[j-1])
							break
						}
					}
					words = append(words[:i], words[i+1:]...) // Remove the tag
					i--                                       // Adjust the index after removal
				}

			case "(cap)":
				if i == 0 {
					words = words[i+1:]
					i--
				}
				if i > 0 {
					for j := i; j > 0; j-- {
						if ponc.MatchString(words[j-1]) {
							continue
						} else {
							words[j-1] = titleCase(words[j-1])
							break
						}
					}
					words = append(words[:i], words[i+1:]...) // Remove the tag
					i--                                       // Adjust the index after removal
				}

			case "(hex)":
				if i == 0 {
					words = words[i+1:]
					i--
				}
				if i > 0 {
					for j := i; j > 0; j-- {
						if ponc.MatchString(words[j-1]) {
							continue
						} else if regHex.MatchString(words[j-1]) {
							words[j-1] = HextoInt(words[j-1])
							break
						} else {
							fmt.Println("error : invalid hex number")
							return
						}
					}
					words = append(words[:i], words[i+1:]...) // Remove the tag
					i--                                       // Adjust the index after removal
				}

			case "(bin)":
				if i == 0 {
					words = words[i+1:]
					i--
				}
				if i > 0 {
					for j := i; j > 0; j-- {
						if ponc.MatchString(words[j-1]) {
							continue
						} else if regBin.MatchString(words[j-1]) {
							words[j-1] = BintoInt(words[j-1])
							break
						} else {
							fmt.Println("error : invalid bin number")
							return
						}
					}
					words = append(words[:i], words[i+1:]...) // Remove the tag
					i--                                       // Adjust the index after removal
				}
			case "(cap,":
				if i == 0 && i+1 <= len(words)-1 && (reg.MatchString(words[i+1]) || regPlus.MatchString(words[i+1]) || regMoins.MatchString(words[i+1])) {
					words = words[i+2:]
					i--

				}
				if i > 0 && i+1 <= len(words)-1 && (reg.MatchString(words[i+1]) || regPlus.MatchString(words[i+1])) {
					a := strings.TrimSuffix(words[i+1], ")")
					numero, _ := strconv.Atoi(string(a))
					if numero > len(words[:i]) {
						numero = len(words[:i])
					}
					countponc := 0
					countwords := 0

					for r := i - 1; countwords < numero && r >= 0; r-- {
						if ponc.MatchString(words[r]) {
							countponc++
						} else {
							countwords++
						}
						total := countponc + countwords

						for j := 1; j <= total; j++ {
							words[i-j] = titleCase(words[i-j])
						}

					}
					words = append(words[:i], words[i+2:]...)
					i--

				}

			case "(low,":
				if i == 0 && i+1 <= len(words)-1 && (reg.MatchString(words[i+1]) || regPlus.MatchString(words[i+1]) || regMoins.MatchString(words[i+1])) {
					words = words[i+2:]
					i--

				}
				if i > 0 && i+1 <= len(words)-1 && (reg.MatchString(words[i+1]) || regPlus.MatchString(words[i+1])) {
					a := strings.TrimSuffix(words[i+1], ")")
					numero, _ := strconv.Atoi(string(a))
					if numero > len(words[:i]) {
						numero = len(words[:i])
					}
					countponc := 0
					countwords := 0

					for r := i - 1; countwords < numero && r >= 0; r-- {
						if ponc.MatchString(words[r]) {
							countponc++
						} else {
							countwords++
						}
						total := countponc + countwords

						for j := 1; j <= total; j++ {
							words[i-j] = strings.ToLower(words[i-j])
						}

					}
					words = append(words[:i], words[i+2:]...)
					i--

				}
			case "(up,":
				if i == 0 && i+1 <= len(words)-1 && (reg.MatchString(words[i+1]) || regPlus.MatchString(words[i+1]) || regMoins.MatchString(words[i+1])) {
					words = words[i+2:]
					i--

				}
				if i > 0 && i+1 <= len(words)-1 && (reg.MatchString(words[i+1]) || regPlus.MatchString(words[i+1])) {
					a := strings.TrimSuffix(words[i+1], ")")
					numero, _ := strconv.Atoi(string(a))
					if numero > len(words[:i]) {
						numero = len(words[:i])
					}
					countponc := 0
					countwords := 0

					for r := i - 1; countwords < numero && r >= 0; r-- {
						if ponc.MatchString(words[r]) {
							countponc++
						} else {
							countwords++
						}
						total := countponc + countwords

						for j := 1; j <= total; j++ {
							words[i-j] = strings.ToUpper(words[i-j])
						}

					}
					words = append(words[:i], words[i+2:]...)
					i--

				}
			}
		}
		for i := 0; i < len(words); i++ {
			switch words[i] {
			case "a":
				for _, char := range voiles {
					if i+1 < len(words) {
						if char == rune(words[i+1][0]) && i+1 < len(words) {
							words[i] = "an"
						}
					}
				}
			case "A":
				for _, char := range voiles {
					if i+1 < len(words) {
						if char == rune(words[i+1][0]) && i+1 < len(words) {
							words[i] = "An"
						}
					}
				}
			}
		}
		Lines := strings.Join(words, " ")
		punctuations := []rune{'.', ',', '!', '?', ';', ':'}
		runes := []rune(Lines)

		for i := 1; i < len(runes); i++ {
			for _, ponct := range punctuations {
				if runes[i] == ponct && runes[i-1] == ' ' {
					runes = append(runes[:i-1], runes[i:]...)
					i-- // Adjust index to account for removed character
				}
				if runes[i] == ponct && i != len(runes)-1 && (runes[i+1] != ' ' && runes[i+1] != ponct) {
					// Insert a space before punctuation if there isn't already one
					runes = append(runes[:i+1], append([]rune{' '}, runes[i+1:]...)...)
					i++ // Move past the inserted space
				}
			}
		}

		// quot
		count := 0
		for i := 0; i < len(runes)-1; i++ {
			if i > 0 && runes[i] == '\'' && runes[i-1] != ' ' && runes[i+1] == ' ' {
				runes[i] = ' '
				runes[i+1] = '\''
			}
		}

		for i := 0; i <= len(runes)-1; i++ {
			if i == 0 && runes[i] == '\'' && len(runes) > 1 {
				count++
				if count%2 == 1 && runes[i+1] == ' ' {
					runes[i+1] = '\a'
				} else {
					continue
				}
			} else if i == len(runes)-1 && len(runes) > 1 {
				count++
				if count%2 == 0 && runes[i-1] == ' ' {
					runes[i-1] = '\a'
				} else {
					continue
				}
			}
			if i > 0 && i < len(runes)-1 && runes[i] == '\'' && (runes[i-1] >= 33 && runes[i-1] <= 126 && runes[i+1] >= 33 && runes[i+1] <= 126) {
				continue
			}
			if i > 0 && i < len(runes)-1 && runes[i] == '\'' && !(runes[i-1] >= 33 && runes[i-1] <= 126 && runes[i+1] >= 33 && runes[i+1] <= 126) {
				count++

				if count%2 == 0 && (runes[i+1] >= 33 && runes[i+1] <= 126) {
					runes[i] = ' '
					runes[i-1] = '\''
				}
				if count%2 == 1 && runes[i+1] == ' ' {
					runes[i+1] = '\a'
				} else if count%2 == 0 && runes[i-1] == ' ' {
					runes[i-1] = '\a'
				}
			}

		}

		finalresult := strings.ReplaceAll(string(runes), "\a", "")
		allLines = append(allLines, finalresult)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file: ", err)
	}

	finaltxt := strings.Join(allLines, "\n")
	finalfile, err := os.Create(result)
	if err != nil {
		panic(err)
	}
	if _, err := io.WriteString(finalfile, finaltxt); err != nil {
		panic(err)
	}
}

func HextoInt(hex string) string {
	number, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(number)
}

func BintoInt(bin string) string {
	number, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(number)
}

func titleCase(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(strings.ToLower(word))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
