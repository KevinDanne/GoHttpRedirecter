package redirection

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Redirection struct {
	From string
	To   string
}

func GetRedirections(fileName string) ([]Redirection, error) {
	file, err := os.Open("routes/" + fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var redirections []Redirection

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=>", 2)
		if len(parts) != 2 {
			return nil, errors.New(fmt.Sprintf("Invalid line '%s'", line))
		}
		redirection := Redirection{
			From: strings.TrimSpace(parts[0]),
			To:   strings.TrimSpace(parts[1]),
		}
		redirections = append(redirections, redirection)
	}

	return redirections, nil
}
