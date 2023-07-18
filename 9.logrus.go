package main

import "github.com/sirupsen/logrus"

func main() {

	logrus.SetLevel(logrus.WarnLevel) // 比WarnLevel级别低的就不会再输出了

	logrus.Error("出错了")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("debug")
	logrus.Println("打印")
}
