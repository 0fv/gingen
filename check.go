package gingen

import (
	"go/ast"
	"go/token"
)

func checkStruct(n ast.Node, comments []*ast.CommentGroup) (info *RouteInfo, ok bool) {
	// info = &routeInfo{}
	if spec, ok := n.(*ast.GenDecl); ok {
		if spec.Tok == token.TYPE {
			if len(spec.Specs) == 1 {
				sps := spec.Specs[0]
				if m, ok := sps.(*ast.TypeSpec); ok {
					if m.Name != nil {
						info = &RouteInfo{
							Name:       m.Name.Name,
							structInfo: spec,
						}
					}
				}
			}
			for _, v := range comments {
				if commentCheck(v) {
					info.C = comment(v.Text())
					info.Route, info.Parent = info.C.checkGroup(info.Name)
					return info, true
				}
			}
		}

	}
	return nil, false
}

func checkFunc(n ast.Node, comments []*ast.CommentGroup) (info *FunctionInfo, err error) {
	info = &FunctionInfo{}
	if spec, ok := n.(*ast.FuncDecl); ok && spec.Recv != nil && len(spec.Recv.List) != 0 {
		if star, ok := spec.Recv.List[0].Type.(*ast.StarExpr); ok {
			if ident, ok := (star.X).(*ast.Ident); ok {
				if spec.Name != nil {
					info = &FunctionInfo{
						Recv:    ident.Name,
						Name:    spec.Name.Name,
						astInfo: spec,
					}
				}
			}
		}
		//paramsCheck
		if spec.Type.Params.NumFields() != 1 {
			return nil, nil
		}
		li := spec.Type.Params.List[0]
		if p, ok := li.Type.(*ast.StarExpr); ok {
			if sel, ok := p.X.(*ast.SelectorExpr); ok {
				if f, ok := sel.X.(*ast.Ident); ok && f.Name == "gin" {
					if sel.Sel.Name == "Context" {
						for _, v := range comments {
							if commentCheck(v) {
								info.C = comment(v.Text())
								info.Route, info.Method, info.Middleware, err = info.C.routeFuncProcess(info.Name)
								return info, err
							}
						}
					}
				}
			}
		}
	}
	return nil, nil
}
