package lib

import (
	"K-BANK/model"
)

// DuplicateCheck : 주어진 칼럼, 값을 가지고 그 값이 존재하는지 확인하는 함수
func DuplicateCheck(column string, value string) bool {
	var count int64
	model.DB.Limit(1).Where(column+"= ?", value).Find(&model.User{}).Count(&count)

	// 0: 중복 없음, 1: 중복 있음
	if count == 0 {
		return true
	}
	return false
}
