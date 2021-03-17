package exmaple

import "github.com/gin-gonic/gin"

//! parent=/foo
type bar struct {
}

//! method=POST route=/:test2
func (f *bar) post(ctx *gin.Context) {
	ctx.JSON(200, 212)
}

//! middleware
func (f *bar) ware(ctx *gin.Context) {

}
