package model

import (
	structpb "google.golang.org/protobuf/types/known/structpb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"svc-user/config"
)

var M = New()

func New() *gorm.DB {
	Conf := config.Parse()
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               Conf.Mysql.User + ":" + Conf.Mysql.Pass + "@tcp(" + Conf.Mysql.Host + ":" + Conf.Mysql.Port + ")/" + Conf.Mysql.Name + "?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	s, _ := db.DB()
	//设置空闲连接池中连接的最大数量
	s.SetMaxIdleConns(10)
	//设置打开数据库连接的最大数量
	//s.SetMaxOpenConns(1000)
	//设置了连接可复用的最大时间
	//s.SetConnMaxLifetime(time.Hour)
	return db
}

type Pages struct {
	Page  int         `json:"page"`
	Pages int64       `json:"pages"`
	Total int64       `json:"total"`
	Limit int         `json:"limit"`
	Docs  interface{} `json:"docs"`
}

type SearchKeyWord struct {
	Search string   // 搜索的值
	Keys   []string // 搜索的字段
}

func WithSearch(search []SearchKeyWord) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		for _, item := range search {
			if item.Search != "" {
				for i, key := range item.Keys {
					if i == 0 {
						db.Where(key+" LIKE ?", "%"+item.Search+"%")
					} else {
						db.Or(key+" LIKE ?", "%"+item.Search+"%")
					}
				}
			}
		}

		return db
	}
}

func Paginate(m interface{}, page, limit int, search []SearchKeyWord, association ...string) (Pages, error) {
	if page <= 0 {
		page = 10
	}
	if limit <= 10 {
		limit = 10
	}

	var total int64
	M.Scopes(WithSearch(search)).Model(m).Count(&total)
	pageNum := total / int64(limit)
	if total%int64(limit) != 0 {
		pageNum++
	}

	result := M.Order("created_at desc").Scopes(WithSearch(search))
	if association != nil {
		for _, ass := range association {
			result = result.Preload(ass)
		}
		//result = result.Preload(clause.Associations)
	}

	result = result.Limit(limit).Offset((page - 1) * limit).Find(m)

	return Pages{
		Page:  page,
		Pages: pageNum,
		Total: total,
		Limit: limit,
		Docs:  m,
	}, result.Error
}

// 转换Protobuf的map<string, google.protobuf.ListValue>类型
func ConversionSearchType(in map[string]*structpb.ListValue) []SearchKeyWord {
	search := []SearchKeyWord{}

	for value, keys := range in {
		k := []string{}
		for _, key := range keys.GetValues() {
			k = append(k, key.GetStringValue())
		}
		search = append(search, SearchKeyWord{value, k})
	}

	return search
}

func ConversionSearchType2(in map[string][]string) []SearchKeyWord {
	search := []SearchKeyWord{}

	for value, keys := range in {
		search = append(search, SearchKeyWord{value, keys})
	}

	return search
}
