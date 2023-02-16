package cache



func Init(){
	MapLoginUser = make(map[string]User, 0)
	MapUser = make(map[int64]User, 0)
}