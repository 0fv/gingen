# GinGen

gin route file generator

## Install

```bash
go get -u github.com/0fv/gingen/cmd/gingen
```

## Usage Example

```go
//! group=/foo2 parent=/bar               <== create group /foo2 under gourp /bar(if group not defined,group name=struct name)
type foo2 struct {
}

//! route=/:test1 method=GET,POST         <== create route /:test1 under group /foo2 ,method:get or post          
func (f *foo2) find(ctx *gin.Context) {                      
	ctx.JSON(200, 212)
}

//! route=/:test2                         <== create route /:test2 under group /foo2, method:post(method not defined,use function name,ignore case)          
func (f *foo2) post(ctx *gin.Context) {
	ctx.JSON(200, 212)
}

//! middleware                            <== create middleware under group /foo2
func (f *foo2) ware(ctx *gin.Context) {

}

```

* [example](https://github.com/0fv/gingen/tree/master/example)

