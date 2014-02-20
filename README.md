#scws4go
这是一个对[scws](github.com/hightman/scws)的go绑定,区别于[bylevel/goscws](github.com/bylevel/goscws)的是提供更加简单的可并发的接口.

#安装方法
##安装hightman/scws
请参考[scws安装方法](https://github.com/hightman/scws#%E5%AE%89%E8%A3%85).

在我机器上,我是这样把第三方库安装在我的$HOME/usr/目录下的:

    cd scws-1.2.1 ; ./configure --prefix=$HOME/usr/ ; make install
    
如果你也这样安装,请确认环境变量CPLUS_INCLUDE_PATH包含了$HOME/usr/include;LD_LIBRARY_PATH包含了$HOME/usr/lib.cgo使用gcc编译器的时候需要靠这两个变量来找到scws.(这些都是太基础东西了,请原谅我啰嗦了一次.)
    
##安装
    go get github.com/getwe/scws4go
##下载词典
到[xunsearch](http://www.xunsearch.com/scws/download.php)下载词典和规则文件.
#例子
##例子1,简单使用
    package main
    func main(){
        // TODO
    }
##例子2,多goroutine下使用
    package main
    func main(){
        // TODO
    }

