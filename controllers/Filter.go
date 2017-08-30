package controllers

import (
	"github.com/astaxie/beego/context"
	"fmt"
)

func BeforeExec(ctx *context.Context)  {
	//fmt.Println("握草", ctx.Request.RequestURI)
	fmt.Println("我执行了吗 BeforeExec")
	ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Output.Header("Pragma", "no-cache")
	ctx.Output.Header("Expires","0")
}

func BeforeStatic(ctx *context.Context)  {
	fmt.Println("我执行了吗 BeforeStatic")
	ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Output.Header("Pragma", "no-cache")
	ctx.Output.Header("Expires","0")
}

func BeforeRouter(ctx *context.Context)  {
	fmt.Println("我执行了吗 BeforeRouter")
	ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Output.Header("Pragma", "no-cache")
	ctx.Output.Header("Expires","0")
}