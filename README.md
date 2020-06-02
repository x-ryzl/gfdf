# gfdf
A file directory comparison tool (一个文件夹对比工具)

# 功能：

查看两个文件夹所有差异的文件，以及每个文件夹的文件数量和文件差数量


# 编译
go build -o gfdf ./main.go

# 查看运行参数
./gfdf -h

  -e 忽略文件夹(非必填)
        
  -s 源路径(必填)
        
  -t 目标路径(必填)
  
  -c 命令参数(0：文件夹差异 1：md5判断文件)
         
# 文件夹差异文件执行命令：
./gfdf -s src -t target -e ext -c 0

# 根据md5判断文件是否相同命令:
./gfdf -s src -t target -c 1



