# Go Server on Android

一个非常简单的，在android上运行服务器的方案！


## Why

go打包的程序可以跨平台运行，那么android既然本质上是一个linux系统，在上面运行服务器理论上是没问题的。

## How

可以直接运行`python build_server.py`自动完成。

他可以：

- 编译出符合android规格的包 (linux arm)
    - `GO_ENABLED=0 GOARCH=arm GOOS=linux go build go_android_server.go`
- 移动到手机中并赋予运行权限
    - `adb push go_android_server /data/local/tmp`
- 启动服务器
    - `adb shell /data/local/tmp/go_android_server`
- 在设置中查看你的手机ip
    
通过设置参数，可以：

- `python build_server.py -k` 杀死现在正在运行的服务器
- `python build_server.py -b` 只进行打包
