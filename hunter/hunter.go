package main

import (
	"context"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log/slog"
	"os"
	"runtime"
	"time"
)

// 脚本循环的频率
const FREQUENCY int64 = 10

var stop bool

// 生成可以取消的 context
var ctx, cancel = context.WithCancel(context.Background())
var 目标数量 string = "AOE"
var 数量 string = "AOE"

const (
	多重射击 int64 = 1000 + iota
	倒刺射击
	杀戮命令
	狂野怒火
	死亡飞轮
	夺命射击
	束缚射击
	荒野的召唤
	眼镜蛇射击
	意气风发
	饰品药水
	大红瓶
	治疗石
	打断
	停止施法
	停止攻击
)

// 技能和按键映射
var SpellKeyMap = map[int64][]string{
	//以下按键中不能有快捷键对应按键，比如：不能有"4","6","0","7","8","9"！！！
	多重射击:  []string{"l"},
	束缚射击:  []string{"l", "ctrl"},
	倒刺射击:  []string{"j"},
	杀戮命令:  []string{"k"},
	狂野怒火:  []string{"y", "shift"},
	死亡飞轮:  []string{"l", "shift"},
	夺命射击:  []string{"k", "ctrl"},
	荒野的召唤: []string{"u", "shift"},
	眼镜蛇射击: []string{"o"},
	意气风发:  []string{","},
	饰品药水:  []string{"[", "ctrl", "shift"},
	大红瓶:   []string{";"},
	治疗石:   []string{"'"},
	打断:    []string{"."},
	停止施法:  []string{"y", "ctrl"},
	停止攻击:  []string{"u", "alt"},
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
			//以下是快捷键，不能与施放技能的按键相同！！！
			//按快捷键“4”开始脚本
			if ev.Rawcode == 52 {
				stop = false
				目标数量 = 数量
				go start()
			}
			//按快捷键“6”暂停脚本
			if ev.Rawcode == 54 {
				stop = true
				cast(停止施法)
				cast(停止攻击)
				if 目标数量 == "单体" || 目标数量 == "AOE" {
					数量 = 目标数量
				}
			}
			//按快捷键小键盘“-”停止脚本软件
			if ev.Rawcode == 109 {
				os.Exit(0)
			}
			//按快捷键“0”为打断
			if ev.Rawcode == 48 {
				数量 = 目标数量
				目标数量 = "打断"
			}
			//按快捷键“8”为单体
			if ev.Rawcode == 56 {
				目标数量 = "单体"
			}
			//按快捷键“7”为爆发
			if ev.Rawcode == 55 {
				数量 = 目标数量
				目标数量 = "爆发"
			}
			//按快捷键“9”为AOE
			if ev.Rawcode == 57 {
				目标数量 = "AOE"
			}
			//按快捷键小键盘"1"大红瓶
			if ev.Rawcode == 49 {
				cast(大红瓶)
			}
			//按快捷键小键盘"2"术士治疗石35
			if ev.Rawcode == 50 {
				cast(治疗石)
			}
			//按快捷键小键盘"3"加血34
			if ev.Rawcode == 51 {
				cast(意气风发)
			}
		}
	}
}

// 脚本开始
func start() {
	for {
		timedelay()
		//脚本开始时间
		//timestart := time.Now()
		switch 目标数量 {
		case "打断":
			目标数量 = 数量
			cast(停止施法)
			cast(停止攻击)
			cast(打断)
		case "单体":
			cast(死亡飞轮)
			cast(狂野怒火)
			cast(倒刺射击)
			cast(杀戮命令)
			cast(眼镜蛇射击)
			cast(夺命射击)
		case "AOE":
			cast(死亡飞轮)
			cast(狂野怒火)
			cast(多重射击)
			cast(倒刺射击)
			cast(杀戮命令)
			cast(夺命射击)
		case "爆发":
			目标数量 = 数量
			cast(停止施法)
			cast(停止攻击)
			cast(饰品药水)
			time.Sleep(time.Second)
			cast(荒野的召唤)

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
		//按下按键
		robotgo.KeyToggle(v, "down")
	}
	//滚动鼠标
	robotgo.ScrollDir(1, k0)
	for _, v := range k1n {
		//抬起按键
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

func delay(timestart, timeend time.Time) {
	timelong := timeend.Sub(timestart).Microseconds()
	//脚本每次循环的周期
	var cycle int64 = 1000000 / FREQUENCY
	slog.Info("周期=", cycle)
	delaytime := cycle - timelong
	slog.Info("delaytime=", delaytime)
	time.Sleep(time.Duration(delaytime * 1000))
	slog.Info("Sleep:", time.Duration(delaytime*1000))
}

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
