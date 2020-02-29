/**
* curl http://localhost:8080/
* curl --data "users=yifan,yifan2" http://localhost:8080/import
* curl http://localhost:8080/lucky
*/
package main

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/mvc"
    "time"
    "fmt"
    "strings"
    "math/rand"
)

var userList []string

type lotteryController struct {
    Ctx iris.Context
}

func newApp() * iris.Application {
    app := iris.New()
    mvc.New(app.Party(/*relativeParth:*/ "/")).Handle(&lotteryController{})
    return app
}

func main() {
    app := newApp()
    // useList = make([]string, 0)
    userList = []string{}
    app.Run(iris.Addr(/*addr:*/ ":8080"))
}

func (c *lotteryController) Get() string {
    count := len(userList)
    return fmt.Sprintf(/*format*/ "当前总共参与抽奖的用户数: %d\n", count)
}

// POST http://localhost:8080/import
// params: users
func (c *lotteryController) PostImport() string {
    strUsers := c.Ctx.FormValue("users")
    users := strings.Split(strUsers, ",")
    count1 := len(userList)
    for _, u := range users {
        u = strings.TrimSpace(u)
        if len(u) > 0 {
            userList = append(userList, u)
        }
    }
    count2 := len(userList)
    return fmt.Sprintf(/*format*/ "当前总共参与当前抽奖的用户数: %d, 成功导入的用户数: %d\n",
        count1, (count2-count1))
}

// GET http://localhost:8080/lucky
func (c *lotteryController) GetLucky() string {
    count := len(userList)
    if count > 1 {
        seed := time.Now().UnixNano()
        index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
        user := userList[index]
        userList = append(userList[0:index], userList[index+1:]...)
        return fmt.Sprintf(/*format*/ "当前中奖用户: %s, 剩余用户数: %d\n", user, count-1)
    } else if count == 1{
        user := userList[0]
        userList = []string{}
        return fmt.Sprintf(/*format*/ "当前中奖用户: %s, 剩余用户数: %d\n", user, count-1)
    } else {
        return fmt.Sprintf(/*format*/ "当前没有参与用户，请先通过 /import 导入用户\n")
    }
}
