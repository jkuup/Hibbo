<p align="center">
    <a href="https://www.alphkiee.xyz" target="_blank">
        <img src="https://raw.githubusercontent.com/jkuup/Hibbo/master/img/gopher.png?v=0.2.2" width="180" />
    </a>
    <h3 align="center">Hibbo</h3>
    <p align="center">Golang 实现的Blog</p>
    <p align="center">
        <a href="https://travis-ci.com/jkuup/Hibbo"><img src="https://travis-ci.com/jkuup/Hibbo.svg?branch=master"></a>
        <a href="https://github.com/jkuup/Hibbo/releases"><img src="https://img.shields.io/badge/Version-v0.0.1-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/jkuup/Hibbo"><img src="https://goreportcard.com/badge/github.com/jkuup/Hibbo?v=0.0.1"></a>
        <a href="https://hub.docker.com/r/jkuup/Hibbo"><img src="https://img.shields.io/badge/Docker-Latest-orange"></a>
        <a href="https://github.com/jkuup/Hibbo/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>
<br/>

Jkuup's Blog


## Documentation

- 新建一个 git 仓库，用来存放你的文章，但是对你的目录结构有一些要求
如下：
```
    |-- assets       //博客静态文件，存放一些图片资源，方便显示到文档里
    |-- content
    |   |-- GOLANG   //分类目录
    |       |-- GOLANG基础   //  子分类目录，但是在页面上不会产生分类目录
    |       |--- GOLANG基础语法.md
    |   |-- 其他分类
    |       |--- xxx.md
    |-- extra_nav  

```
content目录下的一级目录代表一个分类，如果一级目录下有子级目录也不会产生新的分类,子级目录的文档也会属于一级目录的分类。

如下：
```
    ```json
    {
        "date":"2019.01.02 14:33"，//最少需要
        "tags": ["BLOG"，"其它tag"]，//可以不填，不过最好添加一些tag，后面可以做一些好玩的东西。
        "title": "文章的标题，一般不用填写，默认使用文件名"，
        "description": "文章描述，不填写自动取正文200个字，可以在app.json中配置"，
        "author": "Jkuup"， //文章作者，可以不用填写，现在也没有使用到
        "musicId":"网易云的音乐ID" //文章的配歌
    }
    ```
```
## Deployment

Hibbo uses Go Modules to manage dependencies.

```zsh
$ git clone https://github.com/jkuup/Hibbo.git
$ go run main.go
```

## License

© 2020, jkuup. Released under [Apache2.0 License](http://www.apache.org/licenses/LICENSE-2.0).

**Hibbo** is authored and maintained by [@Jkuup](http://github.com/jkuup).