package cmd

import (
    "github.com/MiracleWong/tour/internal/timer"
    "github.com/spf13/cobra"
    "log"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
        nowTime := timer.GetNowTime()
        log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:06"), nowTime.Unix())
        //log.Printf("输出结果：%s, %d", nowTime, nowTime.Unix())
	},
}

func init() {

}
