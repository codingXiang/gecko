
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
	s.service = service.New{{Module}}Service(repo)
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
//TestGetUser 用於測試 Service 中的 GetUser
func (s *Suite) TestGetUser() {
	panic("implement me")
}
//TestGetUserList 用於測試 Service 中的 GetUserList
func (s *Suite) TestGetUserList() {
	panic("implement me")
}
//TestCreateUser 用於測試 Service 中的 CreateUser
func (s *Suite) TestCreateUser() {
	panic("implement me")
}
//TestUpdateUser 用於測試 Service 中的 UpdateUser
func (s *Suite) TestUpdateUser() {
	panic("implement me")
}
//TestModifyUser 用於測試 Service 中的 ModifyUser
func (s *Suite) TestModifyUser() {
	panic("implement me")
}
//TestDeleteUser 用於測試 Service 中的 DeleteUser
func (s *Suite) TestDeleteUser() {
	panic("implement me")
}

