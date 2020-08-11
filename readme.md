## 基于golang的个人简历项目

GitHub：https://github.com/otiedano

邮箱：1392401799@qq.com



> **该内容为后端API**



### 项目配置

后端：golang + mysql + redis

前端：vue-cli + vuex + vue-router + axios



### 项目运行

1. 拉取项目到本地

2. 项目使用```go mod```来管理依赖。需要配置gomodule111.

   ```
   GO111MODULE="on"
   GOPROXY="https://goproxy.cn,direct"
   ```

   在shell输入```go env -w yourconfig=value ```可写入go环境配置,然后通过```go env```查看是否配置成功

3. shell中cd到项目根目录，```go build```编译，然后运行生成的可执行文件，项目即可运行。

4. 如果需要进行开发，本项目需要配置全局变量名```SZRESUME```

   格式：

   ```
   SZRESUME=/Users/YOUR/FILE/PATH
   ```

   项目依赖配置文件。由于```go build``` 、```go run```、```go test```实际运行位置都不一样。开发时使用相对路径往往读不到配置文件和日志文件。

   mac电脑可到～/.profile内加入。或者在shell中导出全局变量。

   我使用的shell是zsh,在.zshrc文件中追加：

   ```
   export SZRESUME=/Users/YOU/FILE/PATH
   ```

   如果不配置全局，则只有项目根目录运行```go bulid```能够正常运行。```go run```、```go test```由于是在系统临时文件生成可执行文件并运行，找不到相对路径的配置文件和日志文件。
   目前项目部署到阿里云,linux版本centos8,golang的os.env无法读取系统全局变量，暂时未做进一步测试确认。不影响项目运行，忽略报错就可以 。

   > 看过一些文章，读取配置文件有几种方式：1.运行时通过go提供的解析目录不能解决根本问题。2.使用命令行参数，配起来也麻烦。```go test```和其他两种模式，配置命令行参数不一样。3.全局变量这种方式最无脑有效。当然也可自行使用系统已有的变量。比如HOME。

​     

### VS code 配置

#### go开发相关插件

使用VS code开发golang应用，需要安装go插件。在VS code按快捷键```command+shift+p```搜索go，可以将关于go的插件都安装。国内十有八九安装失败，需要科学上网。我是试了好多次，结果只能一个一个去github拉，但是最后有一次莫名其妙就成功了。

#### 

安装gopls后应该在VS code的settings.json中输入如下内容。

```
    "go.gopath":"/Users/mac/Documents/go",
    "go.toolsGopath":"/Users/mac/Documents/go",
    "window.zoomLevel": 0,
    "go.useLanguageServer": true,
    "[go]": {
        "editor.snippetSuggestions": "none",
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true
        }
    },
    "gopls": {
        "usePlaceholders": true, 
        "wantCompletionDocumentation": true
    },
    "files.eol": "\n", // formatting only supports LF line endings
```



### 库

#### gin

轻量级后端框架，用起来很方便。

#### zap

日志收集工具，轻量但功能十分强大，执行效率很高。数据导出需要搭配lumberjack使用。包括日志json格式字段导出。日志按大小或日志分割。日志分级过滤。多个输出源自由定制。执行效率很高。

#### beego valid

表单验证，使用方便。支持很多常见的表单验证方式。

#### go-ini

配置文件加载和解析工具。



