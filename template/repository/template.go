package repository

//var (
//	Get     = []byte(`Get{{method}}(data model.{{ .module }}Interface) (*model.{{ .module }}, error)`)
//	GetList = []byte(`Get{{method}}List(data map[string]interface{}) ([]*model.{{ .module }}, error)`)
//	Create  = []byte(`Create{{method}}(data model.{{ .module }}Interface) (*model.{{ .module }}, error)`)
//	Update  = []byte(`Update{{method}}(data model.{{ .module }}Interface) (*model.{{ .module }}, error)`)
//	Modify  = []byte(`Modify{{method}}(data model.{{ .module }}Interface, data map[string]interface{}) (*model.{{ .module }}, error)`)
//	Delete  = []byte(`Delete{{method}}(data model.{{ .module }}Interface) error`)
//)

var (
	INTERFACE = []byte(`
/*
Auto Create By Moduler
{{ .package }} 模組的 Repository Interface
*/
package {{ .package }}

//Repository 用於與資料庫進行存取的封裝方法
//go:generate mockgen -destination mock/mock_repository.go -package mock -source repository.go
type Repository interface {
/*
    以下宣告 Repository 方法
*/
{{ .implement }}
}
`)
	IMPLEMENT = []byte(`
/*
Auto Create By Moduler
{{ .module }} 模組的 Repository implement
*/
package repository

import (
    "github.com/jinzhu/gorm"
)

//Repository 實例
type {{ .module }}Repository struct {
	orm *gorm.DB
}

//建立
func New{{ .module }}Repository(orm *gorm.DB) {{ .package }}.Repository {
	return &{{ .module }}Repository{
		orm: orm,
	}
}

/*
    以下實作 Repository 方法
*/
{{ .implement }}

`)
	TEST = []byte(`
//Package repository_test 用於測試 {{ .module }} 模組的 Repository
package repository_test

import (
    "database/sql"
    "github.com/jinzhu/gorm"
    "github.com/DATA-DOG/go-sqlmock"
   	"github.com/stretchr/testify/require"
   	"github.com/stretchr/testify/suite"
   	"testing"
)

//Suite 集成 sql mock, repository
type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository {{ .package }}.Repository
}

//初始化 Suite
func (s *Suite) SetupSuite() {
	/*
		宣告參數
	 */
	var (
		db  *sql.DB //SQL連線
		err error   //錯誤
	)
	// 初始化 sql mock，建立 db 的 instance
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	// 透過 gorm 建立 mysql 的 instance
	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)
	// 設定 log 模式
	s.DB.LogMode(false)
	// 設定要測試的 repository
	s.repository = repository.New{{ .module }}Repository(s.DB)
}

//AfterTest 用於測試完畢之後的檢查
func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

/*
    以下撰寫測試用例
*/
{{ .implement }}
`)
	SUBSTITUTION = []byte(`
`)
)

var (
	INTERFACE_ABSTRACT_METHOD = []byte(`    {{ .method.verb }}{{ .module }}{{ .module.extension }}({{ .method.param }}) ({{ .method.return.type }})`)
	INTERFACE_METHOD          = []byte(`//此為自動產生，建議不要進行更動
func (g *{{ .module }}Repository) {{ .method.verb }}{{ .module }}{{ .module.extension }}({{ .method.param }}) ({{ .method.return.type }}) {
	{{ .module.variable }}
	{{ .method.action }}
	{{ .method.return }}
}`)
	INTERFACE_METHOD_VARIABLE = []byte(`var (
		err error
		in  = {{ .module.variable }}
	)`)
	INTERFACE_METHOD_ACTION = []byte(`err = g.orm.{{ .method.action }}.Error`)
	CRUD = []string{"GetList", "Get", "Create", "Update", "Modify", "Delete"}

	INTERFACE_TEST = []byte(`//Test{{ .method.verb }}{{ .module }}{{ .module.extension }} 用於測試 Repository 中的 {{ .method.verb }}{{ .module }}{{ .module.extension }}
func (s *Suite) Test{{ .method.verb }}{{ .module }}{{ .module.extension }}() {
	panic("implement me")
}`)
)
