package cmd

import (
	"github.com/MiracleWong/tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var nowTimeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:06"), nowTime.Unix())
		//log.Printf("输出结果：%s, %d", nowTime, nowTime.Unix())
	},
}

var calculateTime string
var duration string

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:06"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64((t)), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	nowTimeCmd.AddCommand(nowTimeCmd)
	nowTimeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间，有效单位为时间戳或已格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}
