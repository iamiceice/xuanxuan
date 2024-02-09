# 玄玄一键宏
魔兽世界脚本
- 适用于魔兽世界正式服10.2.5版
- 基于robotgo设计的脚本，原理是在技能序列循环里，把每一个技能转变成按键，就是模拟手动按键。
- 暂时只有兽王猎专精，不能根据BUFF施放技能，只能运行固定序列。
- 建议自己修改源码（主要是按键），自己编译。如果用我编译好的.exe文件，须用我的按键，在myslot兽王.txt里，我使用罗技G600的12侧键鼠标。
- 使用说明：按"4"开始，按"6"暂停，按小键盘"-"结束。运行中，按”8“单体，按"7"是爆发（可能要按两三下，默认爆发后自动转AOE，可参考"爆发监控TMW"），按”9“是AOE,按"0"是打断，按小键盘”7“，”5“，”9“是加血。
- 欢迎合作开发

### 编译
使用如下命令，可生成不带DOS黑窗口的exe文件(在hunter目录下执行)
```
go build -ldflags -H=windowsgui
```

### 计划
根据BUFF施放技能

### License
MIT

### 捐助
![](https://github.com/iamiceice/xuanxuan/blob/main/donate/mm.png)
![](https://github.com/iamiceice/xuanxuan/blob/main/donate/22.jpg)
