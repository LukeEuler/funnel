# Funnel



## Installing ANTLR v4

```
# 安装到其他目录
wget http://www.antlr.org/download/antlr-4.9.2-complete.jar
alias antlr='java -jar $PWD/antlr-4.9.2-complete.jar'

# 项目目录
# 生成语法解析的 golang 代码
cd antlr4
antlr -Dlanguage=Go -o parser Match.g4
```
