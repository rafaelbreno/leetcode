package calculator

import (
	"log"
	"strings"
)

type Parser struct {
	RawStr   string
	Result   int
	Keys     []Key
	MaxDepth int
}

type Key struct {
	Rune string
	T    string
}

var Tokens map[string]string
var Rules map[string][]string
var TokenRules interface{}

func init() {
	Tokens = make(map[string]string, 6)
	Tokens["+"] = "ADD"
	Tokens["-"] = "ADD"
	Tokens["*"] = "MUL"
	Tokens["/"] = "MUL"
	Tokens["("] = "LPAR"
	Tokens[")"] = "RPAR"

	Rules = make(map[string][]string, 4)
	Rules["add"] = []string{"mul ADD add", "mul"}
	Rules["mul"] = []string{"atom MUL mul", "atom"}
	Rules["atom"] = []string{"NUM", "LPAR add RPAR", "neg"}
	Rules["neg"] = []string{"ADD atom"}

	TokenRules = matchTokenRules()
}
func Parse(str string) int {
	p := Parser{
		RawStr: str,
	}

	p.
		cleanWhitespaces().
		setParts()

	log.Println("%v\n", TokenRules)

	log.Printf("%v\n", p.Keys)

	return p.Result
}

func (p *Parser) cleanWhitespaces() *Parser {
	p.RawStr = strings.ReplaceAll(p.RawStr, " ", "")
	return p
}

func (p *Parser) setParts() *Parser {
	for _, val := range p.RawStr {
		key := Key{}

		if k, found := Tokens[string(val)]; found {
			key = Key{
				Rune: string(val),
				T:    k,
			}
		} else {
			key = Key{
				Rune: string(val),
				T:    "NUM",
			}
		}
		p.Keys = append(p.Keys, key)
	}

	return p
}

func matchTokenRules() interface{} {
	var a interface{}
	a = "qaa"

	return a
}
