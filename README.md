# gfdf
A file directory comparison tool (一个文件夹对比工具)


# 打包
go build -o gfdf ./main.go

# 查看运行参数
./gfdf -h

  -e 忽略文件夹(非必填)
        
  -s 源路径(必填)
        
  -t 目标路径(必填)
         
# 执行命令
./gfdf -s src -t target -e ext


