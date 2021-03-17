package main

import (
	"fmt"
	"go/ast"
	"log"
	"strings"
)

type routeInfo struct {
	route        string
	parent       string
	c            comment
	name         string
	structInfo   *ast.GenDecl
	functionList functionList
	routeList    routeList
}

type functionInfo struct {
	route      string
	method     []string
	c          comment
	middleware bool
	name       string
	recv       string
	astInfo    *ast.FuncDecl
}

var errNotFound = func(t, v string) error {
	return fmt.Errorf("type:%v not found %v", t, v)
}

var errDeplicateRoute = func(s1, s2 string) error {
	return fmt.Errorf("deplicate route in diffrent struct:%v %v", s1, s2)
}

type functionList []*functionInfo

func (f functionList) String() string {
	fL := []string{}
	for _, v := range f {
		fL = append(fL, fmt.Sprintf("  %v", v))
	}
	return strings.Join(fL, "\n")
}

func (f functionList) splitByRecv(recv string) (include, exclude functionList) {
	for _, v := range f {
		if v.recv == recv {
			include = append(include, v)
		} else {
			exclude = append(exclude, v)
		}
	}
	return
}

func (r routeList) buildTree() (tree routeList) {
	m := make(map[string]routeList)
	for i := range r {
		m[r[i].route] = append(m[r[i].route], r[i])
	}
	for _, v := range m {
		for i := range v {
			if v[i].parent == "" {
				tree = append(tree, v[i])
			} else {
				parents, ok := m[v[i].parent]
				if !ok {
					log.Printf("warn: missing parent:%v", v[i].parent)
				}
				for i := range parents {
					parents[i].routeList = append(parents[i].routeList, v[i])
				}
			}
		}
	}
	return
}

type routeList []*routeInfo

func (r routeList) String() string {
	rL := []string{}
	for _, v := range r {
		rL = append(rL, fmt.Sprint(v))
	}
	return strings.Join(rL, "\n")
}

func (r routeList) findByRouteName(name string) (*routeInfo, error) {
	for _, v := range r {
		if v.route == name {
			return v, nil
		}
	}
	return nil, errNotFound(route, name)
}

func (r routeList) findByParent(parent string) []*routeInfo {
	var rl []*routeInfo
	for _, v := range r {
		if v.parent == parent {
			rl = append(rl, v)
		}
	}
	return rl
}

func (r routeInfo) String() string {
	return fmt.Sprintf("struct:%v route:%v parent:%v function:\n   %v\n children:\n   %v", r.name, r.route, r.parent, r.functionList, r.routeList)
}

func (f functionInfo) String() string {
	return fmt.Sprintf("recv:%v route:%v method:%v name:%v middleware:%v", f.recv, f.route, f.method, f.name, f.middleware)
}
