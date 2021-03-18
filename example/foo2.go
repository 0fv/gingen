package exmaple

import "github.com/gin-gonic/gin"

//! group=/foo2 parent=/bar
type foo2 struct {
}

//! route=/:test2
func (f *foo2) post(ctx *gin.Context) {
	ctx.JSON(200, 212)
	gin.Default().Use()
}

//! middleware
func (f *foo2) ware(ctx *gin.Context) {

}
