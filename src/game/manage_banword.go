package game

import (
	"fmt"
	"regexp"
)

// 违禁词库

var manageBanWord *ManageBanWord

type ManageBanWord struct {
	BanWordBase  []string
	BanWordExtra []string // 外部扩展
}

// 单例
func GetManageBanWord() *ManageBanWord {
	if manageBanWord == nil {
		manageBanWord = new(ManageBanWord)
		manageBanWord.BanWordBase = []string{"外挂", "工具"}
		manageBanWord.BanWordExtra = []string{"TMD"}
	}

	return manageBanWord
}

func (self *ManageBanWord) IsBanWord(txt string) bool {
	// 基础词库查找
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}

	for _, v := range self.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}

	return false
}
