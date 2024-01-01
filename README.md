# 玄玄一键宏
魔兽世界脚本
- 适用于魔兽世界正式服10.2版
- 基于robotgo设计的脚本，原理是在技能序列循环里，把每一个技能转变成按键，就是模拟手动按键。
- 暂时不能根据BUFF施放技能，只能运行固定序列。
- 建议自己修改源码（主要是按键），自己编译。如果用我编译好的.exe文件（天赋是火流），须用我的按键，在myslot元素.txt里。
- 欢迎合作开发
### 编译
- 使用如下命令，可生成不带DOS黑窗口的exe文件
```
go build -ldflags -H=windowsgui
```

### License
MIT

### 捐助
![](https://github.com/iamiceice/xuanxuan/blob/main/donate/mm.png)
![](https://github.com/iamiceice/xuanxuan/blob/main/donate/22.jpg)
