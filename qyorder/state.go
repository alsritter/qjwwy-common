package qyorder

// 定义订单状态
const (
	StatusCreate    = State(0)
	StatusPrepare   = State(10)
	StatusSuccess   = State(20)
	StatusCancel    = State(30)
	StatusRefund    = State(50)
	StatusRefunding = State(60)
	StatusClose     = State(70)
)

// statusText 定义订单状态文案
var statusText = map[State]string{
	StatusCreate:    "无效",
	StatusPrepare:   "待支付",
	StatusSuccess:   "支付成功",
	StatusCancel:    "已取消",
	StatusRefunding: "退款中",
	StatusRefund:    "已退款",
	StatusClose:     "已关闭",
}

// statusEvent 定义订单状态对应的可操作事件
var statusEvent = map[State][]Event{
	StatusCreate:    {EventInit},
	StatusPrepare:   {EventPay, EventCancel, EventClose},
	StatusSuccess:   {EventRefund},
	StatusCancel:    {EventClose},
	StatusRefunding: {EventIsRefund},
}

func StatusText(status State) string {
	return statusText[status]
}
