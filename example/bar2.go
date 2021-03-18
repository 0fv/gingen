package exmaple

import "github.com/gin-gonic/gin"

//! parent=/foo
type bar2 struct {
}

//! method=POST route=/:test2
func (f *bar2) post(ctx *gin.Context) {
	ctx.JSON(200, 212)
}

//! route=/:test3
func (f *bar2) get(ctx *gin.Context) {
	ctx.JSON(200, 212)
}

//! middleware
func (f *bar2) ware(ctx *gin.Context) {

}
