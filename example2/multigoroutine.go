package main

import (
    "github.com/getwe/scws4go"
    "fmt"
    "os"
    "runtime"
    "sync"
    "time"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Printf("Usage : %s dict rule\n",os.Args[0])
        os.Exit(0)
    }

    runtime.GOMAXPROCS(runtime.NumCPU())


    scws:= scws4go.NewScws()

    err := scws.SetDict(os.Args[1], scws4go.SCWS_XDICT_XDB)
    if err != nil { fmt.Println(err);os.Exit(0)}
    err = scws.SetRule(os.Args[2])
    if err != nil { fmt.Println(err);os.Exit(0)}

    scws.SetCharset("utf8")
    scws.SetIgnore(1)
    scws.SetMulti(scws4go.SCWS_MULTI_SHORT & scws4go.SCWS_MULTI_DUALITY & scws4go.SCWS_MULTI_ZMAIN)


    SegRoutineNum := 4

    scws.Init(SegRoutineNum)

    wg := sync.WaitGroup{}

    t1 := time.Now().UnixNano()

    for i:=0;i<SegRoutineNum;i++ {
        wg.Add(1)
        go func() {
            // 每个协程执行10000次切词后退出
            for i:=0;i<10000;i++ {
                _,err := scws.Segment("scws4go是c语言版本的scws的go封装.")
                if err != nil {
                    fmt.Println(err)
                    os.Exit(0)
                }
            }
            wg.Done()
        }()
    }

    // 等待全部协程退出
    wg.Wait()
    t2 := time.Now().UnixNano()
    fmt.Printf("耗时%d ns\n",(t2-t1))

    scws.Free()
}
