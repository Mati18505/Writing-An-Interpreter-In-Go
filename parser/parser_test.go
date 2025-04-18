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
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

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