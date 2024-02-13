# 玄玄一键宏
魔兽世界脚本
- 适用于魔兽世界正式服10.2.5版
- 基于robotgo设计的脚本，原理是在技能序列循环里，把每一个技能转变成按键，就是模拟手动按键。
- 暂时只有兽王猎和增强萨职业专精，不能根据BUFF施放技能，只能运行固定序列。
- 建议自己修改源码（主要是按键），自己编译。如果用我编译好的.exe文件，须用我的按键，在myslot兽王猎.txt和myslot增强萨里，我使用罗技G600的12侧键鼠标。
- 使用说明：按"4"开始（默认AOE），按"6"暂停，按小键盘"-"结束。运行中，按”8“单体（增强闪电箭），按"7"是爆发，按”9“是AOE（增强闪电链）,按"0"是打断，按”1“，”2“，”3“是加血。(注：除了打断，其它按键都会延迟1秒执行)
- 欢迎合作开发

### 编译
使用如下命令(在hunter目录下执行)，生成带DOS黑窗口的exe文件，黑窗口不见就是一键宏停止了，所以可随时查看一键宏还在不在。
```
go build 
```
使用如下命令(在hunter目录下执行)，可生成不带DOS黑窗口的exe文件
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
