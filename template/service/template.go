package service

var (
	Get = []byte(`Get{{method}}(data model.{{ .module }}Interface) (*model.{{ .module }}, error)`)
	GetList = []byte(`Get{{method}}List(data map[string]interface{}) ([]*model.{{ .module }}, error)`)
	Create = []byte(`Create{{method}}(data model.{{ .module }}Interface) (*model.{{ .module }}, error)`)
	Update = []byte(`Update{{method}}(data model.{{ .module }}Interface) (*model.{{ .module }}, error)`)
	Modify = []byte(`Modify{{method}}(data model.{{ .module }}Interface, data map[string]interface{}) (*model.{{ .module }}, error)`)
	Delete = []byte(`Delete{{method}}(data model.{{ .module }}Interface) error`)
)

var (
	INTERFACE = []byte(`
/*
Auto Create By Moduler
{{ .package }} 模組 module 的 Service Interface
*/
package {{ .package }}

//Service 用於封裝商業邏輯的方法
//go:generate mockgen -destination mock/mock_service.go -package mock -source service.go
type Service interface {
    /*
        寫入封裝方法
    */
{{ .implement }}
}
`)
	IMPLEMENT = []byte(`
/*
Auto Create By Moduler
//{{ .module }} 模組的 Service implement
*/
package service

//Service 實例
type {{ .module }}Service struct {
	{{ .package }}Repo {{ .package }}.Repository
}

//建立
func New{{ .module }}Service({{ .module }}Repo {{ .package }}.Repository) {{ .package }}.Service {
	return &{{ .module }}Service{
		{{ .package }}Repo: {{ .package }}Repo,
	}
}

/*
    以下實作 Service 方法
*/

{{ .implement }}

`)
	TEST = []byte(`
//Package service_test 用於測試 {{ .module }} 模組的 Repository
package service_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Suite 為設定 mock repository
type Suite struct {
	suite.Suite
	service {{ .package }}.Service
}

//測試變數
//var (
//	testData = []*model.Test{
//		&model.Test{
//			ID:                   "test",
//			Name:                 "test",
//		}
//	}
//)

//初始化 Suite
func (s *Suite) SetupSuite() {
	// 建立 mock controller
	ctrl := gomock.NewController(s.T())
	// 透過 mock 建立 repository
	repo := mock.NewMockRepository(ctrl)
	/*
		建立 repository mock data
	*/
	//// Example
	//{
	//	//Example
	//	repo.EXPECT().CreateXXX(testData).Return(testData, nil)
	//}
	// 初始化 demoService
	s.service = service.New{{ .module }}Service(repo)
}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

/*
    以下撰寫測試用例
*/


//TestCase 用於測試 Repository 中的 xxx方法
//func (s *Suite) TestCase() {
    // 情境
	//{
	//	data, err := s.service.CreateXXX(testData)
	//	require.NoError(s.T(), err)
	//	require.NotNil(s.T(), data)
	//	require.IsType(s.T(), testData, data)
	//}
//}
{{ .implement }}
`)
)


var (
	INTERFACE_ABSTRACT_METHOD = []byte(`    {{ .method.verb }}{{ .module }}{{ .method.extension }}({{ .method.param }}) ({{ .method.return.type }})`)
	INTERFACE_METHOD          = []byte(`//此為自動產生，建議不要進行更動
func (g *{{ .module }}Service) {{ .method.verb }}{{ .module }}{{ .method.extension }}({{ .method.param }}) ({{ .method.return.type }}) {
	{{ .method.action }}
	return {{ .method.return }}
}`)
	INTERFACE_METHOD_VARIABLE = []byte(`var (
		err error
        {{ .result }}
	)`)
	INTERFACE_METHOD_ACTION = []byte(`{{ .method.variable }} := g.{{ .package }}Repo.{{ .method.verb }}{{ .module }}{{ .method.extension }}({{ .method.param.value }})`)
	CRUD = []string{"Get", "GetList", "Create", "Update", "Modify", "Delete"}

	INTERFACE_TEST = []byte(`//Test{{ .method.verb }}{{ .module }}{{ .method.extension }} 用於測試 Service 中的 {{ .method.verb }}{{ .module }}{{ .method.extension }}
func (s *Suite) Test{{ .method.verb }}{{ .module }}{{ .method.extension }}() {
	panic("implement me")
}`)
)
