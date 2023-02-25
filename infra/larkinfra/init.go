package larkinfra

import (
	"github.com/larksuite/oapi-sdk-go/v3"
	"math/rand"
	"time"
)

const (
	appid     = "cli_a31ac04fbef8900b"
	appsecret = "83BUHQBSPLm136PftJXhcBBrPUGyqv1I"
)

var Client *lark.Client

func init() {
	Client = lark.NewClient(appid, appsecret, lark.WithLogReqAtDebug(true),
		lark.WithEnableTokenCache(true))
	rand.Seed(time.Now().Unix())
}
