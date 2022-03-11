# funnel 过程简述

### 第一次上色

设计`funnel`的初衷, 是为了可配置化的对日志数据做监控. 主要的需求目标为以下几点:
- 利用词法规则编写过滤规则`rule`
- `rule`应包含以下判别功能:存在、等于、包含, 与、或、非、优先
- 数据`data`与`rule`可以做匹配判别, 结果有且仅有两种结果: `match`, `not match`

#### 图例

我们可以依据以上规则, 对`data`和`rule`做匹配: 上色
```ditaa {cmd=true args=["-E"]}
  +--------+
  |        |
  |  Data  |--+
  |  cBLU  |  |         +----------+
  +---+----+  | match   |Data cBLU |
              +-------> +----------+
  +--------+  |         |r1 cRED   |
  |  r1    |  |         +----------+
  +--------+--+         | ...      |
  | rule   |            +----------+
  | cRED   |
  +--------+

  ```

当然,实际生产中, 日志和规则都远不止一条. 从监控的角度看, 我们更关心的是数据本身.  所以, 对于那些不匹配的规则, 是可以忽略的.

```ditaa {cmd=true args=["-E"]}
  +----------+  +----------+  +----------+  +----------+
  |d1   cBLU |  |d2   cBLU |  |d3   cBLU |  |d4   cBLU |
  +----------+  +----------+  +----------+  +----------+
  |r1 cRED   |  |r2 cRED   |  |          |  |r3 cRED   |
  +----------+  +----------+  |          |  +----------+
  |r3 cRED   |  |          |  |          |  |r6 cRED   |
  +----------+  |          |  |          |  +----------+
  |r18 cRED  |  |          |  |          |  |          |
  +----------+  |          |  |          |  |          |
  |r27 cRED  |  |          |  |          |  |          |
  +----------+  +----------+  +----------+  +----------+
 ```

#### 判别功能

对于`rule`判别功能, `funnel`定义了自己的语法规则:
|判别|EXIST|EQUAL|CONTAINS|AND|OR|NOT|优先|
|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
|符号|`+`| `=` | `>` | `&` | `|` | `!` | `()` |


举例说明:
我们可以假设, 日志都是`json` 格式的
```json
{
    "level":"error",
    "message":"connect to www.example.com failed, timeout",
    "region":"antarctic",
    "func": "GetRemainPeople",
    "file": "/data/.../client.go:123",
    "host": "a.bb.pc.com",
    "cost": {
        "response": 213,
    },
    "tag": "request",
    "time": "2022-09-27T17:15:21+08:00",
}
```
那么规则编写就可以如下

`level = 'error' & message > 'imeout' & (region = 'africa' | region = 'antarctic' | cost.response+)`

通过`第一次上色`, 我们可以很容易的将数据区分开. 对于日志而言, 我们就可以依据规则过滤出部分数据了.
然而, 在实际场景中, 这是远远不够的.
- 通常我们会有一个更加高效的办法直接提取错误日志, 如 `level = error`
- 对于符合某些规则的日志, 我们可以忍受其发生在一定的频率下而不做报警. 如在高频网络请求中, 偶发的连接超时
- 规则之间也应可以相互影响. 如, 某些情况下, 我们明确知道特殊类别的日志（来自某个坏掉的组件,并且不需要修理), 可以忽略. 那么这些符合特征的日志, 就可以不再参与另外的规则的匹配, 避免无效的报警发生(这些不需要修理的组件,造成的超时连接日志都可以去掉)

### 第二次上色

日志报警总是对时间比较敏感:
- 最新的报警更加重要
- 密集发生的同类事件预示着有情况了

所以`data`需要升级为时序数据(当然, 实际生产中, 我们一般会关注最近一段时间的数据, 如最近一小时).

同样`rule`也需要升级
- **level**: 用以定义规则等级, 用以优先匹配, 以及影响低级别`rule`
- **mutex**: 用以决定是在在匹配时, 将`data`排除在低级别`rule`的匹配过程
- **start, end**: `rule`的有效时间范围. 0 值代表忽略边界
- **duration, times**: 匹配的频次阈值

因此, 对于一个`data`, 一个`rule`的上色场景中, `data`的上色状态将有三种.
- **not match**: 数据与规则不匹配
- **match & hit**: 数据与规则匹配, 且达到报警条件
- **match & miss**: 数据与规则匹配, 但未达到报警条件

#### 图例

```ditaa {cmd=true args=["-E"]}
  +----------+  +----------+  +----------+  +----------+
  |d1   cBLU |  |d2   cBLU |  |d3   cBLU |  |d4   cBLU |
  +----------+  +----------+  +----------+  +----------+
  |r1 cPNK   |  |r2 cRED   |  |          |  |r3 cPNK   |
  +----------+  +----------+  |          |  +----------+
  |r3 cPNK   |  |          |  |          |  |r6 cRED   |
  +----------+  |          |  |          |  +----------+
  |r18 cPNK  |  |          |  |          |  |          |
  +----------+  |          |  |          |  |          |
  |r27 cPNK  |  |          |  |          |  |          |
  +----------+  +----------+  +----------+  +----------+
 ```

### 重要补充
实际生产中, 通常我们只会把error数据做过滤。一方面，这样提取数据, 通常比较简单且快速; 另一方面, 这样可以避免浪费 funnel 的计算资源. 所以, 经过处理前的数据, 一般都是要报警的. funnel 只是把其中的噪音抹掉.

基于以上事实, 我们做如下讨论: 一个`data`会与多个`rule`做匹配.
- 如果所有的规则都都未匹配, 那么这条数据应该参与报警. 此时发生的事情, 是预期之外的
- 如果存在部分规则与该数据匹配, 但所有规则都不命中. 那么, 这条数据可以在报警中忽略
- 如果存在部分规则与该数据匹配, 且这其中存在规则命中了. 那么, 这条数据应报警

> 如何使用该程序, 是使用方的自由. 所以大可不必拘泥于本人的建议. 建议仅供参考

> 报警数据当然不需要逐条日志报警, 频繁的报警等于没有报警. 那么, 为了有效话报警, 报警数据应做聚合处理. 个人建议在收集完所有需要报警的日志数据后, 应该忘记其匹配的规则, 使用更加贴合业务的方法来聚合日志(如`json`日志中的key字段)