package in

import (
    "fmt"
    "testing")

func TestCreateLinkService(t *testing.T) {
    linkService := NewLinkService()
    uid, err := linkService.Create("url", "name", "description")
    fmt.Println(uid, err)
    if err != nil {
        t.Fail()
    }
}
