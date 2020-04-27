
/*
Auto Create By Moduler
user 模組的 Repository Interface
*/
package user

//Repository 用於與資料庫進行存取的封裝方法
//go:generate mockgen -destination mock/mock_repository.go -package mock -source repository.go
type Repository interface {
/*
    以下宣告 Repository 方法
*/
    GetUser(data model.UserInterface) (model.User ,error)
    GetUserList(data model.UserInterface) (model.User ,error)
    CreateUser(data model.UserInterface) (model.User ,error)
    UpdateUser(data model.UserInterface) (model.User ,error)
    ModifyUser(model model.UserInterface, data map[string]interface{}) (model.User ,error)
    DeleteUser(data model.UserInterface) (error)

}
