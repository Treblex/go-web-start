package model

import (
	"math"
)

type (

	// Result 分页方法返回结果
	Result struct {
		Count     int         `json:"count"`
		PageSize  int         `json:"page_size"`
		PageCount int         `json:"page_count"`
		PageNow   int         `json:"page_now"`
		List      interface{} `json:"list"`
	}
)

// DataBaselimit  mysql数据分页
func DataBaselimit(limit int, page int, where interface{}, _model listModel, key string, orderBy string, userID string) *Result {
	db := DB
	list := _model.PointerList()

	// 初始化数据库对象
	resultModal := db.Table(_model.TableName())
	if where != nil {
		resultModal = resultModal.Where(where)
	}

	if !_model.IsPublic() {
		resultModal = resultModal.Where(map[string]interface{}{
			"user_id": userID,
		})
	}

	resultModal = _model.Search(resultModal, key)
	// 总数
	var count int
	resultModal.Select("count(id)").Where("deleted_at is NULL").Count(&count)
	// 绑定总数
	// 查询绑定用户列表
	resultModal = _model.Joins(resultModal)

	resultModal.Offset(limit * (page - 1)).Limit(limit).Order(orderBy).Find(list)

	list = _model.Result(list)
	var pageCount int = int(math.Ceil(float64(count) / float64(limit)))
	if list == nil {
		list = []interface{}{}
	}
	return &Result{
		Count:     count,
		PageCount: pageCount,
		PageNow:   page,
		PageSize:  limit,
		List:      list,
	}
}
