// Code generated by Kitex v0.4.4. DO NOT EDIT.

package feedservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	videodemo "github.com/gitgou/simple_douyin/kitex_gen/videodemo"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return feedServiceServiceInfo
}

var feedServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FeedService"
	handlerType := (*videodemo.FeedService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":    kitex.NewMethodInfo(feedHandler, newFeedArgs, newFeedResult, false),
		"Publish": kitex.NewMethodInfo(publishHandler, newPublishArgs, newPublishResult, false),
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

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videodemo.FeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videodemo.FeedService).Feed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FeedArgs:
		success, err := handler.(videodemo.FeedService).Feed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FeedResult)
		realResult.Success = success
	}
	return nil
}
func newFeedArgs() interface{} {
	return &FeedArgs{}
}

func newFeedResult() interface{} {
	return &FeedResult{}
}

type FeedArgs struct {
	Req *videodemo.FeedRequest
}

func (p *FeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(videodemo.FeedRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FeedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FeedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FeedArgs) Unmarshal(in []byte) error {
	msg := new(videodemo.FeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FeedArgs_Req_DEFAULT *videodemo.FeedRequest

func (p *FeedArgs) GetReq() *videodemo.FeedRequest {
	if !p.IsSetReq() {
		return FeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FeedArgs) IsSetReq() bool {
	return p.Req != nil
}

type FeedResult struct {
	Success *videodemo.FeedResponse
}

var FeedResult_Success_DEFAULT *videodemo.FeedResponse

func (p *FeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(videodemo.FeedResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FeedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FeedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FeedResult) Unmarshal(in []byte) error {
	msg := new(videodemo.FeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FeedResult) GetSuccess() *videodemo.FeedResponse {
	if !p.IsSetSuccess() {
		return FeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*videodemo.FeedResponse)
}

func (p *FeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(videodemo.PublishRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(videodemo.FeedService).Publish(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishArgs:
		success, err := handler.(videodemo.FeedService).Publish(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishResult)
		realResult.Success = success
	}
	return nil
}
func newPublishArgs() interface{} {
	return &PublishArgs{}
}

func newPublishResult() interface{} {
	return &PublishResult{}
}

type PublishArgs struct {
	Req *videodemo.PublishRequest
}

func (p *PublishArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(videodemo.PublishRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PublishArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PublishArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PublishArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishArgs) Unmarshal(in []byte) error {
	msg := new(videodemo.PublishRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishArgs_Req_DEFAULT *videodemo.PublishRequest

func (p *PublishArgs) GetReq() *videodemo.PublishRequest {
	if !p.IsSetReq() {
		return PublishArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishResult struct {
	Success *videodemo.PublishResponse
}

var PublishResult_Success_DEFAULT *videodemo.PublishResponse

func (p *PublishResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(videodemo.PublishResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PublishResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PublishResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PublishResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishResult) Unmarshal(in []byte) error {
	msg := new(videodemo.PublishResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishResult) GetSuccess() *videodemo.PublishResponse {
	if !p.IsSetSuccess() {
		return PublishResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishResult) SetSuccess(x interface{}) {
	p.Success = x.(*videodemo.PublishResponse)
}

func (p *PublishResult) IsSetSuccess() bool {
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

func (p *kClient) Feed(ctx context.Context, Req *videodemo.FeedRequest) (r *videodemo.FeedResponse, err error) {
	var _args FeedArgs
	_args.Req = Req
	var _result FeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Publish(ctx context.Context, Req *videodemo.PublishRequest) (r *videodemo.PublishResponse, err error) {
	var _args PublishArgs
	_args.Req = Req
	var _result PublishResult
	if err = p.c.Call(ctx, "Publish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}