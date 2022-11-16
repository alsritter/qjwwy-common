package qyorder

import "fmt"

var (
	// handlerInit 初始化订单
	handlerInit = Handler(func(opt *Opt) (State, error) {
		fmt.Println(fmt.Sprintf("正在处理创建订单逻辑，订单ID(%d), 订单名称(%s) ... 处理完毕！", opt.OrderId, opt.OrderName))
		return StatusPrepare, nil
	})

	// handlerCancel 关闭订单
	handlerCancel = Handler(func(opt *Opt) (State, error) {
		return StatusCancel, nil
	})

	// handlerClose 关闭订单
	handlerClose = Handler(func(opt *Opt) (State, error) {
		return StatusClose, nil
	})

	// handlerPay 支付订单
	handlerPay = Handler(func(opt *Opt) (State, error) {
		if opt.HandlerSendSMS != nil {
			_ = opt.HandlerSendSMS("18888888888", "恭喜你预定成功了！")
		}
		return StatusSuccess, nil
	})

	// handlerRefund 退款订单
	handlerRefund = Handler(func(opt *Opt) (State, error) {
		return StatusRefunding, nil
	})

	// handlerRefund 退款完成订单
	handlerIsRefund = Handler(func(opt *Opt) (State, error) {
		return StatusRefund, nil
	})
)
