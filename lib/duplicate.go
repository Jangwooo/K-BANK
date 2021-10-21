package lib

import (
	"K-BANK/model"
)

// DuplicateCheck : 주어진 칼럼, 값을 가지고 그 값이 존재하는지 확인하는 함수
func DuplicateCheck(column string, value string) bool {
	result := model.DB.Limit(1).Where(column+"= ?", value).Find(&model.User{}).RowsAffected

	// 0: 중복 없음, 1: 중복 있음
	if result == 0 {
		return true
	}
	return false
}
