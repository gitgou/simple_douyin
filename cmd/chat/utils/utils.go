package utils

import (
	"fmt"
	"strconv"
	"strings"
)
func GenChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}

func SpliceChatKey(key string)([]int64){
	strUserId := strings.Split(key, "_")
	userId1, _ := strconv.ParseInt(strUserId[0], 10, 64)
	userId2, _ := strconv.ParseInt(strUserId[0], 10, 64)
	return []int64{ userId1, userId2}
}