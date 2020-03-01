/**
* curl http://localhost:8080/
* curl --data "users=yifan,yifan2" http://localhost:8080/import
* curl http://localhost:8080/lucky
*/
package main

import (
    "github.com/kataras/iris"
    "annualMetting/controller"
)

func main() {
    app := controller.NewApp()
    // useList = make([]string, 0)
    //UserList = []string{}
    controller.Init()
    app.Run(iris.Addr(/*addr:*/ ":8080"))
}
