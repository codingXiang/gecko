
//Package service_test 用於測試 User 模組的 Repository
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
	service user.Service
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
	s.service = service.NewUserService(repo)
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
//TestGetUser{{ .method.extension }} 用於測試 Service 中的 GetUser{{ .method.extension }}
func (s *Suite) TestGetUser{{ .method.extension }}() {
	panic("implement me")
}
//TestGetUser{{ .method.extension }} 用於測試 Service 中的 GetUser{{ .method.extension }}
func (s *Suite) TestGetUser{{ .method.extension }}() {
	panic("implement me")
}
//TestCreateUser{{ .method.extension }} 用於測試 Service 中的 CreateUser{{ .method.extension }}
func (s *Suite) TestCreateUser{{ .method.extension }}() {
	panic("implement me")
}
//TestUpdateUser{{ .method.extension }} 用於測試 Service 中的 UpdateUser{{ .method.extension }}
func (s *Suite) TestUpdateUser{{ .method.extension }}() {
	panic("implement me")
}
//TestModifyUser{{ .method.extension }} 用於測試 Service 中的 ModifyUser{{ .method.extension }}
func (s *Suite) TestModifyUser{{ .method.extension }}() {
	panic("implement me")
}
//TestDeleteUser{{ .method.extension }} 用於測試 Service 中的 DeleteUser{{ .method.extension }}
func (s *Suite) TestDeleteUser{{ .method.extension }}() {
	panic("implement me")
}

