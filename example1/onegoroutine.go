package main

import (
    "github.com/getwe/scws4go"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Printf("Usage : %s dict rule\n",os.Args[0])
        os.Exit(0)
    }
    scws:= scws4go.NewScws()

    err := scws.SetDict(os.Args[1], scws4go.SCWS_XDICT_XDB)
    if err != nil { fmt.Println(err);os.Exit(0)}
    err = scws.SetRule(os.Args[2])
    if err != nil { fmt.Println(err);os.Exit(0)}

    scws.SetCharset("utf8")
    scws.SetIgnore(1)
    scws.SetMulti(scws4go.SCWS_MULTI_SHORT & scws4go.SCWS_MULTI_DUALITY & scws4go.SCWS_MULTI_ZMAIN)

    // 单协程只需要一个切词实例就够了
    scws.Init(1)

    segResult,err := scws.Segment("scws4go是c语言版本的scws的go封装.")
    if err != nil {
        fmt.Println(err)
    } else {
        for _,t := range segResult {
            fmt.Printf("term[%s],attr[%s],idf[%f]\n",t.Term,t.Attr,t.Idf)
        }
    }

    scws.Free()
}
