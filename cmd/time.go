package cmd

import (
	"fmt"
	"github.com/MiracleWong/tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Var Create")
	},
}

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

// init后，才可以在rootCmd中使用到
func init() {
	fmt.Println("Hello Init")
	timeCmd.AddCommand(nowTimeCmd)
}
