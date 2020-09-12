package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
	"tour/internal/timer"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use: "time",
	Short: "时间格式处理",
	Long: "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use: "now",
	Short: "获得当前时间",
	Long: "获得当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var locationTimeCmd = &cobra.Command{
	Use: "location",
	Short: "时区设置的用法",
	Long: "举例说明在解析时间字符串时，指定时区后，会以指定时区的标准来对待字符串",
	Run: func(cmd *cobra.Command, args []string) {
		location, _ := time.LoadLocation("Asia/Shanghai")
		inputTime := "2020-09-12 12:30:45"
		layout := "2006-01-02 15:04:05"

		// parse时间字符串的时候，指定时区，当做指定时区的时间字符串来解析
		t, _ := time.ParseInLocation(layout, inputTime, location)
		dateTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
		log.Printf("输入时间(指定东8区)：%s，输出时间(转换到东8区时间)：%s, unix时间: %d", inputTime, dateTime, t.Unix())

		// parse时间字符串的时候，不指定时区， 被当做UTC时区的时间字符串来解析
		t1, _ := time.Parse(layout, inputTime)
		dateTime1 := time.Unix(t1.Unix(), 0).In(location).Format(layout)
		log.Printf("输入时间(不指定时区)：%s，输出时间(转换到东8区时间)：%s, unix时间戳(秒数): %d", inputTime, dateTime1, t1.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use: "calc",
	Short: "计算所需时间",
	Long: "计算所需时间",
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
			log.Fatalf("timer.GetCalculateTime err: %#v", err)
		}
		log.Printf("输出结果：%s, %d", calculateTime.Format(layout), calculateTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	timeCmd.AddCommand(locationTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "",
		"需要计算的时间，有效单位为时间戳或已格式化后的时间字符串")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "",
		`持续时间，有效时间单位为"ns", "us", "ms", "s", "m", "h"`)
}
