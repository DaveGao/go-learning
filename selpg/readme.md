# 关于selpg命令行程序的开发

---
## selpg程序介绍

    selpg 程序允许用户以一定的名命令参数从指定输入文本抽取的页的范围，这些输入文本可以来自文件或另一个进程，然后这个程序将所获得的输入按要求输出到屏幕，指定的文件或者打印机。
 
---
## 程序的使用说明

    使用方法：
    
```
USAGE:  %v -s StartPage -e EndPage [ -f | -l lines_per_page ] [ -d dest ] [ in_filename ]
```
    
    selpg程序的命令介绍如下：
    必选参数----
    -s startPage 代表输入文件开始打印的页序号
    -e endPage 代表输入文件结束打印的页序号
    inputFile 代表选取的输入文件
    可选参数----
    -l lineOfPage 代表每页打印的固定行数
    -f EOF 代表一直打印文件直到读取到输入文件的EOF结束符
    -d destinationOfFile 代表最终打印的输出文件

---
## 程序的设计说明

flag包

    该程序使用了flag包来解析传入程序的各种参数。

os包

    使用了os来处理输入输出的文件。
    
    
---
## 程序测试样例
输入文件为input.txt内容：

> good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere

    
    
下列是各类型命令的测试结果(注意$表示命令行的起始位置)：

1.打印输入文件input.txt的第一页
```
$ ./selpg -s 1 -e 1 input.txt
```
测试结果：

> good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere


2.打印来子管道的第一页
```
$ ./selpg -s 1 -e 1 < input.txt
```
测试结果：

> good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere

3.打印文件的第一页到目标lp打印机
```
$ ./selpg -s 1 -e 1 -d lp
```
测试结果：

> lp: The printer or class does not exist.

4.打印input.txt到out.txt
```
$ ./selpg -s 1 -e 1 input.txt >output.txt
```
测试结果：

> ./selpg: done

5.打印答应输出的错误信息到error.txt
```
$ ./selpg -s 1 -e 1 input.txt 2>error.txt
```
测试结果：
在错误处理文件中看到了输出信息
> ./selpg: done

6.打印输入文件input.txt的固定页数为5行
```
$ ./selpg -s 1 -e 1 -l 5 input.txt
```
测试结果：

> good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere


7.打印input.txt文件直到文件末尾
```
$ ./selpg -s 1 -e 1 -f input.txt

```
测试结果：

> good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere
good things are everyWhere


8.打印文件并丢弃错误到空设备
```
$ ./selpg -s 1 -e 1 input.txt >/dev/null
```
测试结果：

> ./selpg: done


9.打印input.txt的文件并且显示行数、字数和字符数
```
$ ./selpg -s 1 -e 1 input.txt | wc
```
测试结果：

> ./selpg: done
     19      72     487

