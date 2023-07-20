package time

import (
	"fmt"
	"time"
)

// 时间类型
func TimeDemo() {

	now := time.Now() //获取当前时间
	fmt.Println(now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间戳 时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。
func TimeStampDemo() {
	now := time.Now()            //获取当前时间
	timeStamp1 := now.Unix()     //时间戳
	timeStamp2 := now.UnixNano() //纳秒时间戳
	fmt.Println(timeStamp1)
	fmt.Println(timeStamp2)

	//使用time.Unix将时间戳转换为时间格式
	timeobj := time.Unix(timeStamp1, 0)
	fmt.Println(timeobj)

}

// 时间间隔，ime.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
func TimeDruationDemo() {
	//ADD函数,时间+时间间隔
	now := time.Now()
	laterTime := now.Add(60 * time.Minute) //当前时间加上60分钟后的时间
	fmt.Println(laterTime)
	//Sub函数,求两个时间之间的差值：
	subtime := now.Sub(now.Add(1 * time.Hour))
	fmt.Println(subtime)
	//equal函数,判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。
	b := now.Equal(now.Add(1 * time.Hour))
	fmt.Println(b)
	//Before函数,如果t代表的时间点在u之前，返回真；否则返回假。
	b = now.Before(now.Add(10 * time.Second))
	fmt.Println(b)
	//after函数,如果t代表的时间点在u之后，返回真；否则返回假。
	b = now.After(now.Add(10 * time.Second))
	fmt.Println(b)

}

// 定时器函数，每隔一段时间会往管道发送数据
func TickDemo() {
	ticker := time.Tick(2 * time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}

// 时间格式化
func FormatDemo() {
	now := time.Now()
	/* // 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02")) */

	//解析字符串格式的时间
	fmt.Println(now)
	//加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//按照指定时区和格式解析字符串时间
	timeobj, err := time.ParseInLocation("2006-1-2 15:04:05", "1998-11-02 19:08:10", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeobj)
	fmt.Println(timeobj.Sub(now))
}
