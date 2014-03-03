#scws4go
这是一个对[scws](https://github.com/hightman/scws)的go绑定,区别于[bylevel/goscws](https://github.com/bylevel/goscws)的是提供更加简单的可并发的接口.

#安装方法
##安装hightman/scws
请参考[scws安装方法](https://github.com/hightman/scws#%E5%AE%89%E8%A3%85).

在我机器上,我是这样把第三方库安装在我的$HOME/usr/目录下的:

    cd scws-1.2.1 ; ./configure --prefix=$HOME/usr/ ; make install
    
如果你也这样安装,请确认环境变量的设置:

* CPLUS_INCLUDE_PATH包含了$HOME/usr/include
* LIBRARY_PATH包含了$HOME/usr/lib

cgo使用gcc编译器的时候需要靠这两个变量来找到scws.(这些都是太基础东西了,请原谅我啰嗦了一次.)
    
##安装
    go get github.com/getwe/scws4go
##下载词典
到[xunsearch](http://www.xunsearch.com/scws/download.php)下载词典和规则文件.
#例子
##例子1,简单使用
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

        scws.Init(1)

        segResult,err := scws.Segment("scws4go是c语言版本的scws的go封装.")
        if err != nil {
            fmt.Println(err)
        } else {
            for _,t := range segResult {
                fmt.Printf("term[%s],attr[%s],idf[%f]\n",t.Term,t.Attr,t.Idf)
            }
        }
    }


并发例子,详见example2目录   

