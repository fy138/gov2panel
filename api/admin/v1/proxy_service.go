package v1

import (
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"

	"github.com/gogf/gf/v2/frame/g"
)

type ProxyServiceReq struct {
	g.Meta `path:"/service" tags:"ProxyService" method:"get,post" summary:"服务器管理"`
	SortOrder
	OffsetLimit
	entity.V2ProxyService
}
type ProxyServiceRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.ProxyServiceInfo `json:"data"`
	Totle  int                       `json:"totle"`
}

type ProxyServiceAEReq struct {
	g.Meta `path:"/service/ae" tags:"ProxyService" method:"post" summary:"服务器管理AE"`
	entity.V2ProxyService
}
type ProxyServiceAERes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type ProxyServiceDelReq struct {
	g.Meta `path:"/service/del" tags:"ProxyService" method:"post" summary:"服务器管理删除"`
	Ids    []int `json:"ids"`
}
type ProxyServiceDelRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type ProxyServiceFlowReq struct {
	g.Meta `path:"/service/flow" tags:"ProxyService" method:"post" summary:"获取服务器流量使用排行榜"`
	Date   string `json:"date"`
}
type ProxyServiceFlowRes struct {
	ServiceFlowTop []*model.ProxyServiceFlow `json:"service_flow_top"`
}

type OnlineUserCountAndLastPushAtReq struct {
	g.Meta `path:"/service/online_user_count_and_last_push_at" tags:"ProxyService" method:"post" summary:"获取所有服务器当前在线用户数量和服务器最后提交时间"`
}
type OnlineUserCountAndLastPushAtRes struct {
	Data map[int]map[int]int64 `json:"data"`
}
