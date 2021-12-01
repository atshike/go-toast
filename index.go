package main

import (
	"github.com/go-toast/toast"
	"log"
)

func main() {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "写日报",
		Message: "大佬别忘了写日报！",
		Icon:    "D:\\Users\\dell\\lg.jpg", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "拒绝", "#"},
			{"protocol", "收到", "#"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
