# go_hotreload

[English](README.md) | 中文

[![Release](https://img.shields.io/github/v/release/zhangga/go_hotreload)](https://github.com/zhangga/go_hotreload/releases)
[![License](https://img.shields.io/github/license/zhangga/go_hotreload)](https://github.com/zhangga/go_hotreload/blob/main/LICENSE)

适用于 Go 应用的轻量级热重载框架。

## 支持平台
* Linux amd64
* Linux arm64
* MacOS amd64
* MacOS arm64
* Windows amd64

## 快速开始
1. 编译要热重载的patch文件： `make patch`。
2. 启动示例程序： `make example`。
3. 访问 [WebUI](http://localhost:8080) 界面，点击`执行热更`按钮，触发热重载。
4. 点击`测试后端执行`按钮，测试热重载是否成功。

## 使用说明
参考 [example](example) | 目录下的示例代码。
* [example/internal](example/internal) 中的示例代码为主程序代码。
* [example/patch_v1](example/patch_v1) 中的示例代码为热重载的 patch 代码。
