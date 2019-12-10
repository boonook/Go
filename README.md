# Go
# 断点调试
# go模块包下载的第三方代理
open -e .bash_profile 打开配置文件
export GO111MODULE=on
export GOPROXY=https://goproxy.io
source /etc/profile
关于goproxy，简单来说就是一个代理，让我们更方便的下载哪些由于墙的原因而导致无法下载的第三方包，比如golang.org/x/下的包，虽然也有各种方法解决，但是，如果是你在拉取第三方包的时候，而这个包又依赖于golang.org/x/下的包，你本地又恰恰没有，当然不嫌麻烦的话，也可以先拉取golang.org/x/下的包，再拉取第三方包。