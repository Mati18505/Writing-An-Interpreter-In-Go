package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLetSatements(t *testing.T) {
	input := `
	let x 5;
	let = 10;
	let 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	require.NotNil(t, program)
	require.Equal(t, len(program.Statements), 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetSatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetSatement(t *testing.T, s ast.Statement, name string) bool {
	err := assert.Equal(t, s.TokenLiteral(), "let")
	if err != true {
		return err
	}

	letStmt, ok := s.(*ast.LetStatement)

	err = assert.True(t, ok)
	if err != true {
		return err
	}
	
	err = assert.Equal(t, letStmt.Name.Value, name)
	if err != true {
		return err
	}

	err = assert.Equal(t, letStmt.Name.TokenLiteral(), name)
	if err != true {
		return err
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}