package http

import (
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/xian1367/layout-go-zero/orm/model/scope"
	"github.com/xian1367/layout-go-zero/pkg/helper"
	"github.com/xian1367/layout-go-zero/pkg/orm"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"net/http"
	"strings"
)

type Paging struct {
	CurrentPage int `json:"current_page"` // 当前页
	PerPage     int `json:"per_page"`     // 每页条数
	TotalPage   int `json:"total_page"`   // 总页数
	TotalCount  int `json:"total_count"`  // 总条数
	List        any `json:"list"`         //列表
}

// Query 查询条件
type Query struct {
	PerPage   int                 `json:"per_page"`                                                                                         // 每页条数
	Page      int                 `json:"page"`                                                                                             // 当前页
	Search    map[string][]string `json:"search" swaggertype:"string" example:"search[name]=Jim&search[name]=Jack"`                         // 模糊查询
	Sort      map[string][]string `json:"sort" swaggertype:"string" example:"sort[id]=asc&sort[id]=desc&sort[id]=rand"`                     // 排序规则
	Between   map[string][]string `json:"between" swaggertype:"string" example:"between[time]=12-01,12-31"`                                 // Between范围
	GT        map[string][]string `json:"gt" swaggertype:"string" example:"gt[price]=100"`                                                  // 大于
	LT        map[string][]string `json:"lt" swaggertype:"string" example:"lt[price]=100"`                                                  // 小于
	Avg       map[string][]string `json:"avg" swaggertype:"string" example:"avg[count]=price&avg[sum]=price&avg[max]=price&avg[min]=price"` // 聚合
	Equal     map[string]string   `json:"equal" swaggertype:"string" example:"name=Jim&age=1,2"`                                            // 等于
	Whitelist []string            `swaggerignore:"true"`                                                                                    // 字段白名单

	db      *gorm.DB      `swaggerignore:"true"` // db query 句柄
	request *http.Request `swaggerignore:"true"` // gin context，方便调用
}

var keyWords = []string{"page", "per_page", "sort", "between", "search", "gt", "lt", "avg"}

func Paginate(r *http.Request, db *gorm.DB, data interface{}) Paging {
	// 初始化 Query 实例
	query := &Query{
		request: r,
	}
	query.Whitelist = scope.GetFields(db.Statement.Model)
	query.getQuery()
	query.db = db.Scopes(
		scopeEqual(query),
		scopeGT(query),
		scopeLT(query),
		scopeSearch(query),
		scopeBetween(query),
		scopeAvg(query),
	)
	paging := query.getPaging()
	offset := (query.Page - 1) * query.PerPage

	// 查询数据库
	query.db.
		Scopes(scopeSort(query)).
		Limit(query.PerPage).
		Offset(offset).
		Find(data)

	paging.List = data
	return paging
}

func (query *Query) getQuery() {
	query.Sort = helper.QueryMap(query.request.URL.Query(), "sort")
	query.Search = helper.QueryMap(query.request.URL.Query(), "search")
	query.Between = helper.QueryMap(query.request.URL.Query(), "between")
	query.GT = helper.QueryMap(query.request.URL.Query(), "gt")
	query.LT = helper.QueryMap(query.request.URL.Query(), "lt")
	query.Avg = helper.QueryMap(query.request.URL.Query(), "avg")

	filterQuery := lo.PickBy[string, []string](query.request.URL.Query(), func(key string, _ []string) bool {
		return !lo.Contains[string](keyWords, key)
	})
	query.Equal = lo.MapValues[string, []string, string](filterQuery, func(item []string, _ string) string {
		return item[0]
	})
}

func (query *Query) getPaging() Paging {
	paging := Paging{
		TotalCount: query.getTotalCount(),
	}
	query.PerPage = query.getPerPage()
	paging.TotalPage = query.getTotalPage(paging.TotalCount)
	paging.CurrentPage = query.getCurrentPage(paging.TotalPage)
	query.Page = paging.CurrentPage
	paging.PerPage = query.PerPage
	return paging
}

func scopeEqual(query *Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for key := range query.Equal {
			if query.fieldFilter(key) {
				values := strings.Split(query.Equal[key], ",")
				if len(values) > 1 {
					db.Where(key+" IN ?", values)
				} else {
					db.Where(key+" = ?", values[0])
				}
			}
		}
		return db
	}
}

func scopeGT(query *Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for key := range query.GT {
			if query.fieldFilter(key) {
				values := query.GT[key]
				if len(values) > 1 {
					or := orm.DB
					for _, value := range values {
						or.Or(key+" > ?", value)
					}
					db.Where(or)
				} else {
					db.Where(key+" > ?", values[0])
				}
			}
		}
		return db
	}
}

func scopeLT(query *Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for key := range query.LT {
			if query.fieldFilter(key) {
				values := query.LT[key]
				if len(values) > 1 {
					or := orm.DB
					for _, value := range values {
						or.Or(key+" < ?", value)
					}
					db.Where(or)
				} else {
					db.Where(key+" < ?", values[0])
				}
			}
		}
		return db
	}
}

func scopeSearch(query *Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for key := range query.Search {
			if query.fieldFilter(key) {
				values := query.Search[key]
				if len(values) > 1 {
					or := orm.DB
					for _, value := range values {
						or.Or(key+" LIKE ?", "%"+value+"%")
					}
					db.Where(or)
				} else {
					db.Where(key+" LIKE ?", "%"+values[0]+"%")
				}
			}
		}
		return db
	}
}

func scopeBetween(query *Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for key := range query.Between {
			if query.fieldFilter(key) {
				values := query.Between[key]
				or := orm.DB
				for _, between := range values {
					if len(values) == 2 {
						or.Or(key+" BETWEEN ? AND ?", between[0], between[1])
					}
				}
				db.Where(or)
			}
		}
		return db
	}
}

func scopeAvg(query *Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for key := range query.Avg {
			if query.avgFilter(key) {
				value := query.Avg[key][0]
				if query.fieldFilter(value) {
					if key == "count" {
						db.Select("COUNT(" + value + ") AS count")
					} else if key == "sum" {
						db.Select("SUM(" + value + ") AS sum")
					} else if key == "max" {
						db.Select("MAX(" + value + ") AS max")
					} else if key == "min" {
						db.Select("MIN(" + value + ") AS min")
					}
				}
			}
		}
		return db
	}
}

func scopeSort(query *Query) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var sqlString []string
		var SqlVars []interface{}
		for key := range query.Sort {
			if query.fieldFilter(key) {
				values := query.Sort[key]
				if query.orderByFilter(values[0]) {
					if values[0] == "rand" {
						sqlString = append(sqlString, "RAND()")
					} else {
						sqlString = append(sqlString, key+" "+values[0])
					}
				} else {
					sqlString = append(sqlString, "FIELD("+key+", ?) DESC")
					SqlVars = append(SqlVars, values)
				}
			}
		}
		sqlString = append(sqlString, "id DESC")
		db.Clauses(clause.OrderBy{
			Expression: clause.Expr{SQL: strings.Join(sqlString, ", "), Vars: SqlVars, WithoutParentheses: true},
		})
		return db
	}
}

func (query *Query) fieldFilter(key string) bool {
	return lo.Contains[string](query.Whitelist, key[strings.Index(key, ".")+1:])
}

func (query *Query) orderByFilter(value string) bool {
	return lo.Contains[string]([]string{"asc", "desc", "field", "rand"}, value)
}

func (query *Query) avgFilter(value string) bool {
	return lo.Contains[string]([]string{"count", "sum", "max", "min"}, value)
}

func (query *Query) getPerPage() int {
	perPage := 10
	queryPerPage := query.request.URL.Query().Get("per_page")
	if len(queryPerPage) > 0 {
		perPage = cast.ToInt(queryPerPage)
	}

	return perPage
}

// getCurrentPage 返回当前页码
func (query *Query) getCurrentPage(totalPage int) int {
	// 优先取用户请求的 page
	page := cast.ToInt(query.request.URL.Query().Get("page"))
	if page <= 0 {
		// 默认为 1
		page = 1
	}
	// TotalPage 等于 0 ，意味着数据不够分页
	if totalPage == 0 {
		return 0
	}
	// 请求页数大于总页数，返回总页数
	if page > totalPage {
		return totalPage
	}
	return page
}

// getTotalCount 返回的是数据库里的条数
func (query *Query) getTotalCount() int {
	var count int64
	if err := query.db.Count(&count).Error; err != nil {
		return 0
	}
	return cast.ToInt(count)
}

// getTotalPage 计算总页数
func (query *Query) getTotalPage(totalCount int) int {
	if totalCount == 0 {
		return 0
	}
	nums := math.Ceil(float64(totalCount) / float64(query.PerPage))
	if nums == 0 {
		nums = 1
	}
	return cast.ToInt(nums)
}
