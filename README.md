# Discard
因為 microsoft的原因 此項目 已經被轉移到 [gitlab.com/king011/ng-xi18n](https://gitlab.com/king011/ng-xi18n)

github上的 版本將不在進行 任何維護


# ng-xi18n
Angular xi18n tools

Angular 提供了 ng xi18n 命令 可以將模板中的 待翻譯 條目 整理 出來 得到檔案 messages.xlf 而 ng-xi18n 則提供了 將 messages.xlf 更新到 dist.xlf 的 update 指令 update 會使用 messages.xlf 覆蓋 dist.xlf 然會 和 copy 不同的 是 會保留 dist.xlf 原本的 語言設置 和 已經翻譯好的 target 內容

ng-xi18n 只 支持 xlf 格式的 檔案

# Install
1. go get -u -d github.com/zuiwuchang/ng-xi18n
2. cd $GOPATH/src/github.com/zuiwuchang/ng-xi18n
3. ./build-go.sh
4. go install

# How To Use
1. cd ${yourProject}
2. mkdir -p src/locale
3. ng-xi18n update -l zh-Hant
```bash
$ ./ng-xi18n update -h
new message file
	ng-xi18n update -s src/messages.xlf -d src/locale/zh-Hant.xlf 
	ng-xi18n update -s src/messages.xlf -d src/locale/zh-Hant.xlf
	ng-xi18n update -s src/messages.xlf -l zh-Hant

Usage:
  ng-xi18n update [flags]

Flags:
  -d, --dist string     distribution file,if empty use src/locale/$locale.xlf
  -h, --help            help for update
  -l, --locale string   locale zh-Hant zh-Hant-TW zh-Hant-HK ...
  -s, --src string      source file,use ng xi18n get it. (default "src/messages.xlf")
```
