package helper

var tokenList = make(map[string]uint)

/**
生成新token
@param uint id admin的id
@return string
*/
func SaveToken(id uint) string {
	//清除id对应token,这样同一账号只能一个人登录
	DeleteTokenById(id)
	token := UUidV4()
	tokenList[token] = id
	return token
}

/**
删除token
@param string token
*/
func DeleteToken(token string) {
	delete(tokenList, token)
}

/**
使用id删除token
@param uint id
*/
func DeleteTokenById(id uint) {
	for token, i := range tokenList {
		if i == id {
			delete(tokenList, token)
		}
	}
}

/**
通过token获取id
@param string token
@return uint id
*/
func GetIdByToken(token string) uint {
	if tokenList[token] > 0 {
		return tokenList[token]
	}
	return 0
}
