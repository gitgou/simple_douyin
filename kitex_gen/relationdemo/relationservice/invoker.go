// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	server "github.com/cloudwego/kitex/server"
	relationdemo "github.com/gitgou/simple_douyin/kitex_gen/relationdemo"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler relationdemo.RelationService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
