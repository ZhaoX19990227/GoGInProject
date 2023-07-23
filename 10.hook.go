package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type MyHook struct {
}

func (hook MyHook) Fire(entry *logrus.Entry) error {
	fmt.Println(entry)
	return nil
}

func (hook MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func main() {
	logrus.AddHook(&MyHook{})
}
