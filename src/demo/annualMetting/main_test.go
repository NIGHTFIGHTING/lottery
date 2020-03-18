package main

import (
    "testing"
    "github.com/kataras/iris/httptest"
    "sync"
    //"github.com/iris-contrib/httpexpect"
    "annualMetting/controller"
    "fmt"
)

func TestMVC(t *testing.T) {
    e := httptest.New(t, controller.NewApp())

    var wg sync.WaitGroup
    e.GET(/*path*/ "/").Expect().Status(httptest.StatusOK).
        Body().Equal(/*value*/ "当前总共参与抽奖的用户数: 0\n")

    for i := 0; i < 100; i++ {
        wg.Add(/*delta*/ 1)
        go func(i int) {
            defer wg.Done()
            e.POST(/*path*/ "/import").WithFormField(/*key*/"users",
                fmt.Sprintf(/*format*/ "test_u%d", i)).Expect().
                Status(httptest.StatusOK)
        }(i)
    }
    wg.Wait()

    e.GET(/*path*/ "/").Expect().Status(httptest.StatusOK).
        Body().Equal(/*value*/ "当前总共参与抽奖的用户数: 100\n")
    e.GET(/*path*/ "/lucky").Expect().Status(httptest.StatusOK)
    e.GET(/*path*/ "/").Expect().Status(httptest.StatusOK).
        Body().Equal(/*value*/ "当前总共参与抽奖的用户数: 99\n")
}
