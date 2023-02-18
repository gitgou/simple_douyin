package utils

import (
	"fmt"
	"strconv"
	"strings"
)
func GenChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%x_%x", userIdB, userIdA)
	}
	return fmt.Sprintf("%x_%x", userIdA, userIdB)
}

func SpliceChatKey(key string)([]int64){
	strUserId := strings.Split(key, "_")
	userId1, _ := strconv.ParseInt(strUserId[0], 10, 64)
	userId2, _ := strconv.ParseInt(strUserId[0], 10, 64)
	return []int64{ userId1, userId2}
}