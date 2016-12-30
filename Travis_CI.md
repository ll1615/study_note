# 简介
> Travis CI 是目前新兴的开源持续集成构建项目，它与jenkins，GO的很明显的特别在于采用yaml格式，简洁清新独树一帜。目前大多数的github项目都已经移入到Travis CI的构建队列中，据说Travis CI每天运行超过4000次完整构建。

# 新手入门
> 参考: https://docs.travis-ci.com/user/for-beginners

- 在 Github 上 fork [PHP示例仓库](https://github.com/plaindocs/travis-broken-example)

- 使用 github 账号[登录 Travis CI](https://travis-ci.org/auth), 通过 Github 的授权确认.

- 登录之后 Travis CI 会同步你 Github 上的仓库, 进入你的[管理页面](https://travis-ci.org/profile), 确认 Travis CI 构建你 fork 的仓库 "travis-broken-example".
- travis.yaml 示例
``` yaml
  language: php
  php:
  - 5.5
  - 5.4
  - hhvm
  script: phpunit Test.php
```
> 说明: 这个文件告诉 Travis CI "该项目使用 PHP 语言书写, 使用 PHP 版本5.5|5.4|HHVM, 执行脚本命令 'phpunit Test.php' 进行单元测试".

- 编辑空文件**NewUser.txt**, 将你的名字添加进去. 然后使用 git 将该文件提交至 Github 仓库, 开始进行 Travis CI 构建. 示例如下:
``` bash
$ git add -A
$ git commit -m 'Testing Travis CI'
$ git push
```
等待 Travis CI 在你 fork 的示例仓库 travis-broken-example 上进行构建, 检查[构建状态](https://travis-ci.org/repositories), 你会注意到构建失败了.(发生错误时 Travis CI 会发送邮件至你的邮箱)
> 注意: Travis 只会在提交的仓库添加至 Travis 之后进行构建.

- 修复代码, 确认 __2=1+1__ 在 __Test.php__ 文件中. 重新提交至 Github. 这次构建将会成功.
```
$ git add -A
$ git commit -m 'Testing Travis CI: fixing the build'
$ git push
```

恭喜, 至此你已添加了一个 Github 仓库到 Travis 并学习了构建和测试代码的基础.

> 注意: 你不需要提交 pull request 至原始仓库, 构建是在你 fork 的仓库上进行的.
