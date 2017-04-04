package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type ConfigHandle struct {
	beego.Controller
}

func (this *ConfigHandle) Get() {
	var (
		out  LotteryConfigOut
		list []*LotteryConfig
		zst  []*LotteryZst
		zst1 []*LotteryZst
	)

	zst = append(zst, &LotteryZst{Zid: 1, Name: "湖北快三走势图一", Sname: "走势图一", Url: "http://abc.com/1.htm", Picture: "http://abc.com/1.png", Isdefault: 1, State: 0})
	zst = append(zst, &LotteryZst{Zid: 2, Name: "湖北快三走势图二", Sname: "走势图二", Url: "http://abc.com/2.htm", Picture: "http://abc.com/2.png", Isdefault: 0, State: 0})

	zst1 = append(zst1, &LotteryZst{Zid: 3, Name: "福彩3D走势图一", Sname: "走势图一", Url: "http://abc.com/1.htm", Picture: "http://abc.com/1.png", Isdefault: 1, State: 0})
	zst1 = append(zst1, &LotteryZst{Zid: 4, Name: "福彩3D走势图二", Sname: "走势图二", Url: "http://abc.com/2.htm", Picture: "http://abc.com/2.png", Isdefault: 0, State: 0})

	list = append(list, &LotteryConfig{Cid: 30, Cname: "福彩湖北快3", Zst: zst})
	list = append(list, &LotteryConfig{Cid: 1, Cname: "福彩3D", Zst: zst1})

	out.Err = 0
	out.Msg = ""
	out.Servertime = time.Now().Unix()
	out.Version = "0.0.1"
	out.Data = list

	this.Data["json"] = &out
	this.ServeJSON()

}

type LotteryConfigOut struct {
	Err        int64            `json:"err"`
	Msg        string           `json:"msg"`
	Servertime int64            `json:"servertime"`
	Version    string           `json:"version"`
	Data       []*LotteryConfig `json:"data"`
}

type LotteryConfig struct {
	Cid     int64         `json:"cid"`
	Cname   string        `json:"cname"`
	Default int64         `json:"default"`
	Zst     []*LotteryZst `json:"zst"`
}

type LotteryZst struct {
	Zid       int64  `json:"zid"`
	Name      string `json:"name"`
	Sname     string `json:"sname"`
	Url       string `json:"url"`
	Picture   string `json:"picture"`
	Isdefault int64  `json:"isdefault"`
	State     int64  `json:"state"`
}
