package repository

import "gorm.io/gorm"

type Repositories struct {
	Users  UserRepository
	Menus  MenuRepository
	Orders OrderRepository
}

func NewRepositories(db *gorm.DB) Repositories {
	return Repositories{
		Users:  NewUserRepository(db),
		Menus:  NewMenuRepository(db),
		Orders: NewOrderRepository(db),
	}
}
