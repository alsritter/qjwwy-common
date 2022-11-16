package qyorder

// 定义订单事件
const (
	EventInit     = Event("初始化")
	EventPay      = Event("支付")
	EventCancel   = Event("取消")
	EventClose    = Event("关闭")
	EventRefund   = Event("退款")
	EventIsRefund = Event("已退款")
)

// 定义订单事件对应的处理方法
var eventHandler = map[Event]Handler{
	EventInit:     handlerInit,
	EventCancel:   handlerCancel,
	EventClose:    handlerClose,
	EventPay:      handlerPay,
	EventRefund:   handlerRefund,
	EventIsRefund: handlerIsRefund,
}
