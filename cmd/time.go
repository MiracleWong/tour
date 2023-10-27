package cmd

import (
	"fmt"
	"github.com/MiracleWong/tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Parent Cmd")
	},
}

// 子命令：now 获取当前时间

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取Now时间",
	Long:  "获取当前的时间",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Sub Cmd")
		nowTime := timer.GetNowTime()
		log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

// 子命令：calc 推算时间

var calcTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			// 三种判定的方式
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}

		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatal("timer.GetCalculateTime error : %v", err)
		}
		log.Printf("输出结果：%s, %d", t.Format(layout), t.Unix())
	},
}

// init后，才可以在rootCmd中使用到
func init() {
	fmt.Println("Hello Init")

	// 对相关联的命令进行注册
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calcTimeCmd)
	calcTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或已格式化后的时间")
	calcTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间，有效时间单位为\"ns\",\"us\"(or\"μs\"),\"ms\",\"s\",\"m\",\"h\"")
}
