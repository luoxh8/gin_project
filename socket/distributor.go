package socket

/**
	获得一个connection对象，获取内部的（）分发到controller里面处理业务
 */

type Distributor struct {
	conn *Connection
}

func initDistributor() *Distributor {
	var dis *Distributor
	return dis
}
