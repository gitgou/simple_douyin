// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package constants

const (
	VideoTableName           = "video"
	UserTableName            = "user"
	FollowTableName          = "follow"
	CommentTableName         = "comments"
	FriendTableName          = "friend"
	ChatMessageTableName 	 = "chat_message"
	FavoriteTableName		 = "favorite"
	SecretKey                = "secret key"
	IdentityKey              = "id"
	StatusCode               = "status_code"
	StatusMsg                = "status_msg"
	User                     = "user"
	UserID                   = "user_id"
	UserList 				 = "user_list"
	Token                    = "token"
	VideoList                = "video_list"
	NextTime                 = "next_time"
	MsgList					 = "message_list"
	Comment					 = "comment"
	CommentList				 = "comment_list"
	ChatMsgIdKey			 = "ChatMsgId"
	MinioBucketName 		 = "dousheng"
	MinioEndPoint 			 = "127.0.0.1:9000"
	MinioAccessID 			 = "admin"
	MinioAccessKey 			 = "12345678"
	Location 				 = "GuangZhou"
	ApiServiceName           = "demoapi"
	VideoServiceName         = "videodemo"
	UserServiceName          = "userdemo"
	ChatServiceName 		 = "chatdemo"
	RedisServiceName 		 = "redisdemo"
	RelationServiceName		 = "relationdemo"
	InteractionServiceName	 = "interationdemo"
	VideoServiceAddress  	 = "127.0.0.1:8888"
	UserServiceAddress  	 = "127.0.0.1:8889"
	ChatServiceAddress  	 = "127.0.0.1:8890"
	RedisServiceAddress  	 = "127.0.0.1:8891"
	RelationServiceAddress   = "127.0.0.1:8892"
	InteractionServiceAddress= "127.0.0.1:8893"
	RedisZSetKeyFollow		 = "follow_count" //关注数
	RedisZSetKeyFollower	 = "follower_count" //粉丝数
	RedisZSetKeyVideo 		 = "video_count"
	RedisZSetKeyFavorite 	 = "favorite_count" //点赞
	RedisZSetKeyFavorited 	 = "favorited_count" //获赞
	MySQLDefaultDSN          = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress              = "127.0.0.1:2379"
	CPURateLimit     float64 = 80.0
	DefaultLimit             = 10
)
