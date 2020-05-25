package main

var HttpTemplate = `package {{PACKAGE}} 

import (
	"net/http"

	"github.com/goees/yago"
	"github.com/goees/yago/base/basehttp"
)

type {{NAME}}Http struct {
	basehttp.BaseHttp
}

func init() {
	h := new({{NAME}}Http)
	yago.AddHttpRouter("", http.MethodPost, h.ListAction, h)
}

func (h *{{NAME}}Http) ListAction(c *yago.Ctx) {
	return
}
`

var RpcTemplate = `package {{PACKAGE}} 

import (
	"context"
	"log"

	"github.com/goees/yago"

	pb "github.com/goees/yago/example/app/modules/home/homerpc/homepb"
)

type {{NAME}}Rpc struct {
}

func init() {
	h := new({{NAME}}Rpc)
	pb.RegisterHomeServer(yago.RpcServer, h)
}

func (r *{{NAME}}Rpc) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Data: "Hello " + in.Name}, nil
}
`

var CmdTemplate = `package {{PACKAGE}} 

import (
	"github.com/goees/yago"
	"github.com/goees/yago/base/basecmd"
	"github.com/spf13/cobra"
)

type {{NAME}}Cmd struct {
	basecmd.BaseCmd
}

func init() {
	c := new({{NAME}}Cmd)
	// 注册路由
	yago.AddCmdRouter("demo", "Demo action", c.DemoAction, yago.CmdStringArg{
		Name: "arg", Shorthand: "a", Value: "value", Usage: "参数", Required: true,
	})
}

func (c *{{NAME}}Cmd) DemoAction(cmd *cobra.Command, args []string) {

}

`

var TaskTemplate = `package {{PACKAGE}} 

import (
	"github.com/goees/yago"
	"github.com/goees/yago/base/basetask"
)

type {{NAME}}Task struct {
	basetask.BaseTask
}

func init() {
	t := new({{NAME}}Task)
	yago.AddTaskRouter("@loop", t.HelloAction)
	yago.AddTaskRouter("0 */1 * * * *", t.HelloAction)
}

func (t *{{NAME}}Task) HelloAction() {
	//t.RunLoop(func() {
	//})
}
`

var ModelTemplate = `package {{PACKAGE}} 

type {{LNAME}}Model struct {
}

func New{{NAME}}Model() *{{LNAME}}Model {
	return &{{LNAME}}Model{}
}

`

var ServiceTemplate = `package {{PACKAGE}} 

type {{LNAME}}Service struct {
}

func New{{NAME}}Service() *{{LNAME}}Service {
	return &{{LNAME}}Service{}
}

`
