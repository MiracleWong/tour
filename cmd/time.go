package cmd

import (
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
	Long:  "时间转换的工具类",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		// "2006-01-02 15:04:05" 是Format 输出的既定项
		//log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
		log.Printf("输出结果: %s, %d", nowTime.Format(time.RFC3339), nowTime.Unix())
	},
}

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
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}

		calculateTime, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Printf("CalculateTime error : %v", err)
		}
		log.Printf("输出结果：%s, %d", calculateTime.Format(layout), calculateTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calcTimeCmd)

	calcTimeCmd.Flags().StringVarP(&calculateTime, "calculateTime", "c", "", "需要计算的时间")
	calcTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "持续时间")
}
