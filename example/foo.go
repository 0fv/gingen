package exmaple

import "github.com/gin-gonic/gin"

//! group=/foo
type foo struct {
}

//! route=/:test2
func (f *foo) post(ctx *gin.Context) {
	ctx.JSON(200, 212)
}

//! 
func (f *foo) get(ctx *gin.Context) {
	ctx.JSON(200, 212)
}

//! route=/update
func (f *foo) put(ctx *gin.Context) {
	ctx.JSON(200, 212)
}


//! middleware
func (f *foo) ware(ctx *gin.Context) {

}
