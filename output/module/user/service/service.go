
/*
Auto Create By Moduler
//User 模組的 Service implement
*/
package service

//Service 實例
type UserService struct {
	userRepo user.Repository
}

//建立
func NewUserService(UserRepo user.Repository) user.Service {
	return &UserService{
		userRepo: userRepo,
	}
}

/*
    以下實作 Service 方法
*/

//此為自動產生，建議不要進行更動
func (g *UserService) GetUser(data model.UserInterface) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.First(&in).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserService) GetUserList(data map[string]interface{}) ([]*model.User, error) {
	var (
		err error
		in  = make([]*model.User, 0)
	)
	err = g.orm.Find(&result, data).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserService) CreateUser(data model.UserInterface) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.Create(&in).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserService) UpdateUser(data model.UserInterface) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.Update(&in).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserService) ModifyUser(data model.UserInterface, column map[string]interface{}) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.Model(&in).Updates(column).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserService) DeleteUser(data model.UserInterface) (error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.Delete(&in).Error
	return err
}


