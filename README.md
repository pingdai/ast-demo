# ast-demo
基于Golang AST所进行的部分测试用例

## Example

```go
package ast_demo

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"testing"
)

func TestImports(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "ast_traversal.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, i := range f.Imports {
		t.Logf("import:	%s", i.Path.Value)
	}

	return
}

func TestComments(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "ast_traversal.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, i := range f.Comments {
		t.Logf("comment:	%s", i.Text())
	}

	return
}

func TestFunctions(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "ast_traversal.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, i := range f.Decls {
		fn, ok := i.(*ast.FuncDecl)
		if !ok {
			continue
		}
		t.Logf("function:	%s", fn.Name.Name)
	}

	return
}

func TestInspect(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "ast_traversal.go", nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
		return
	}

	ast.Inspect(f, func(n ast.Node) bool {
		// Find Return Statements
		ret, ok := n.(*ast.ReturnStmt)
		if ok {
			t.Logf("return statement found on line %d:\n\t", fset.Position(ret.Pos()).Line)
			printer.Fprint(os.Stdout, fset, ret)
			return true
		}

		// Find Functions
		fn, ok := n.(*ast.FuncDecl)
		if ok {
			var exported string
			if fn.Name.IsExported() {
				exported = "exported "
			}
			t.Logf("%sfunction declaration found on line %d: %s", exported, fset.Position(fn.Pos()).Line, fn.Name.Name)
			return true
		}

		return true
	})

	return
}

```
