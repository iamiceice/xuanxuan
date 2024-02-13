package main

import (
	"context"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"os"
	"sync"
	"time"
)

// 脚本循环的频率
// const FREQUENCY int64 = 10
var ch = make(chan bool)
var run bool
var flag bool
var lock sync.Mutex

// 创建一个管道，用于接收按键事件
var pipe = make(chan byte)
var char byte

// 生成可以取消的 context
var ctx, cancel = context.WithCancel(context.Background())

const (
	闪电箭 int64 = 1000 + iota
	闪电链
	切换目标
	始源之潮
	烈焰震击
	冰霜震击
	风暴打击
	野性狼魂
	毁灭闪电
	霜刃打击
	熔岩猛击
	元素冲击
	裂地术
	毁灭之风
	治疗之涌
	饰品药水
	大红瓶
	治疗石
	打断
)

// 技能和按键映射
var SpellKeyMap = map[int64][]string{
	闪电箭:  []string{"k"},
	闪电链:  []string{"l"},
	始源之潮: []string{"o", "shift"},
	冰霜震击: []string{"u"},
	烈焰震击: []string{"o"},
	风暴打击: []string{"5", "ctrl"},
	野性狼魂: []string{"l", "alt"},
	毁灭闪电: []string{"l", "ctrl"},
	霜刃打击: []string{"k", "alt"},
	熔岩猛击: []string{"k", "ctrl"},
	元素冲击: []string{"u", "ctrl"},
	裂地术:  []string{"l", "shift"},
	毁灭之风: []string{"5", "shift"},
	治疗之涌: []string{"-"},
	打断:   []string{"."},
	切换目标: []string{"5"},
	饰品药水: []string{"[", "ctrl", "shift"},
	大红瓶:  []string{";"},
	治疗石:  []string{"'"},
}

func main() {
	//技能循环
	go loop()
	//处理快捷键
	go fork()
	//快捷键
	shortcutkey()
}

// 快捷键
func shortcutkey() {
	hooks := hook.Start()
	defer hook.End()
	for ev := range hooks {
		//	监听键盘弹起
		if ev.Kind == hook.KeyDown {
			//以下是快捷键，不能与施放技能的按键相同！！！
			//按快捷键“4”开始脚本
			if ev.Rawcode == 52 {
				ch <- true
			}
			//按快捷键“6”暂停脚本
			if ev.Rawcode == 54 {
				lock.Lock()
				run = false
				lock.Unlock()
				ch <- false
			}
			//按快捷键小键盘“-”停止脚本软件
			if ev.Rawcode == 109 {
				os.Exit(0)
			}
			//按快捷键“0”为打断
			if ev.Rawcode == 48 {
				cast(打断)
			}
			//按快捷键“7”为饰品药水
			if ev.Rawcode == 55 {
				pipe <- '7'
			}
			//按快捷键“8”为闪电箭
			if ev.Rawcode == 56 {
				pipe <- '8'
			}
			//按快捷键“9”为闪电链
			if ev.Rawcode == 57 {
				pipe <- '9'
			}
			//按快捷键小键盘"1"大红瓶
			if ev.Rawcode == 49 {
				pipe <- '1'
			}
			//按快捷键小键盘"2"术士治疗石35
			if ev.Rawcode == 50 {
				pipe <- '2'
			}
			//按快捷键小键盘"3"加血34
			if ev.Rawcode == 51 {
				pipe <- '3'
			}
		}
	}
}

// 脚本开始
func loop() {
	for {
		//脚本开始时间
		//timestart := time.Now()
		if run == true {
			cast(始源之潮)
			cast(风暴打击)
			cast(野性狼魂)
			cast(毁灭闪电)
			cast(霜刃打击)
			cast(冰霜震击)
			cast(裂地术)
			cast(毁灭之风)
		} else {
			run = <-ch
		}
		//脚本结束时间
		//timeend := time.Now()
		//delay(timestart, timeend)
		//根据暂停flag判断是否暂停
	}
}

func fork() {
	for {
		char = <-pipe
		lock.Lock()
		flag = run
		run = false
		lock.Unlock()
		ch <- false
		time.Sleep(time.Second)
		switch char {
		case '7':
			cast(饰品药水)
		case '8':
			cast(闪电箭)
		case '9':
			cast(闪电链)
		case '1':
			cast(大红瓶)
		case '2':
			cast(治疗石)
		case '3':
			cast(治疗之涌)
		default:

		}
		if flag {
			ch <- true
		}
	}
}

// 施放技能
func cast(spell int64) {
	// 获得技能对应的按键  spell:技能 key:按键
	key := spell2key(spell)
	//技能对应的第一个按键
	k0 := key[0]
	// 技能对应的控制键序列，如ctrl shift alt
	k1n := key[1:]
	// 如果是涉及鼠标滚轮，单独处理
	switch k0 {
	case "上滚":
		k0 = "up"
	case "下滚":
		k0 = "down"
	case "左滚":
		k0 = "left"
	case "右滚":
		k0 = "right"
	default:

	}
	presskey(k0, k1n)
}

// 按下按键
func presskey(k0 string, k1n []string) {
	// robotgo的按键函数
	robotgo.KeyTap(k0, k1n)
}

// robotgo的鼠标滚轮操作
func wheelkey(k0 string, k1n []string) {
	for _, v := range k1n {
		robotgo.KeyToggle(v, "down")
	}
	robotgo.ScrollDir(1, k0)
	for _, v := range k1n {
		robotgo.KeyToggle(v, "up")
	}
}

// 根据技能，从“技能-按键映射”中获得按键序列
func ToKey(c int64) []string {
	tokey := SpellKeyMap[c]
	return tokey
}

// 输入技能，输出按键  spell:技能 key:按键
func spell2key(spell int64) (key []string) {
	// 从“技能-按键映射”中获得按键序列
	key = ToKey(spell)
	return key
}

//func delay(timestart, timeend time.Time) {
//	timelong := timeend.Sub(timestart).Microseconds()
//	//脚本每次循环的周期
//	var cycle int64 = 1000000 / FREQUENCY
//	slog.Info("周期=", cycle)
//	delaytime := cycle - timelong
//	slog.Info("delaytime=", delaytime)
//	time.Sleep(time.Duration(delaytime * 1000))
//	slog.Info("Sleep:", time.Duration(delaytime*1000))
//}
