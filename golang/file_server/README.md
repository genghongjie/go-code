# 文件服务器

bin目录中有对应的可执行文件

## 下载 

linux可执行文件 
```
wget --no-check-certificate https://github.com/genghongjie/go-code/blob/main/golang/file_server/bin/linux/64/fserver-1.0
```

mac可执行文件

```
wget --no-check-certificate https://github.com/genghongjie/go-code/blob/main/golang/file_server/bin/mac/64/fserver

```
windows可执行文件
```
wget --no-check-certificate https://github.com/genghongjie/go-code/blob/main/golang/file_server/bin/windows/64/fserver.exe
```



指定代理目录需要在启动时指定参数 dir，默认为当前路径

启动示例
```
#指定文件路径
./fserver -dir /tmp
#使用默认路径
./fserver
```

浏览器访问文件内容
```
http://127.0.0.1:10000
```

