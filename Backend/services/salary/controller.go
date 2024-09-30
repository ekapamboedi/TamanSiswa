package salary

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"TamanSiswa/helper"
	"TamanSiswa/middleware"
	"TamanSiswa/services/salary/request"
	"TamanSiswa/services/salary/response"
)

func SalaryRoute(r *gin.RouterGroup) {
	r.GET("", GetAll)
	r.POST("", Create)
	r.GET("/:id", GetOne)
	r.PUT("/:id", Update)
}

func GetAll(ctx *gin.Context) {
	var res response.ResponseList
	var err error
	identity := helper.GetIdentity(ctx)

	data, err := GetAllRecord(identity)
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
		res.Message = "Error occured"
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

	res.Status = "OK"
	res.Message = "Success"
	res.Set(*data)

	ctx.IndentedJSON(200, res)
}

func GetOne(ctx *gin.Context) {
	var res response.ResponseDetail
	var err error
	identity := helper.GetIdentity(ctx)

	id := ctx.Param("id")

	data, err := GetWithId(identity, id)
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
		res.Message = "Error occured"
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

	res.Status = "OK"
	res.Message = "Success"
	res.Set(*data)

	ctx.IndentedJSON(200, res)
}

func Update(ctx *gin.Context) {
	var req request.Request
	var res response.Response
	var err error
	identity := helper.GetIdentity(ctx)

	id := ctx.Param("id")

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err.Error())
		res.Status = "Bad Request"
		res.Message = "Error occured"
		ctx.AbortWithStatusJSON(400, res)
		go middleware.LogError(ctx.Request.URL.Path, ctx.Request.Method, err.Error(), identity)
		return
	}

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
		res.Message = "Error occured"
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
