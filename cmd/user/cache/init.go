package cache



func Init(){
	MapLoginUser = make(map[string]int64, 0)
	MapUser = make(map[int64]User, 0)
}