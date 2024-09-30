package menuaccess

import (
	"fmt"
	"strings"

	"TamanSiswa/helper"
	"TamanSiswa/middleware"
	"TamanSiswa/services/menuAccess/request"
	"TamanSiswa/services/menuAccess/response"
	"github.com/gin-gonic/gin"
)

func MenuRoute(r *gin.RouterGroup) {
	r.POST("", Create)
	r.PUT("/:id", Update)
	r.DELETE("/:id", Delete)
	r.GET("", GetAll)
	r.GET("/:id", GetOne)
}

func GetAll(ctx *gin.Context) {
	var res response.ResponseList
	var err error
	identity := helper.GetIdentity(ctx)

	menuAccessData, positionData, err := GetAllRecord(identity)

	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = err.Error()
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

	res.Status = "OK"
	res.Message = "Success"
	res.Set(*menuAccessData, *positionData)

	ctx.IndentedJSON(200, res)
}

func GetOne(ctx *gin.Context) {
	var res response.ResponseDetail
	var err error
	identity := helper.GetIdentity(ctx)

	id := ctx.Param("id")

	result, err := GetWithId(identity, id)
	if err != nil && strings.Contains(err.Error(), "not found") {
		res.Status = "Not Found"
		res.Message = "Cannot find with specified ID"
		ctx.AbortWithStatusJSON(404, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, "Not Found", identity)
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = err.Error()
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

	res.Status = "OK"
	res.Message = "Success"
	res.Set(*result)

	ctx.IndentedJSON(200, res)
}

func Delete(ctx *gin.Context) {
	var res response.Response
	var err error
	identity := helper.GetIdentity(ctx)

	id := ctx.Param("id")

	err = DeleteOne(identity, id)
	if err != nil && err.Error() == "not found" {
		res.Status = "Not Found"
		res.Message = "Cannot find with specified ID"
		ctx.AbortWithStatusJSON(404, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, "Not Found", identity)
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = err.Error()
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

	res.Status = "OK"
	res.Message = "Success"

	ctx.IndentedJSON(200, res)
}

func Update(ctx *gin.Context) {
	var req request.RequestUpdate
	var res response.Response
	var err error
	identity := helper.GetIdentity(ctx)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = "Error occured"
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}
	id := ctx.Param("id")

	err = UpdateOne(identity, req, id)
	if err != nil && err.Error() == "not found" {
		res.Status = "Not Found"
		res.Message = "Cannot find with specified ID"
		ctx.AbortWithStatusJSON(404, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, "Not Found", identity)
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = err.Error()
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

	res.Status = "OK"
	res.Message = "Success"

	ctx.IndentedJSON(200, res)
}

func Create(ctx *gin.Context) {
	var req request.Request
	var res response.Response
	var err error
	identity := helper.GetIdentity(ctx)

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = "Error occured"
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

	err = CreateOne(identity, req)
	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = err.Error()
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

	res.Status = "Created"
	res.Message = "Success"

	ctx.IndentedJSON(201, res)
}
