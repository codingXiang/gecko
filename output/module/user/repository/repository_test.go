
//Package repository_test 用於測試 User 模組的 Repository
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
	repository user.Repository
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
	s.repository = repository.NewUserRepository(s.DB)
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
//TestGetUser 用於測試 Repository 中的 GetUser
func (s *Suite) TestGetUser() {
	panic("implement me")
}
//TestGetUserList 用於測試 Repository 中的 GetUserList
func (s *Suite) TestGetUserList() {
	panic("implement me")
}
//TestCreateUser 用於測試 Repository 中的 CreateUser
func (s *Suite) TestCreateUser() {
	panic("implement me")
}
//TestUpdateUser 用於測試 Repository 中的 UpdateUser
func (s *Suite) TestUpdateUser() {
	panic("implement me")
}
//TestModifyUser 用於測試 Repository 中的 ModifyUser
func (s *Suite) TestModifyUser() {
	panic("implement me")
}
//TestDeleteUser 用於測試 Repository 中的 DeleteUser
func (s *Suite) TestDeleteUser() {
	panic("implement me")
}

