/**
 * 1 即开即得型
 * 2 双色球自选型
 */
package main

import (
    "github.com/kataras/iris"
    "demo/ticket/controller"
)

func main() {
    app := controller.NewApp()
    app.Run(iris.Addr(/*addr:*/ ":8080"))
}
