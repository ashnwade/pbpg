// generated by pbpg, do not modify

package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

// type String string
// type Code string
// type Literal string
// type Lex string
// type Repetition *GOR
// type Option *GOR
// type Group *GOR
// type Term *Term
// type Alternative *Alternative
// type Expression *Expression
// type CodeBlock string
// type Error string
// type Action string
//  The top level production is the initial state to attempt to reduce.
// Program = [ Comment { Comment } ] [ Header ] { Types } Line { Line }
func (p *pbpgParser) stateProgram() (err error) {
	// option
	p = p.predict()
	err = p.stateComment()
	if err == nil {
		// repetition
		for {
			p = p.predict()
			err = p.stateComment()
			if err != nil {
				p = p.backtrack()
				p.lastErr = err
				err = nil
				break
			} else {
				p = p.accept()
			}
		}
	}
	if err != nil {
		p = p.backtrack()
		p.lastErr = err
		err = nil
	} else {
		p = p.accept()
	}
	if err == nil {
		// option
		p = p.predict()
		err = p.stateHeader()
		if err != nil {
			p = p.backtrack()
			p.lastErr = err
			err = nil
		} else {
			p = p.accept()
		}
		if err == nil {
			// repetition
			for {
				p = p.predict()
				err = p.stateTypes()
				if err != nil {
					p = p.backtrack()
					p.lastErr = err
					err = nil
					break
				} else {
					p = p.accept()
				}
			}
			if err == nil {
				err = p.stateLine()
				if err == nil {
					// repetition
					for {
						p = p.predict()
						err = p.stateLine()
						if err != nil {
							p = p.backtrack()
							p.lastErr = err
							err = nil
							break
						} else {
							p = p.accept()
						}
					}
				}
			}
		}
	}
	return err
}

// Header = CodeBlock
func (p *pbpgParser) stateHeader() (err error) {
	err = p.stateCodeBlock()
	if err == nil {
		p.Data.actionHeader(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionHeader(whitespace bool, lit, lex string) {
	p.out.WriteString(doNotModify)
	p.out.WriteString(v1)
}

// Types = "type" String String
func (p *pbpgParser) stateTypes() (err error) {
	err = p.literal("type")
	if err == nil {
		err = p.stateString()
		if err == nil {
			err = p.stateString()
		}
	}
	if err == nil {
		p.Data.actionTypes(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionTypes(whitespace bool, lit, lex string) {
	if _, ok := p.typeMap[v2]; ok {
		log.Fatalf("type %v redeclared", v2)
	}
	p.typeMap[v2] = v3

}

// Line = Comment | Production
func (p *pbpgParser) stateLine() (err error) {
	err = p.stateComment()
	if err != nil {
		err = p.stateProduction()
	}
	return err
}

// Production = Name "=" [ Expression ] "." [ Action ] [ Error ]
func (p *pbpgParser) stateProduction() (err error) {
	err = p.stateName()
	if err == nil {
		err = p.literal("=")
		if err == nil {
			// option
			p = p.predict()
			err = p.stateExpression()
			if err != nil {
				p = p.backtrack()
				p.lastErr = err
				err = nil
			} else {
				p = p.accept()
			}
			if err == nil {
				err = p.literal(".")
				if err == nil {
					// option
					p = p.predict()
					err = p.stateAction()
					if err != nil {
						p = p.backtrack()
						p.lastErr = err
						err = nil
					} else {
						p = p.accept()
					}
					if err == nil {
						// option
						p = p.predict()
						err = p.stateError()
						if err != nil {
							p = p.backtrack()
							p.lastErr = err
							err = nil
						} else {
							p = p.accept()
						}
					}
				}
			}
		}
	}
	if err == nil {
		p.Data.actionProduction(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionProduction(whitespace bool, lit, lex string) {
	if p.stateMap[v1] != nil {
		log.Fatalf("%v redeclared", v1)
	}
	p.stateMap[v1] = v3

	a, err := p.patchTypes(v3, v5)
	if err != nil {
		log.Fatal(err)
	}
	e, err := p.patchTypes(v3, v6)
	if err != nil {
		log.Fatal(err)
	}

	p.emitState(v1, v3, a, e)

	if p.entryPoint == "" {
		p.entryPoint = v1
	}

}

// Action = "Action" CodeBlock
func (p *pbpgParser) stateAction() (err error) {
	err = p.literal("Action")
	if err == nil {
		err = p.stateCodeBlock()
	}
	if err == nil {
		p.Data.actionAction(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionAction(whitespace bool, lit, lex string) {
	return v2
}

// Error = "Error" CodeBlock
func (p *pbpgParser) stateError() (err error) {
	err = p.literal("Error")
	if err == nil {
		err = p.stateCodeBlock()
	}
	if err == nil {
		p.Data.actionError(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionError(whitespace bool, lit, lex string) {
	return v2
}

// CodeBlock = "{" Code "}"
func (p *pbpgParser) stateCodeBlock() (err error) {
	err = p.literal("{")
	if err == nil {
		err = p.stateCode()
		if err == nil {
			err = p.literal("}")
		}
	}
	if err == nil {
		p.Data.actionCodeBlock(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionCodeBlock(whitespace bool, lit, lex string) {
	return v2
}

// Expression = Alternative { "|" Alternative }
func (p *pbpgParser) stateExpression() (err error) {
	err = p.stateAlternative()
	if err == nil {
		// repetition
		for {
			p = p.predict()
			err = p.literal("|")
			if err == nil {
				err = p.stateAlternative()
			}
			if err != nil {
				p = p.backtrack()
				p.lastErr = err
				err = nil
				break
			} else {
				p = p.accept()
			}
		}
	}
	if err == nil {
		p.Data.actionExpression(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionExpression(whitespace bool, lit, lex string) {
	return &Expression{alternatives: append([]*Alternative, v1, v3...)}
}

// Alternative = Term { Term }
func (p *pbpgParser) stateAlternative() (err error) {
	err = p.stateTerm()
	if err == nil {
		// repetition
		for {
			p = p.predict()
			err = p.stateTerm()
			if err != nil {
				p = p.backtrack()
				p.lastErr = err
				err = nil
				break
			} else {
				p = p.accept()
			}
		}
	}
	if err == nil {
		p.Data.actionAlternative(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionAlternative(whitespace bool, lit, lex string) {
	return &Alternative{terms: append([]*Term, v1, v2...)}
}

// Term = Lex | Name | Literal | Group | Option | Repetition
func (p *pbpgParser) stateTerm() (err error) {
	err = p.stateLex()
	if err != nil {
		err = p.stateName()
		if err != nil {
			err = p.stateLiteral()
			if err != nil {
				err = p.stateGroup()
				if err != nil {
					err = p.stateOption()
					if err != nil {
						err = p.stateRepetition()
					}
				}
			}
		}
	}
	if err == nil {
		p.Data.actionTerm(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionTerm(whitespace bool, lit, lex string) {
	t := &Term{
		option: vP1,
	}
	switch t.option {
	case 1:
		t.Lex = v1
	case 2:
		t.Name = v1
	case 3:
		t.Literal = v1
	case 4, 5, 6:
		t.GOR = v1
	}
	return t

}

// Group = "(" Expression ")"
func (p *pbpgParser) stateGroup() (err error) {
	err = p.literal("(")
	if err == nil {
		err = p.stateExpression()
		if err == nil {
			err = p.literal(")")
		}
	}
	if err == nil {
		p.Data.actionGroup(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionGroup(whitespace bool, lit, lex string) {
	return &GOR{option: TYPE_GROUP, expression: v2}
}

// Option = "[" Expression "]"
func (p *pbpgParser) stateOption() (err error) {
	err = p.literal("[")
	if err == nil {
		err = p.stateExpression()
		if err == nil {
			err = p.literal("]")
		}
	}
	if err == nil {
		p.Data.actionOption(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionOption(whitespace bool, lit, lex string) {
	return &GOR{option: TYPE_OPTION, expression: v2}
}

// Repetition = "{" Expression "}"
func (p *pbpgParser) stateRepetition() (err error) {
	err = p.literal("{")
	if err == nil {
		err = p.stateExpression()
		if err == nil {
			err = p.literal("}")
		}
	}
	if err == nil {
		p.Data.actionRepetition(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionRepetition(whitespace bool, lit, lex string) {
	return &GOR{option: TYPE_REPETITION, expression: v2}
}

// Lex = "lex" "(" String ")"
func (p *pbpgParser) stateLex() (err error) {
	err = p.literal("lex")
	if err == nil {
		err = p.literal("(")
		if err == nil {
			err = p.stateString()
			if err == nil {
				err = p.literal(")")
			}
		}
	}
	if err == nil {
		p.Data.actionLex(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionLex(whitespace bool, lit, lex string) {
	return v3
}

// Literal = """ String """
func (p *pbpgParser) stateLiteral() (err error) {
	err = p.literal("\"")
	if err == nil {
		err = p.stateString()
		if err == nil {
			err = p.literal("\"")
		}
	}
	if err == nil {
		p.Data.actionLiteral(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionLiteral(whitespace bool, lit, lex string) {
	return v2
}

// Name = String
func (p *pbpgParser) stateName() (err error) {
	err = p.stateString()
	if err == nil {
		p.Data.actionName(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionName(whitespace bool, lit, lex string) {
	return v1
}

//  Lexer directives.
// Code = code
func (p *pbpgParser) stateCode() (err error) {
	{
		n, lexeme, lerr := p.Data.lexcode(p.input[p.pos:])
		p.pos += n
		if lerr != nil {
			err = fmt.Errorf("%v: %w", p.position(), lerr)
		} else {
			err = nil
			p.lexeme = lexeme
		}
	}
	if err == nil {
		p.Data.actionCode(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionCode(whitespace bool, lit, lex string) {
	return v1
}

// String = string
func (p *pbpgParser) stateString() (err error) {
	{
		n, lexeme, lerr := p.Data.lexstring(p.input[p.pos:])
		p.pos += n
		if lerr != nil {
			err = fmt.Errorf("%v: %w", p.position(), lerr)
		} else {
			err = nil
			p.lexeme = lexeme
		}
	}
	if err == nil {
		p.Data.actionString(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionString(whitespace bool, lit, lex string) {
	return v1
}

// Comment = "#" comment
func (p *pbpgParser) stateComment() (err error) {
	err = p.literal("#")
	if err == nil {
		{
			n, lexeme, lerr := p.Data.lexcomment(p.input[p.pos:])
			p.pos += n
			if lerr != nil {
				err = fmt.Errorf("%v: %w", p.position(), lerr)
			} else {
				err = nil
				p.lexeme = lexeme
			}
		}
	}
	if err == nil {
		p.Data.actionComment(p.lastWhitespace, p.lastLiteral, p.lexeme)
	}

	return err
}

func (p *pbpgData) actionComment(whitespace bool, lit, lex string) {
	p.out.WriteString("// " + v2)
}

func Parsepbpg(input string) (*pbpgParser, error) {
	p := newpbpgParser(input)

	err := p.stateProgram()
	if err == nil {
		if strings.TrimSpace(p.input[p.pos:]) != "" {
			return p, p.lastErr
		}
	}
	return p, err
}

type pbpgParser struct {
	input          string
	pos            int
	lineOffsets    []int
	lexeme         string
	Data           *pbpgData
	lastErr        error
	lastLiteral    string
	lastWhitespace bool

	predictStack []*pbpgParser
}

func newpbpgParser(input string) *pbpgParser {
	return &pbpgParser{
		input:       input,
		lineOffsets: pbpgGenerateLineOffsets(input),
		Data:        &pbpgData{},
	}
}

func pbpgGenerateLineOffsets(input string) []int {
	var ret []int

	lines := strings.Split(input, "\n")

	offset := 0
	for _, v := range lines {
		ret = append(ret, len(v)+1+offset)
		offset += len(v) + 1
	}
	return ret
}

func (p *pbpgParser) position() string {
	for i, v := range p.lineOffsets {
		if p.pos < v {
			return fmt.Sprintf("line %v", i)
		}
	}
	return fmt.Sprintln("impossible line reached", p.pos)
}

func (p *pbpgParser) literal(want string) error {
	count := 0
	for r, s := utf8.DecodeRuneInString(p.input[p.pos+count:]); s > 0 && unicode.IsSpace(r); r, s = utf8.DecodeRuneInString(p.input[p.pos+count:]) {
		count += s
	}

	if strings.HasPrefix(p.input[p.pos+count:], want) {
		p.pos += count + len(want)
		p.lastLiteral = want
		p.lastWhitespace = count > 0
		return nil
	}

	return fmt.Errorf("%v: expected \"%v\"", p.position(), want)
}

func (p *pbpgParser) predict() *pbpgParser {
	p.predictStack = append(p.predictStack, p)
	return &pbpgParser{
		input:        p.input,
		pos:          p.pos,
		lineOffsets:  p.lineOffsets,
		lexeme:       p.lexeme,
		Data:         p.Data.fork(),
		predictStack: p.predictStack,
		lastErr:      p.lastErr,
	}
}

func (p *pbpgParser) backtrack() *pbpgParser {
	pp := p.predictStack[len(p.predictStack)-1]
	pp.predictStack = pp.predictStack[:len(pp.predictStack)-1]
	pp.lastErr = p.lastErr
	return pp
}

func (p *pbpgParser) accept() *pbpgParser {
	pp := p.backtrack()
	pp.pos = p.pos
	pp.lexeme = p.lexeme
	pp.Data.merge(p.Data)
	return pp
}
