package cmd

import (
	"github.com/MiracleWong/tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"time"
)

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

func init() {
	timeCmd.AddCommand(nowTimeCmd)
}
