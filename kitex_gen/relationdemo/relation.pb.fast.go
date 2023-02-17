// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package relationdemo

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	userdemo "github.com/gitgou/simple_douyin/kitex_gen/userdemo"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *RelationRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationRequest[number], err)
}

func (x *RelationRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ActionType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RelationResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RelationResponse[number], err)
}

func (x *RelationResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v userdemo.BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *GetFollowRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowRequest[number], err)
}

func (x *GetFollowRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetFollowResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowResponse[number], err)
}

func (x *GetFollowResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v userdemo.BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *GetFollowResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v userdemo.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserList = append(x.UserList, &v)
	return offset, nil
}

func (x *GetFollowerRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowerRequest[number], err)
}

func (x *GetFollowerRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetFollowerResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowerResponse[number], err)
}

func (x *GetFollowerResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v userdemo.BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *GetFollowerResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v userdemo.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserList = append(x.UserList, &v)
	return offset, nil
}

func (x *FriendUser) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FriendUser[number], err)
}

func (x *FriendUser) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v userdemo.User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *FriendUser) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.MsgType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *FriendUser) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Message, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetFriendRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFriendRequest[number], err)
}

func (x *GetFriendRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetFriendResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFriendResponse[number], err)
}

func (x *GetFriendResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v userdemo.BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *GetFriendResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v FriendUser
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserList = append(x.UserList, &v)
	return offset, nil
}

func (x *CheckFollowRelationRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckFollowRelationRequest[number], err)
}

func (x *CheckFollowRelationRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CheckFollowRelationRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ToUserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CheckFollowRelationRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.RelationType, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CheckFollowRelationResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckFollowRelationResponse[number], err)
}

func (x *CheckFollowRelationResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Follow, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *RelationRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *RelationRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *RelationRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ToUserId)
	return offset
}

func (x *RelationRequest) fastWriteField3(buf []byte) (offset int) {
	if x.ActionType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.ActionType)
	return offset
}

func (x *RelationResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *RelationResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *GetFollowRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetFollowRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *GetFollowResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetFollowResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *GetFollowResponse) fastWriteField2(buf []byte) (offset int) {
	if x.UserList == nil {
		return offset
	}
	for i := range x.UserList {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.UserList[i])
	}
	return offset
}

func (x *GetFollowerRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetFollowerRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *GetFollowerResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetFollowerResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *GetFollowerResponse) fastWriteField2(buf []byte) (offset int) {
	if x.UserList == nil {
		return offset
	}
	for i := range x.UserList {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.UserList[i])
	}
	return offset
}

func (x *FriendUser) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *FriendUser) fastWriteField1(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.User)
	return offset
}

func (x *FriendUser) fastWriteField2(buf []byte) (offset int) {
	if x.MsgType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.MsgType)
	return offset
}

func (x *FriendUser) fastWriteField3(buf []byte) (offset int) {
	if x.Message == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.Message)
	return offset
}

func (x *GetFriendRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetFriendRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *GetFriendResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetFriendResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *GetFriendResponse) fastWriteField2(buf []byte) (offset int) {
	if x.UserList == nil {
		return offset
	}
	for i := range x.UserList {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.UserList[i])
	}
	return offset
}

func (x *CheckFollowRelationRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CheckFollowRelationRequest) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.UserId)
	return offset
}

func (x *CheckFollowRelationRequest) fastWriteField2(buf []byte) (offset int) {
	if x.ToUserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.ToUserId)
	return offset
}

func (x *CheckFollowRelationRequest) fastWriteField3(buf []byte) (offset int) {
	if x.RelationType == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.RelationType)
	return offset
}

func (x *CheckFollowRelationResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CheckFollowRelationResponse) fastWriteField1(buf []byte) (offset int) {
	if !x.Follow {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.Follow)
	return offset
}

func (x *RelationRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *RelationRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *RelationRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ToUserId)
	return n
}

func (x *RelationRequest) sizeField3() (n int) {
	if x.ActionType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.ActionType)
	return n
}

func (x *RelationResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *RelationResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *GetFollowRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetFollowRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *GetFollowResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetFollowResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *GetFollowResponse) sizeField2() (n int) {
	if x.UserList == nil {
		return n
	}
	for i := range x.UserList {
		n += fastpb.SizeMessage(2, x.UserList[i])
	}
	return n
}

func (x *GetFollowerRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetFollowerRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *GetFollowerResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetFollowerResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *GetFollowerResponse) sizeField2() (n int) {
	if x.UserList == nil {
		return n
	}
	for i := range x.UserList {
		n += fastpb.SizeMessage(2, x.UserList[i])
	}
	return n
}

func (x *FriendUser) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *FriendUser) sizeField1() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.User)
	return n
}

func (x *FriendUser) sizeField2() (n int) {
	if x.MsgType == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.MsgType)
	return n
}

func (x *FriendUser) sizeField3() (n int) {
	if x.Message == "" {
		return n
	}
	n += fastpb.SizeString(3, x.Message)
	return n
}

func (x *GetFriendRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetFriendRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *GetFriendResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetFriendResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *GetFriendResponse) sizeField2() (n int) {
	if x.UserList == nil {
		return n
	}
	for i := range x.UserList {
		n += fastpb.SizeMessage(2, x.UserList[i])
	}
	return n
}

func (x *CheckFollowRelationRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CheckFollowRelationRequest) sizeField1() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.UserId)
	return n
}

func (x *CheckFollowRelationRequest) sizeField2() (n int) {
	if x.ToUserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.ToUserId)
	return n
}

func (x *CheckFollowRelationRequest) sizeField3() (n int) {
	if x.RelationType == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.RelationType)
	return n
}

func (x *CheckFollowRelationResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CheckFollowRelationResponse) sizeField1() (n int) {
	if !x.Follow {
		return n
	}
	n += fastpb.SizeBool(1, x.Follow)
	return n
}

var fieldIDToName_RelationRequest = map[int32]string{
	1: "UserId",
	2: "ToUserId",
	3: "ActionType",
}

var fieldIDToName_RelationResponse = map[int32]string{
	1: "BaseResp",
}

var fieldIDToName_GetFollowRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetFollowResponse = map[int32]string{
	1: "BaseResp",
	2: "UserList",
}

var fieldIDToName_GetFollowerRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetFollowerResponse = map[int32]string{
	1: "BaseResp",
	2: "UserList",
}

var fieldIDToName_FriendUser = map[int32]string{
	1: "User",
	2: "MsgType",
	3: "Message",
}

var fieldIDToName_GetFriendRequest = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetFriendResponse = map[int32]string{
	1: "BaseResp",
	2: "UserList",
}

var fieldIDToName_CheckFollowRelationRequest = map[int32]string{
	1: "UserId",
	2: "ToUserId",
	3: "RelationType",
}

var fieldIDToName_CheckFollowRelationResponse = map[int32]string{
	1: "Follow",
}

var _ = userdemo.File_userdemo_proto
