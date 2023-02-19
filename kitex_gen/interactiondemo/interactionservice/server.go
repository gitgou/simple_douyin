// Code generated by Kitex v0.4.4. DO NOT EDIT.
package interaction

import (
	server "github.com/cloudwego/kitex/server"
	interaction "simple_douyin-master/cmd/interaction/kitex_gen/interaction"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler interaction.Interaction, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
