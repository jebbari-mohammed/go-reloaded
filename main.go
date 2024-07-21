package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid input")
		return
	}
	sample := os.Args[1]
	result := os.Args[2]
	if result == "main.go" {
		fmt.Println("haha nice try")
		return
	}
	if os.Args[2][len(result)-3:len(result)] != "txt" {
		fmt.Println("invalid input: the file should be .txt")
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
						}
					}
					words = append(words[:i], words[i+1:]...) // Remove the tag
					i--                                       // Adjust the index after removal
				}
			case "(cap,":
				if i == 0  {
					if i+2 < len(words) && reg.MatchString(words[i+1])  {
					words = words[i+2:]
					i--
				}
				}
				if i != 0 && (reg.MatchString(words[i+1]) || regPlus.MatchString(words[i+1])) {
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
				if  i == 0 {
					if i+2 < len(words) && reg.MatchString(words[i+1]) {
					words = words[i+2:]
					i--
				}
				}
				if i != 0 && (reg.MatchString(words[i+1]) || regPlus.MatchString(words[i+1])) {
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
				if  i == 0  {
					if i+2 < len(words) && reg.MatchString(words[i+1]) {
						if regMoins.MatchString(words[i+1]) {
							words = words[i+2:]
							i--
						}
					words = words[i+2:]
					i--
				}
				}
				if i != 0 && (reg.MatchString(words[i+1]) || regPlus.MatchString(words[i+1])) {
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
		Lines := strings.Join(words, " ")
		allLines = append(allLines, Lines)
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
	// Convert the first character to uppercase and the rest to lowercase
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
