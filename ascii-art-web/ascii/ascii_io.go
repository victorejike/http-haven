package ascii

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type AsciiGenerator struct {
	bannerMap map[string]string
	cache     map[string][]string
}

func NewAsciiGenerator() *AsciiGenerator {
	return &AsciiGenerator{
		bannerMap: map[string]string{
			"standard":   "Banners/standard.txt",
			"shadow":     "Banners/shadow.txt",
			"thinkertoy": "Banners/thinkertoy.txt",
		},
		cache: make(map[string][]string),
	}
}

func (g *AsciiGenerator) GenerateFromReader(reader io.Reader, banner string) (string, error) {
	var input bytes.Buffer
	_, err := io.Copy(&input, reader)
	if err != nil {
		return "", fmt.Errorf("error reading input: %v", err)
	}

	return g.Generate(input.String(), banner)
}

func (g *AsciiGenerator) GenerateFromFile(filePath, banner string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	return g.GenerateFromReader(file, banner)
}

func (g *AsciiGenerator) Generate(input, banner string) (string, error) {
	if !g.IsValidBanner(banner) {
		return "", fmt.Errorf("%w: %s", ErrInvalidBanner, banner)
	}

	lines, err := g.loadBanner(banner)
	if err != nil {
		return "", err
	}

	input = strings.ReplaceAll(input, "\r\n", "\n")
	inputLines := strings.Split(input, "\n")

	var result strings.Builder
	bufWriter := bufio.NewWriter(&result)
	defer bufWriter.Flush()

	for idx, line := range inputLines {
		if line == "" && idx != len(inputLines)-1 {
			bufWriter.WriteString("\n")
			continue
		}

		if line == "" {
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range line {
				index := int(char-32)*9 + i
				if index < len(lines) {
					bufWriter.WriteString(lines[index])
				} else {
					bufWriter.WriteString(strings.Repeat(" ", 8))
				}
			}
			bufWriter.WriteString("\n")
		}
	}

	bufWriter.Flush()
	return result.String(), nil
}

func (g *AsciiGenerator) GenerateStreaming(input string, banner string, writer io.Writer) error {
	if !g.IsValidBanner(banner) {
		return fmt.Errorf("%w: %s", ErrInvalidBanner, banner)
	}

	lines, err := g.loadBanner(banner)
	if err != nil {
		return err
	}

	input = strings.ReplaceAll(input, "\r\n", "\n")
	inputLines := strings.Split(input, "\n")

	bufWriter := bufio.NewWriter(writer)
	defer bufWriter.Flush()

	for idx, line := range inputLines {
		if line == "" && idx != len(inputLines)-1 {
			bufWriter.WriteString("\n")
			continue
		}

		if line == "" {
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range line {
				index := int(char-32)*9 + i
				if index < len(lines) {
					bufWriter.WriteString(lines[index])
				} else {
					bufWriter.WriteString(strings.Repeat(" ", 8))
				}
			}
			bufWriter.WriteString("\n")
		}
	}

	return bufWriter.Flush()
}

func (g *AsciiGenerator) loadBanner(banner string) ([]string, error) {
	if lines, ok := g.cache[banner]; ok {
		return lines, nil
	}

	bannerPath, ok := g.bannerMap[banner]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrInvalidBanner, banner)
	}

	file, err := os.Open(bannerPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrBannerNotFound, bannerPath)
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

	g.cache[banner] = lines
	return lines, nil
}

func (g *AsciiGenerator) IsValidBanner(banner string) bool {
	_, ok := g.bannerMap[banner]
	return ok
}

func (g *AsciiGenerator) GetAvailableBanners() []string {
	keys := make([]string, 0, len(g.bannerMap))
	for k := range g.bannerMap {
		keys = append(keys, k)
	}
	return keys
}

func (g *AsciiGenerator) ClearCache() {
	g.cache = make(map[string][]string)
}

func (g *AsciiGenerator) GetBannerSize(banner string) (int64, error) {
	bannerPath, ok := g.bannerMap[banner]
	if !ok {
		return 0, fmt.Errorf("banner '%s' not found", banner)
	}

	fileInfo, err := os.Stat(bannerPath)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrBannerNotFound, bannerPath)
	}

	return fileInfo.Size(), nil
}

func ReadBannerLines(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading banner: %v", err)
	}
	return lines, nil
}

func ProcessInputReader(reader io.Reader) ([]string, error) {
	var input bytes.Buffer
	_, err := io.Copy(&input, reader)
	if err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	processed := strings.ReplaceAll(input.String(), "\r\n", "\n")
	return strings.Split(processed, "\n"), nil
}
