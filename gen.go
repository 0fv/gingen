package gingen

import (
	"errors"
	"io"
	"os"
	"text/template"
)

const suffix = "_route_gen.go"

const rootltpl = `
// Code generated by github.com/0fv/gingen DO NOT EDIT

package {{.pkg}}

import "github.com/gin-gonic/gin"

func initRoute(r *gin.Engine){
	{{- range $tr:=.li}}
	{{$routeName:=print "route_" $tr.Name }}
	{{$routeName}} := {{$tr.Name}}{}
	{{$groupName:=print "group_" $tr.Name}}
	{{$groupName}}:= r.Group("{{$tr.Route}}")
	{{- range $fun:=.FunctionList}}
	{{- if $fun.Middleware}}
	{{$groupName}}.Use({{$routeName}}.{{$fun.Name -}})
	{{else}}
	{{- range $me:=$fun.Method}}
	{{$groupName}}.{{$me}}("{{$fun.Route}}",{{$routeName}}.{{$fun.Name -}})
	{{end -}}
	{{end -}}
	{{end -}}
	{{- range $cr:=.RouteList}}
	init_{{$cr.Name}}({{$groupName -}})
	{{end -}}
	{{end}}
}
`
const grouptpl = `

func init_{{.Name}}(r *gin.RouterGroup){
	{{$routeName:=print "route_" $.Name }}
	{{$routeName}} := {{$.Name}}{}
	{{$groupName:=print "group_" $.Name}}
	{{$groupName}}:= r.Group("{{$.Route}}")
	{{- range $fun:=.FunctionList}}
	{{- if $fun.Middleware}}
	{{$groupName}}.Use({{$routeName}}.{{$fun.Name -}})
	{{else}}
	{{- range $me:=$fun.Method}}
	{{$groupName}}.{{$me}}("{{$fun.Route}}",{{$routeName}}.{{$fun.Name -}})
	{{end -}}
	{{end -}}
	{{end -}}
	{{- range $cr:=.RouteList}}
	init_{{$cr.Name}}({{$groupName -}})
	{{end }}
}

`

//GenFile ...
func GenFile(rs RouteList) error {
	fileName := pkgName + suffix
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	if err = genAll(rs, f); err != nil {
		return err
	}
	return nil
}

func genAll(rs RouteList, wirter io.Writer) error {
	if err := genRoot(rs, wirter); err != nil {
		return err
	}
	for _, v := range rs {
		if err := genSub(v, wirter); err != nil {
			return err
		}
	}
	return nil
}

func genRoot(rs RouteList, wirter io.Writer) error {
	data := map[string]interface{}{
		"pkg": os.Getenv("GOPACKAGE"),
		"li":  rs,
	}
	t, err := template.New("").Parse(rootltpl)
	if err != nil {
		return errors.New("template init err:" + err.Error())
	}
	return t.Execute(wirter, data)
}

func genSub(r *RouteInfo, wirter io.Writer) error {
	t, err := template.New("").Parse(grouptpl)
	if err != nil {
		return errors.New("template init err:" + err.Error())
	}
	if err = t.Execute(wirter, r); err != nil {
		return err
	}
	for _, v := range r.RouteList {
		if err = genSub(v, wirter); err != nil {
			return err
		}
	}
	return nil
}
