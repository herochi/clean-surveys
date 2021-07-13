package interfaces

import (
	"clean_survey/domain"

)

type UserPresenter interface {
	UserResponse(d *viewmodel.UserVM) *domain.User
	UsersResponse(d []*viewmodel.UserVM) []*domain.User
	UserVMResponse(d *domain.User) *viewmodel.UserVM
	UsersVMResponse(ds []*domain.User) []*viewmodel.UserVM
}
