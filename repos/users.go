package repos

type User struct {
	Id   int
	Name string
	Age  int
}

type Users struct {
	users []User
}

func NewUsers(users []User) *Users {
	return &Users{
		users: users,
	}
}

func (u *Users) GetIds() []int {
	ids := make([]int, 0, len(u.users))
	for _, f := range u.users {
		ids = append(ids, f.Id)
	}
	return ids
}
