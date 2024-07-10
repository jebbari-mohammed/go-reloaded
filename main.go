package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var sample = os.Args[1]

var result = os.Args[2]

func titleCase(word string) string {
	if len(word) == 0 {
		return word
	}
	// Convert the first character to uppercase and the rest to lowercase
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

func main() {
	file, err := os.ReadFile(sample)
	if err != nil {
		panic(err)
	}
	words := strings.Fields(string(file))
	reg := regexp.MustCompile(`^\d+\)$`)

	// Loop through the words and process tags
	for i := 0; i < len(words); i++ {
		switch words[i] {
		case "(up)":
			if i == 0 {
				words = words[i+1:]
			}
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...) // Remove the tag
				i--                                       // Adjust the index after removal
			}

		case "(low)":
			if i == 0 {
				words = words[i+1:]
			}
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...) // Remove the tag
				i--                                       // Adjust the index after removal
			}

		case "(cap)":
			if i == 0 {
				words = words[i+1:]
			}
			if i > 0 {
				words[i-1] = titleCase(words[i-1])
				words = append(words[:i], words[i+1:]...) // Remove the tag
				i--                                       // Adjust the index after removal
			}

		case "(hex)":
			if i == 0 {
				words = words[i+1:]
			}
			if i > 0 {
				words[i-1] = HextoInt(words[i-1])
				words = append(words[:i], words[i+1:]...) // Remove the tag
				i--                                       // Adjust the index after removal
			}

		case "(bin)":
			if i == 0 {
				words = words[i+1:]
			}
			if i > 0 {
				words[i-1] = BintoInt(words[i-1])
				words = append(words[:i], words[i+1:]...) // Remove the tag
				i--                                       // Adjust the index after removal
			}
		case "(cap,":
			if i != 0 && reg.MatchString(words[i+1]) {
				a := strings.TrimSuffix(words[i+1], ")")

				num, _ := strconv.Atoi(string(a))
				if num > len(words[:i]) {
					num = len(words[:i])
				}
				for b := 1; b <= num; b++ {
					words[i-b] = titleCase(words[i-b])
				}
				words = append(words[:i], words[i+2:]...) // He removes "(cap,"
				i--
			}

		case "(low,":
			if i != 0 && reg.MatchString(words[i+1]) {
				a := strings.TrimSuffix(words[i+1], ")")

				num, _ := strconv.Atoi(string(a))
				if num > len(words[:i]) {
					num = len(words[:i])
				}
				for b := 1; b <= num; b++ {
					words[i-b] = strings.ToLower(words[i-b])
				}
				words = append(words[:i], words[i+2:]...) // He removes "(cap,"
				i--
			}
		case "(up,":
			if i != 0 && reg.MatchString(words[i+1]) {
				a := strings.TrimSuffix(words[i+1], ")")

				num, _ := strconv.Atoi(string(a))
				if num > len(words[:i]) {
					num = len(words[:i])
				}
				for b := 1; b <= num; b++ {
					words[i-b] = strings.ToUpper(words[i-b])
				}
				words = append(words[:i], words[i+2:]...) // He removes "(cap,"
				i--
			}

		}
	}

	// Join the processed words back into a single string
	// resultt := strings.Join(words, " ")
	resultt := ""
	for i, word := range words {
		if i != len(words)-1 {
			resultt = resultt + word + " "
		} else {
			resultt = resultt + word
		}
	}

	// Write the result to the output file
	files, err := os.Create(result)
	if err != nil {
		panic(err)
	}
	defer files.Close()

	if _, err := io.WriteString(files, resultt); err != nil {
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
