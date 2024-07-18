package main

import (
	"fmt"
	"io"
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

	file, err := os.ReadFile(sample)
	if err != nil {
		panic(err)
	}
	words := strings.Split(string(file), " ")
	reg := regexp.MustCompile(`^\d+\)$`)
	regponc := regexp.MustCompile(`^[.,?!:;]+$`)
	for i := 0; i <= len(words)-1; i++ {
		switch words[i] {

		case "(hex)":
			if i == 0 {
				words = words[i+1:]
			}
			if i != 0 {
				for j := 1; j <= len(words[:i]); j++ {
					if words[i-j] == "" {
						continue
					} else {
						number, _ := strconv.ParseInt(words[i-j], 16, 64)
						words[i-j] = fmt.Sprint(number)
						break
					}
				}
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(bin)":
			if i == 0 {
				words = words[i+1:]
			}
			if i != 0 {
				for j := 1; j <= len(words[:i]); j++ {
					if words[i-j] == "" {
						continue
					} else {
						number, _ := strconv.ParseInt(words[i-j], 2, 64)
						words[i-j] = fmt.Sprint(number)
						break
					}
				}
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(up)":
			if i == 0 {
				words = words[i+1:]
			}
			if i != 0 {
				for j := 1; j <= len(words[:i]); j++ {
					if words[i-j] == "" {
						continue
					} else {
						words[i-j] = strings.ToUpper(words[i-j])
						break
					}
				}
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(cap)":
			if i == 0 {
				words = words[i+1:]
			}
			if i != 0 {
				for j := 1; j <= len(words[:i]); j++ {
					if words[i-j] == "" {
						continue
					} else {
						if len(words[i-j]) > 1 {
							words[i-j] = strings.ToUpper(string(words[i-j][0])) + strings.ToLower(string(words[i-j][1:]))
							break
						}
						if len(words[i-j]) == 0 {
							words[i-j] = strings.ToUpper(words[i-j])
							break
						}
					}
				}
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(low)":
			if i == 0 {
				words = words[i+1:]
			}
			if i != 0 {
				for j := 1; j <= len(words[:i]); j++ {
					if words[i-j] == "" {
						continue
					} else {
						words[i-j] = strings.ToLower(words[i-j])
						break
					}
				}
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(up,":
			if i == 0 {
				words = words[i+1:]
			}
			if i != 0 {
				if reg.MatchString(words[i+1]) {
					a := strings.TrimSuffix(words[i+1], ")")

					num, _ := strconv.Atoi(string(a))
					if num > len(words[:i]) {
						num = len(words[:i])
					}
					if num <= len(words[:i]) {
						for k := 1; k < num; k++ {
							if words[i-k] == "" || regponc.MatchString(words[i-k]) {
								num++
							}
						}
					}
					for l := 1; l <= num; l++ {
						if words[i-l] != "" {
							words[i-l] = strings.ToUpper(words[i-l])
						}
					}

				}
				words = append(words[:i], words[i+2:]...)
				i--
			}

		case "(low,":
			if i == 0 {
				words = words[i+1:]
			}
			if i != 0 {
				if reg.MatchString(words[i+1]) {
					a := strings.TrimSuffix(words[i+1], ")")

					num, _ := strconv.Atoi(string(a))
					if num > len(words[:i]) {
						num = len(words[:i])
					}
					if num <= len(words[:i]) {
						for p := 1; p < num; p++ {
							if words[i-p] == "" || regponc.MatchString(words[i-p]) {
								num++
							}
						}
					}
					for f := 1; f <= num; f++ {
						if words[i-f] != "" {
							words[i-f] = strings.ToLower(words[i-f])
						}
					}

				}
				words = append(words[:i], words[i+2:]...)
				i--
			}

		case "(cap,":
			if i == 0 {
				words = words[i+1:]
			}
			if i != 0 {
				if reg.MatchString(words[i+1]) {
					a := strings.TrimSuffix(words[i+1], ")")

					num, _ := strconv.Atoi(string(a))
					if num > len(words[:i]) {
						num = len(words[:i])
					}
					if num <= len(words[:i]) {
						for k := 1; k < num; k++ {
							if words[i-k] == "" || regponc.MatchString(words[i-k]) {
								num++
							}
						}
					}
					for l := 1; l <= num; l++ {
						if words[i-l] != "" {
							words[i-l] = strings.ToUpper(string(words[i-l][0])) + strings.ToLower(string(words[i-l][1:]))
						}
					}

				}
				words = append(words[:i], words[i+2:]...)
				i--
			}

		}
	}

	finaltxt := strings.Join(words, " ")
	finalfile, err := os.Create(result)
	if err != nil {
		panic(err)
	}
	if _, err := io.WriteString(finalfile, finaltxt); err != nil {
		panic(err)
	}
}
