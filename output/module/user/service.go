
/*
Auto Create By Moduler
user 模組 module 的 Service Interface
*/
package user

//Service 用於封裝商業邏輯的方法
//go:generate mockgen -destination mock/mock_service.go -package mock -source service.go
type Service interface {
    /*
        寫入封裝方法
    */
    GetUser(data model.UserInterface) (*model.User, error)
    GetUserList(data map[string]interface{}) ([]*model.User, error)
    CreateUser(data model.UserInterface) (*model.User, error)
    UpdateUser(data model.UserInterface) (*model.User, error)
    ModifyUser(data model.UserInterface, column map[string]interface{}) (*model.User, error)
    DeleteUser(data model.UserInterface) (error)

}
