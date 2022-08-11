package application

import (
	"fmt"
	"io"
)

func parseArgs(input string) ([]string, error) {
	p := &parser{
		input:  input,
		cursor: 0,
	}
	return p.parse()
}

type parser struct {
	input  string
	cursor int
}

func (p *parser) parse() ([]string, error) {
	args := make([]string, 0)

	for {
		b, err := p.next()
		if err != nil {
			break
		}
		if isSpace(b) {
			p.cursor++
			continue
		}
		switch {
		case isQuote(b):
			arg, err := p.parseQuoted()
			if err != nil {
				return nil, err
			}
			if arg != "" {
				args = append(args, arg)
			}
		case isDash(b):
			argName := p.parseArgName()
			if argName != "" {
				args = append(args, argName)
			}
		case isLetter(b):
			token := p.parseToken()
			if token != "" {
				args = append(args, token)
			}
		default:
			return nil, fmt.Errorf("invalid tokens")
		}
	}
	return args, nil
}

func (p *parser) next() (byte, error) {
	if p.cursor >= len(p.input) {
		return 0, io.EOF
	}
	b := p.input[p.cursor]
	return b, nil
}

func (p *parser) read() (byte, error) {
	if p.cursor >= len(p.input) {
		return 0, io.EOF
	}
	b := p.input[p.cursor]
	p.cursor++
	return b, nil
}

func (p *parser) parseQuoted() (string, error) {
	q, err := p.read()
	if err != nil {
		return "", err
	}
	val := []byte{}
	ended := false
	for {
		if b, err := p.read(); err != nil || b == q {
			if b == q {
				ended = true
			}
			break
		} else {
			val = append(val, b)
		}
	}
	if !ended {
		return "", fmt.Errorf("quote does not match")
	}
	return string(val), nil
}

func (p *parser) parseArgName() string {
	name := []byte{}
	for {
		if b, err := p.read(); err != nil || b == '=' || isSpace(b) {
			break
		} else {
			name = append(name, b)
		}
	}
	return string(name)
}

func (p *parser) parseToken() string {
	token := []byte{}

	for {
		if b, err := p.read(); err != nil || isSpace(b) {
			break
		} else {
			token = append(token, b)
		}
	}
	return string(token)
}

func isDash(c byte) bool {
	return c == '-'
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t'
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' || c <= '9')
}

func isQuote(c byte) bool {
	return c == '\'' || c == '"'
}
