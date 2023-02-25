package utils

import (
	"os/user"
	"strings"
)

const (
	// 测试用
	Janson    = "ou_6e75a4323ad0eef42ae349d71203969e"
	JansonUID = "d14ed441"
	Fish      = "ou_9a1afb986757fa17ee0d4d6ba2289d72"
	FishUID   = "5efg94ff"
	Raft      = "ou_bcd4e27aad8d4cda70bb0d0b091e487c"
	RaftUID   = "65491367"
	All       = "all"
)

const (
	BotOpenID               = "ou_35ac11238b92d3ab3ad98547d3c7fd47" // 机器人openid
	StudentOpenDepartmentID = "od-1aad208d5b2b4b4fd3dd8ee26d8abce6" // 学生部门id
	StudentDepartmentID     = "79dc8e4d9bd8d4db"
	MentorDepartmentID      = "1"
	MentorOpenDepartmentID  = "od-00f02bd1248978fec7311904fba21f01" // 导师团部门id
	BotEncryptedKey         = "NPHlL4eOQV5i8Hn7IHxWsg07J7Ev6LIj"
	BotVerifyToken          = "Nu7hthyTijDXvVZyO6uZ2eFBmH2DEBCc"
)

type MemberType string

const (
	OpenIDType MemberType = "open_id"
	UserIDType MemberType = "user_id"
)

func GetWhiteOpenIDs() []string {
	if IsTestEnv() {
		return testWhiteOpenIDs
	}
	return whiteOpenIDs
}

func Get1V1MentorIDs() []string {
	if IsTestEnv() {
		return []string{
			"ou_6e75a4323ad0eef42ae349d71203969e",
			"ou_bcd4e27aad8d4cda70bb0d0b091e487c", //虎哥
			"ou_9a1afb986757fa17ee0d4d6ba2289d72", //牛哥
		}
	}
	return whiteOpenIDs
}

func GetWorkshopPrepareChatMentors() []string {
	return Get1V1MentorIDs()
}

//todo: whitelist 可配置
//OpenID:同一个 User ID 在不同应用中的 Open ID 不同
var testWhiteOpenIDs = []string{
	//"ou_6e75a4323ad0eef42ae349d71203969e", //狼哥
	//"ou_bcd4e27aad8d4cda70bb0d0b091e487c", //虎哥
	//"ou_9a1afb986757fa17ee0d4d6ba2289d72", //牛哥
}
var whiteOpenIDs = []string{
	"ou_6e75a4323ad0eef42ae349d71203969e", //狼哥
	"ou_bcd4e27aad8d4cda70bb0d0b091e487c", //虎哥
	"ou_9a1afb986757fa17ee0d4d6ba2289d72", //牛哥
	"ou_844198543d625613bb5f9a5e4f2366c2", //鹏哥
}

func GetTaskFocusIDs() []string {
	return focusOpenIDs
}

//todo:根据部门获取
var focusOpenIDs = []string{
	"ou_6e75a4323ad0eef42ae349d71203969e", //狼哥
	"ou_bcd4e27aad8d4cda70bb0d0b091e487c", //虎哥
	"ou_9a1afb986757fa17ee0d4d6ba2289d72", //牛哥
}

const (
	testProduceChatID = "oc_e36d348632effb1192b59c5b27f3c67e" // 导师群测试
	produceChatID     = "oc_3cb3aa4021eee68165c78e731476d3c1" // 导师团
	testTopicChatID   = "oc_c6f4a213c3f2b35bba6e17868af7c072" // 目标话题测试
	topicChatID       = "oc_9ab398ed766e74cc68706c1919b48471" // 系统设计workshop
	mainGroupChatID   = "oc_757c35079993acdb60fc7e668f99050b" // 训练营大群
)

// GetRemindChatWhiteIDs 发言提醒白名单
func GetRemindChatWhiteIDs() map[string]struct{} {
	return map[string]struct{}{
		produceChatID:   {},
		topicChatID:     {},
		mainGroupChatID: {},
	}
}

//GetTestEnvRemindAllowList 发言提醒测试仅发名单
func GetTestEnvRemindAllowList() map[string]struct{} {
	return map[string]struct{}{
		testProduceChatID: {},
	}
}

// GetRemindPostChatID 测试提醒消息用
func GetRemindPostChatID() string {
	if IsTestEnv() {
		return testProduceChatID
	}
	return testProduceChatID
}

// GetJoinMainGroupChatID 测试进群消息用
func GetJoinMainGroupChatID() string {
	if IsTestEnv() {
		return testProduceChatID
	}
	return mainGroupChatID
}

// GetWorkshopProduceChatID workshop发题名单
func GetWorkshopProduceChatID() string {
	if IsTestEnv() {
		return testProduceChatID
	}
	return produceChatID
}

// GetWorkshopTopicGroupChatID 获取接受workshop任务群id
func GetWorkshopTopicGroupChatID() string {
	if IsTestEnv() {
		return testTopicChatID
	}
	return topicChatID
}

func IsTestEnv() bool {
	result := false
	u, _ := user.Current()
	if strings.Contains(u.Username, "bytedance") {
		result = true
	}
	return result
}
