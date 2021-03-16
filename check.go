package main

import (
	"go/ast"
	"go/token"
	"strings"
)

var tag = "! "

type routeInfo struct {
	comment      string
	name         string
	structInfo   *ast.GenDecl
	functionList []*functionInfo
	routeList    []*routeInfo
}

type functionInfo struct {
	comment string
	name    string
	recv    string
	astInfo *ast.FuncDecl
}

var rInfo routeInfo

func checkStruct(n ast.Node, comments []*ast.CommentGroup) (info *routeInfo, ok bool) {
	info = &routeInfo{}
	if spec, ok := n.(*ast.GenDecl); ok {
		if spec.Tok == token.TYPE {
			if len(spec.Specs) == 1 {
				sps := spec.Specs[0]
				if m, ok := sps.(*ast.TypeSpec); ok {
					if m.Name != nil {
						info = &routeInfo{
							name:       m.Name.Name,
							structInfo: spec,
						}
					}
				}
			}
			for _, comment := range comments {
				if commentCheck(comment) {
					info.comment = comment.Text()
					return info, true
				}
			}
		}

	}
	return
}

func checkFunc(n ast.Node, comments []*ast.CommentGroup) (info *functionInfo, ok bool) {
	info = &functionInfo{}
	if spec, ok := n.(*ast.FuncDecl); ok {
		if star, ok := spec.Recv.List[0].Type.(*ast.StarExpr); ok {
			if ident, ok := (star.X).(*ast.Ident); ok {
				if spec.Name != nil {
					info = &functionInfo{
						recv:    ident.Name,
						name:    spec.Name.Name,
						astInfo: spec,
					}
				}
			}
		}
		//paramsCheck
		if spec.Type.Params.NumFields() != 1 {
			return nil, false
		}
		li := spec.Type.Params.List[0]
		if p, ok := li.Type.(*ast.StarExpr); ok {
			if sel, ok := p.X.(*ast.SelectorExpr); ok {
				if f, ok := sel.X.(*ast.Ident); ok && f.Name == "gin" {
					if sel.Sel.Name == "Context" {
						for _, v := range comments {
							if commentCheck(v) {
								info.comment = v.Text()
								return info, true
							}
						}
					}
				}
			}
		}

	}
	return
}

func commentCheck(c *ast.CommentGroup) bool {
	return strings.HasPrefix(c.Text(), tag)
}
