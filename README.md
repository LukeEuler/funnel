# Funnel

本项目力求在最精简依赖的情况下，实现一个`json`格式的日志报警引擎。项目的各个模块也都尽力做到边界清晰，以确保任何人可以自由的改写以便适应其需求场景。甚至，你可以用它做非`json`格式的日志报警引擎。

首先，我想要感谢[antlr4](https://www.antlr.org/)，及其作者[Terence Parr](https://explained.ai/)。这是一个天才的作者，实现的的天才的程序。

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
