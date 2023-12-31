package main

import (
	"context"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"os"
	"runtime"
)

// 脚本循环的频率
// const FREQUENCY int64 = 10
var stop bool

// 生成可以取消的 context
var ctx, cancel = context.WithCancel(context.Background())
var 目标数量 string = "AOE"

const (
	闪电箭 int64 = 1000 + iota
	闪电链
	神器
	岩浆图腾
	幽灵步
	始源之潮
	烈焰震击
	火元素
	熔岩爆裂
	大地震击
	回收图腾
	地震术
	切换目标
)

// 技能和按键映射
var SpellKeyMap = map[int64][]string{
	闪电箭:  []string{"8"},
	闪电链:  []string{"9"},
	神器:   []string{"5", "shift"},
	幽灵步:  []string{"上滚", "control"},
	始源之潮: []string{"6", "shift"},
	烈焰震击: []string{"6"},
	火元素:  []string{"下滚", "shift"},
	熔岩爆裂: []string{"6", "ctrl"},
	大地震击: []string{"5", "ctrl"},
	岩浆图腾: []string{"9", "shift"},
	回收图腾: []string{"9", "shift"},
	地震术:  []string{"9", "ctrl"},
	切换目标: []string{"5"},
}

func main() {
	//快捷键
	shortcutkey()
	fmt.Println("start之前")
	//暂停flag
	stop = true
}

// 快捷键
func shortcutkey() {
	hooks := hook.Start()
	defer hook.End()
	for ev := range hooks {
		//	监听键盘弹起
		if ev.Kind == hook.KeyDown {
			//按快捷键“-”开始脚本
			if ev.Rawcode == 109 {
				stop = false
				go start()
			}
			//按快捷键“+”暂停脚本
			if ev.Rawcode == 107 {
				stop = true
			}
			//按快捷键“/”停止脚本软件
			if ev.Rawcode == 111 {
				os.Exit(0)
			}
			//按快捷键“1”为单体
			if ev.Rawcode == 49 {
				目标数量 = "单体"
			}
			//按快捷键“2”为1-3目标
			if ev.Rawcode == 50 {
				目标数量 = "1-3"
			}
			//按快捷键“3”为AOE
			if ev.Rawcode == 51 {
				目标数量 = "AOE"
			}
		}
	}
}

// 脚本开始
func start() {
	for {
		//根据暂停flag判断是否暂停
		timedelay()
		//脚本开始时间
		//timestart := time.Now()
		if 目标数量 == "单体" {
			cast(火元素)
			cast(烈焰震击)
			cast(切换目标)
			cast(始源之潮)
			cast(岩浆图腾)
			cast(熔岩爆裂)
			cast(神器)
			cast(大地震击)
			cast(闪电箭)
		} else if 目标数量 == "1-3" {
			cast(火元素)
			cast(烈焰震击)
			cast(切换目标)
			cast(始源之潮)
			cast(岩浆图腾)
			cast(熔岩爆裂)
			cast(地震术)
			cast(闪电链)
		} else {
			cast(火元素)
			cast(烈焰震击)
			cast(切换目标)
			cast(始源之潮)
			cast(岩浆图腾)
			cast(熔岩爆裂)
			cast(神器)
			cast(地震术)
			cast(闪电链)
		}
		//脚本结束时间
		//timeend := time.Now()
		//delay(timestart, timeend)
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
	if k0 == "上滚" {
		k0 = "up"
		// 鼠标滚轮
		wheelkey(k0, k1n)
	} else if k0 == "下滚" {
		k0 = "down"
		wheelkey(k0, k1n)
	} else {
		// 正常按键处理（无滚轮操作）
		presskey(k0, k1n)
	}
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

// 根据暂停flag判断是否暂停
func timedelay() {
	for {
		if stop {
			// 如果暂停flag为真，退出施法脚本协程，暂停
			runtime.Goexit()
		} else {
			// 如果暂停flag为假，正常运行脚本
			break
		}
	}
}
