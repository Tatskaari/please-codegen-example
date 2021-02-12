package dao

type dao struct {

}

func New() *dao {
	return new(dao)
}

func (*dao) StoreUser(username, email string) error {
	return nil
}