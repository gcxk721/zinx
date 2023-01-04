// Package ziface 主要提供zinx全部抽象层接口定义.
// 包括:
//
//			IServer 服务mod接口
//			IRouter 路由mod接口
//			IConnection 连接mod层接口
//	     IMessage 消息mod接口
//			IDataPack 消息拆解接口
//	     IMsgHandler 消息处理及协程池接口
//
// 当前文件描述:
// @Title  irequest.go
// @Description  提供连接请求全部接口声明
// @Author  Aceld - Thu Mar 11 10:32:29 CST 2019
package ziface

/*
IRequest 接口：
实际上是把客户端请求的链接信息 和 请求的数据 包装到了 Request里
*/
type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetData() []byte            //获取请求消息的数据

	BindRouter(router IRouter) //绑定这次请求由哪个路由处理
	Next()                     //转进到下一个处理器开始执行 但是调用此方法的函数会根据先后顺序逆序执行
	Abort()                    //终止处理函数的运行 但调用此方法的函数会执行完毕
}
