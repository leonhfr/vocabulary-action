package vocabulary

import (
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type Parser struct {
	gm parser.Parser
}

func NewParser() *Parser {
	return &Parser{parser.NewParser(
		parser.WithBlockParsers(
			util.Prioritized(parser.NewATXHeadingParser(), 100),
			util.Prioritized(parser.NewHTMLBlockParser(), 200),
			util.Prioritized(parser.NewParagraphParser(), 300),
		),
	)}
}

func (p *Parser) Parse(input string) []string {
	source := []byte(input)
	var vocabulary []string
	doc := p.gm.Parse(text.NewReader(source))
	_ = ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if paragraph, ok := node.(*ast.Paragraph); ok && entering && node.HasChildren() {
			vocabulary = append(vocabulary, rawText(paragraph, source))
		}
		return ast.WalkContinue, nil
	})
	return vocabulary
}

func rawText(node ast.Node, source []byte) string {
	var lines []string
	for i := 0; i < node.Lines().Len(); i++ {
		line := node.Lines().At(i)
		lines = append(lines, string(line.Value(source)))
	}
	return strings.Join(lines, "")
}
