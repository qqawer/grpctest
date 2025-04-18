**操作命令**

1. go mod init grpctest
2. go mod tidy
3. 有些最新依赖超出了本地golang1.22.9版本所以需要自己go get  
   例如go get golang.org/x/crypto@v0.32.0
4. go env -w GOTOOLCHAIN=local  
   过设置工具链为 local 来防止 Go 升级  
   go env GOTOOLCHAIN查看命令
5. go build main.go->main.exe -h->main.exe -port 50053
