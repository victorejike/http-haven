package ascii

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrInvalidBanner  = errors.New("invalid banner")
	ErrBannerNotFound = errors.New("banner not found")
)

var banners = map[string]string{
	"standard":   "Banners/standard.txt",
	"shadow":     "Banners/shadow.txt",
	"thinkertoy": "Banners/thinkertoy.txt",
}

func ValidateInput(input string) bool {
	for _, char := range input {
		if char < 32 || char > 126 {
			if char != 10 && char != 13 {
				return false
			}
		}
	}
	return true
}

func IsValidBanner(banner string) bool {
	_, ok := banners[banner]
	return ok
}

func GetAvailableBanners() []string {
	keys := make([]string, 0, len(banners))
	for k := range banners {
		keys = append(keys, k)
	}
	return keys
}

func GenerateAsciiArt(input, banner string) (string, error) {
	bannerPath, ok := banners[banner]
	if !ok {
		return "", fmt.Errorf("%w: %s", ErrInvalidBanner, banner)
	}

	lines, err := readBannerFile(bannerPath)
	if err != nil {
		return "", err
	}

	input = strings.ReplaceAll(input, "\r\n", "\n")
	inputLines := strings.Split(input, "\n")

	var result strings.Builder

	for idx, line := range inputLines {
		if line == "" && idx != len(inputLines)-1 {
			result.WriteString("\n")
			continue
		}

		if line == "" {
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range line {
				index := int(char-32)*9 + i
				if index < len(lines) {
					result.WriteString(lines[index])
				} else {
					result.WriteString(strings.Repeat(" ", 8))
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

func readBannerFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrBannerNotFound, path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading banner file: %v", err)
	}

	if len(lines) < 95*9 {
		return nil, fmt.Errorf("invalid banner file format: expected at least %d lines, got %d", 95*9, len(lines))
	}

	return lines, nil
}

func fallbackBannerLines() []string {
	lines := make([]string, 0, 95*9)
	for char := ' '; char <= '~'; char++ {
		glyph := fallbackGlyph(char)
		lines = append(lines, glyph[:]...)
		lines = append(lines, "")
	}
	return lines
}

func fallbackGlyph(char rune) [8]string {
	if char == ' ' {
		return [8]string{"        ", "        ", "        ", "        ", "        ", "        ", "        ", "        "}
	}

	label := string(char)
	return [8]string{
		"  ____  ",
		" /    \\ ",
		"|  " + label + "   | ",
		"|      | ",
		"|  " + label + "   | ",
		"|      | ",
		" \\____/ ",
		"        ",
	}
}
