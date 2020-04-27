package model

//此為自動產生的 Interface，建議不要進行更動
type UserInterface interface {
	GetId() int64
	SetId(in int64) *User
	GetName() string
	SetName(in string) *User
}


type User struct {
	id     int64 `json:"id"`
	name   string
}

//此為自動產生的方法，建議不要更動
func NewUser() UserInterface {
	return &User{}
}

//此為自動產生的方法，建議不要更動
func (g *User) GetId() int64 {
    
    return g.id
}

//此為自動產生的方法，建議不要更動
func (g *User) SetId(in int64) *User {
    g.id = in
    return g
}

//此為自動產生的方法，建議不要更動
func (g *User) GetName() string {
    
    return g.name
}

//此為自動產生的方法，建議不要更動
func (g *User) SetName(in string) *User {
    g.name = in
    return g
}


