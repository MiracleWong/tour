package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "显示时间，转换时间",
	Long:  "时间转换的工具类",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.Now())
	},
}
