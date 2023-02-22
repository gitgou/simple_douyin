// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	userdemo "github.com/gitgou/simple_douyin/kitex_gen/userdemo"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*userdemo.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetUser":         kitex.NewMethodInfo(getUserHandler, newGetUserArgs, newGetUserResult, false),
		"MGetUser":        kitex.NewMethodInfo(mGetUserHandler, newMGetUserArgs, newMGetUserResult, false),
		"CreateUser":      kitex.NewMethodInfo(createUserHandler, newCreateUserArgs, newCreateUserResult, false),
		"Login":           kitex.NewMethodInfo(loginHandler, newLoginArgs, newLoginResult, false),
		"CheckUserOnline": kitex.NewMethodInfo(checkUserOnlineHandler, newCheckUserOnlineArgs, newCheckUserOnlineResult, false),
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

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userdemo.GetUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userdemo.UserService).GetUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserArgs:
		success, err := handler.(userdemo.UserService).GetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserArgs() interface{} {
	return &GetUserArgs{}
}

func newGetUserResult() interface{} {
	return &GetUserResult{}
}

type GetUserArgs struct {
	Req *userdemo.GetUserRequest
}

func (p *GetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserArgs) Unmarshal(in []byte) error {
	msg := new(userdemo.GetUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserArgs_Req_DEFAULT *userdemo.GetUserRequest

func (p *GetUserArgs) GetReq() *userdemo.GetUserRequest {
	if !p.IsSetReq() {
		return GetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserResult struct {
	Success *userdemo.GetUserResponse
}

var GetUserResult_Success_DEFAULT *userdemo.GetUserResponse

func (p *GetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserResult) Unmarshal(in []byte) error {
	msg := new(userdemo.GetUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserResult) GetSuccess() *userdemo.GetUserResponse {
	if !p.IsSetSuccess() {
		return GetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userdemo.GetUserResponse)
}

func (p *GetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func mGetUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userdemo.MGetUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userdemo.UserService).MGetUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MGetUserArgs:
		success, err := handler.(userdemo.UserService).MGetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MGetUserResult)
		realResult.Success = success
	}
	return nil
}
func newMGetUserArgs() interface{} {
	return &MGetUserArgs{}
}

func newMGetUserResult() interface{} {
	return &MGetUserResult{}
}

type MGetUserArgs struct {
	Req *userdemo.MGetUserRequest
}

func (p *MGetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MGetUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MGetUserArgs) Unmarshal(in []byte) error {
	msg := new(userdemo.MGetUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MGetUserArgs_Req_DEFAULT *userdemo.MGetUserRequest

func (p *MGetUserArgs) GetReq() *userdemo.MGetUserRequest {
	if !p.IsSetReq() {
		return MGetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MGetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type MGetUserResult struct {
	Success *userdemo.MGetUserResponse
}

var MGetUserResult_Success_DEFAULT *userdemo.MGetUserResponse

func (p *MGetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MGetUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MGetUserResult) Unmarshal(in []byte) error {
	msg := new(userdemo.MGetUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MGetUserResult) GetSuccess() *userdemo.MGetUserResponse {
	if !p.IsSetSuccess() {
		return MGetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MGetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userdemo.MGetUserResponse)
}

func (p *MGetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userdemo.CreateUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userdemo.UserService).CreateUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateUserArgs:
		success, err := handler.(userdemo.UserService).CreateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateUserResult)
		realResult.Success = success
	}
	return nil
}
func newCreateUserArgs() interface{} {
	return &CreateUserArgs{}
}

func newCreateUserResult() interface{} {
	return &CreateUserResult{}
}

type CreateUserArgs struct {
	Req *userdemo.CreateUserRequest
}

func (p *CreateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateUserArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateUserArgs) Unmarshal(in []byte) error {
	msg := new(userdemo.CreateUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateUserArgs_Req_DEFAULT *userdemo.CreateUserRequest

func (p *CreateUserArgs) GetReq() *userdemo.CreateUserRequest {
	if !p.IsSetReq() {
		return CreateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateUserResult struct {
	Success *userdemo.CreateUserResponse
}

var CreateUserResult_Success_DEFAULT *userdemo.CreateUserResponse

func (p *CreateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateUserResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateUserResult) Unmarshal(in []byte) error {
	msg := new(userdemo.CreateUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateUserResult) GetSuccess() *userdemo.CreateUserResponse {
	if !p.IsSetSuccess() {
		return CreateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*userdemo.CreateUserResponse)
}

func (p *CreateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userdemo.LoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userdemo.UserService).Login(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *LoginArgs:
		success, err := handler.(userdemo.UserService).Login(ctx, s.Req)
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
	Req *userdemo.LoginRequest
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in LoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(userdemo.LoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *userdemo.LoginRequest

func (p *LoginArgs) GetReq() *userdemo.LoginRequest {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

type LoginResult struct {
	Success *userdemo.LoginResponse
}

var LoginResult_Success_DEFAULT *userdemo.LoginResponse

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in LoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(userdemo.LoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *userdemo.LoginResponse {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*userdemo.LoginResponse)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkUserOnlineHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(userdemo.CheckUserOnlineRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(userdemo.UserService).CheckUserOnline(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckUserOnlineArgs:
		success, err := handler.(userdemo.UserService).CheckUserOnline(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckUserOnlineResult)
		realResult.Success = success
	}
	return nil
}
func newCheckUserOnlineArgs() interface{} {
	return &CheckUserOnlineArgs{}
}

func newCheckUserOnlineResult() interface{} {
	return &CheckUserOnlineResult{}
}

type CheckUserOnlineArgs struct {
	Req *userdemo.CheckUserOnlineRequest
}

func (p *CheckUserOnlineArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckUserOnlineArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckUserOnlineArgs) Unmarshal(in []byte) error {
	msg := new(userdemo.CheckUserOnlineRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckUserOnlineArgs_Req_DEFAULT *userdemo.CheckUserOnlineRequest

func (p *CheckUserOnlineArgs) GetReq() *userdemo.CheckUserOnlineRequest {
	if !p.IsSetReq() {
		return CheckUserOnlineArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckUserOnlineArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckUserOnlineResult struct {
	Success *userdemo.CheckUserOnlineResponse
}

var CheckUserOnlineResult_Success_DEFAULT *userdemo.CheckUserOnlineResponse

func (p *CheckUserOnlineResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckUserOnlineResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckUserOnlineResult) Unmarshal(in []byte) error {
	msg := new(userdemo.CheckUserOnlineResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckUserOnlineResult) GetSuccess() *userdemo.CheckUserOnlineResponse {
	if !p.IsSetSuccess() {
		return CheckUserOnlineResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckUserOnlineResult) SetSuccess(x interface{}) {
	p.Success = x.(*userdemo.CheckUserOnlineResponse)
}

func (p *CheckUserOnlineResult) IsSetSuccess() bool {
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

func (p *kClient) GetUser(ctx context.Context, Req *userdemo.GetUserRequest) (r *userdemo.GetUserResponse, err error) {
	var _args GetUserArgs
	_args.Req = Req
	var _result GetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetUser(ctx context.Context, Req *userdemo.MGetUserRequest) (r *userdemo.MGetUserResponse, err error) {
	var _args MGetUserArgs
	_args.Req = Req
	var _result MGetUserResult
	if err = p.c.Call(ctx, "MGetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUser(ctx context.Context, Req *userdemo.CreateUserRequest) (r *userdemo.CreateUserResponse, err error) {
	var _args CreateUserArgs
	_args.Req = Req
	var _result CreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *userdemo.LoginRequest) (r *userdemo.LoginResponse, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CheckUserOnline(ctx context.Context, Req *userdemo.CheckUserOnlineRequest) (r *userdemo.CheckUserOnlineResponse, err error) {
	var _args CheckUserOnlineArgs
	_args.Req = Req
	var _result CheckUserOnlineResult
	if err = p.c.Call(ctx, "CheckUserOnline", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
