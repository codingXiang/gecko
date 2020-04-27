
/*
Auto Create By Moduler
User 模組的 Repository implement
*/
package repository

import (
    "github.com/jinzhu/gorm"
)

//Repository 實例
type UserRepository struct {
	orm *gorm.DB
}

//建立
func NewUserRepository(orm *gorm.DB) user.Repository {
	return &UserRepository{
		orm: orm,
	}
}

/*
    以下實作 Repository 方法
*/
//此為自動產生，建議不要進行更動
func (g *UserRepository) GetUser(data model.UserInterface) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.First(&in).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserRepository) GetUserList(data map[string]interface{}) ([]*model.User, error) {
	var (
		err error
		in  = make([]*model.User, 0)
	)
	err = g.orm.Find(&result, data).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserRepository) CreateUser(data model.UserInterface) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.Create(&in).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserRepository) UpdateUser(data model.UserInterface) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.Update(&in).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserRepository) ModifyUser(data model.UserInterface, column map[string]interface{}) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.Model(&in).Updates(column).Error
	return in, err
}
//此為自動產生，建議不要進行更動
func (g *UserRepository) DeleteUser(data model.UserInterface) (error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = g.orm.Delete(&in).Error
	return err
}


