package gingen

import (
	"fmt"
	"go/ast"
	"log"
	"strings"
)

type RouteInfo struct {
	Route        string
	Parent       string
	C            comment
	Name         string
	structInfo   *ast.GenDecl
	FunctionList FunctionList
	RouteList    RouteList
}

type FunctionInfo struct {
	Route      string
	Method     []string
	C          comment
	Middleware bool
	Name       string
	Recv       string
	astInfo    *ast.FuncDecl
}

var errNotFound = func(t, v string) error {
	return fmt.Errorf("type:%v not found %v", t, v)
}

var errDeplicateRoute = func(s1, s2 string) error {
	return fmt.Errorf("deplicate route in diffrent struct:%v %v", s1, s2)
}

type FunctionList []*FunctionInfo

func (f FunctionList) String() string {
	fL := []string{}
	for _, v := range f {
		fL = append(fL, fmt.Sprintf("%v", v))
	}
	return strings.Join(fL, "\n   ")
}

func (f FunctionList) splitByRecv(recv string) (include, exclude FunctionList) {
	for _, v := range f {
		if v.Recv == recv {
			include = append(include, v)
		} else {
			exclude = append(exclude, v)
		}
	}
	return
}

func (r RouteList) BuildTree() (tree RouteList) {
	m := make(map[string]RouteList)
	for i := range r {
		m[r[i].Route] = append(m[r[i].Route], r[i])
	}
	for _, v := range m {
		for i := range v {
			if v[i].Parent == "" {
				tree = append(tree, v[i])
			} else {
				parents, ok := m[v[i].Parent]
				if !ok {
					log.Printf("warn: missing parent:%v", v[i].Parent)
				}
				for i := range parents {
					parents[i].RouteList = append(parents[i].RouteList, v[i])
				}
			}
		}
	}
	return
}

type RouteList []*RouteInfo

func (r RouteList) String() string {
	rL := []string{}
	for _, v := range r {
		rL = append(rL, fmt.Sprint(v))
	}
	return strings.Join(rL, "    \n")
}

func (r RouteList) findByRouteName(name string) (*RouteInfo, error) {
	for _, v := range r {
		if v.Route == name {
			return v, nil
		}
	}
	return nil, errNotFound(route, name)
}

func (r RouteList) findByParent(parent string) []*RouteInfo {
	var rl []*RouteInfo
	for _, v := range r {
		if v.Parent == parent {
			rl = append(rl, v)
		}
	}
	return rl
}

func (r RouteInfo) String() string {
	return fmt.Sprintf("struct:%v route:%v parent:%v function:\n   %v\n children:\n   %v", r.Name, r.Route, r.Parent, r.FunctionList, r.RouteList)
}

func (f FunctionInfo) String() string {
	return fmt.Sprintf("recv:%v route:%v method:%v name:%v middleware:%v", f.Recv, f.Route, f.Method, f.Name, f.Middleware)
}
