// Code generated by Kitex v0.4.4. DO NOT EDIT.

package chatservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	chatdemo "github.com/gitgou/simple_douyin/kitex_gen/chatdemo"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return chatServiceServiceInfo
}

var chatServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ChatService"
	handlerType := (*chatdemo.ChatService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetChat":    kitex.NewMethodInfo(getChatHandler, newGetChatArgs, newGetChatResult, false),
		"ChatAction": kitex.NewMethodInfo(chatActionHandler, newChatActionArgs, newChatActionResult, false),
		"Login":      kitex.NewMethodInfo(loginHandler, newLoginArgs, newLoginResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyin",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func getChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chatdemo.ChatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chatdemo.ChatService).GetChat(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetChatArgs:
		success, err := handler.(chatdemo.ChatService).GetChat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetChatResult)
		realResult.Success = success
	}
	return nil
}
func newGetChatArgs() interface{} {
	return &GetChatArgs{}
}

func newGetChatResult() interface{} {
	return &GetChatResult{}
}

type GetChatArgs struct {
	Req *chatdemo.ChatRequest
}

func (p *GetChatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chatdemo.ChatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetChatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetChatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetChatArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetChatArgs) Unmarshal(in []byte) error {
	msg := new(chatdemo.ChatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetChatArgs_Req_DEFAULT *chatdemo.ChatRequest

func (p *GetChatArgs) GetReq() *chatdemo.ChatRequest {
	if !p.IsSetReq() {
		return GetChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetChatResult struct {
	Success *chatdemo.ChatResponse
}

var GetChatResult_Success_DEFAULT *chatdemo.ChatResponse

func (p *GetChatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chatdemo.ChatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetChatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetChatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetChatResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetChatResult) Unmarshal(in []byte) error {
	msg := new(chatdemo.ChatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetChatResult) GetSuccess() *chatdemo.ChatResponse {
	if !p.IsSetSuccess() {
		return GetChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*chatdemo.ChatResponse)
}

func (p *GetChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func chatActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chatdemo.ChatActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chatdemo.ChatService).ChatAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ChatActionArgs:
		success, err := handler.(chatdemo.ChatService).ChatAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChatActionResult)
		realResult.Success = success
	}
	return nil
}
func newChatActionArgs() interface{} {
	return &ChatActionArgs{}
}

func newChatActionResult() interface{} {
	return &ChatActionResult{}
}

type ChatActionArgs struct {
	Req *chatdemo.ChatActionRequest
}

func (p *ChatActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chatdemo.ChatActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChatActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChatActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChatActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChatActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChatActionArgs) Unmarshal(in []byte) error {
	msg := new(chatdemo.ChatActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChatActionArgs_Req_DEFAULT *chatdemo.ChatActionRequest

func (p *ChatActionArgs) GetReq() *chatdemo.ChatActionRequest {
	if !p.IsSetReq() {
		return ChatActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChatActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChatActionResult struct {
	Success *chatdemo.ChatActionResponse
}

var ChatActionResult_Success_DEFAULT *chatdemo.ChatActionResponse

func (p *ChatActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chatdemo.ChatActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChatActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChatActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChatActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChatActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChatActionResult) Unmarshal(in []byte) error {
	msg := new(chatdemo.ChatActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChatActionResult) GetSuccess() *chatdemo.ChatActionResponse {
	if !p.IsSetSuccess() {
		return ChatActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChatActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*chatdemo.ChatActionResponse)
}

func (p *ChatActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(chatdemo.LoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(chatdemo.ChatService).Login(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *LoginArgs:
		success, err := handler.(chatdemo.ChatService).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
	}
	return nil
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *chatdemo.LoginRequest
}

func (p *LoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(chatdemo.LoginRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *LoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *LoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in LoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(chatdemo.LoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *chatdemo.LoginRequest

func (p *LoginArgs) GetReq() *chatdemo.LoginRequest {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

type LoginResult struct {
	Success *chatdemo.LoginResponse
}

var LoginResult_Success_DEFAULT *chatdemo.LoginResponse

func (p *LoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(chatdemo.LoginResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *LoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *LoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in LoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(chatdemo.LoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *chatdemo.LoginResponse {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*chatdemo.LoginResponse)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetChat(ctx context.Context, Req *chatdemo.ChatRequest) (r *chatdemo.ChatResponse, err error) {
	var _args GetChatArgs
	_args.Req = Req
	var _result GetChatResult
	if err = p.c.Call(ctx, "GetChat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ChatAction(ctx context.Context, Req *chatdemo.ChatActionRequest) (r *chatdemo.ChatActionResponse, err error) {
	var _args ChatActionArgs
	_args.Req = Req
	var _result ChatActionResult
	if err = p.c.Call(ctx, "ChatAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *chatdemo.LoginRequest) (r *chatdemo.LoginResponse, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
