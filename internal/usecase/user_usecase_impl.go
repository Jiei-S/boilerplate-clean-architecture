package usecase

import (
	"context"

	pkgErr "github.com/Jiei-S/boilerplate-clean-architecture/pkg/error"
)

var _ UserUsecase = (*UserUsecaseImpl)(nil)

type UserUsecaseImpl struct {
	userRepository UserRepository
}

func (u *UserUsecaseImpl) AddUser(
	ctx context.Context,
	dto *User,
) (*User, *pkgErr.ApplicationError) {
	entity, err := u.userRepository.Save(ctx, dto.ToEntity())
	if err != nil {
		return nil, err
	}
	return ToDTO(entity), nil
}

func (u *UserUsecaseImpl) FindUser(
	ctx context.Context,
	id string,
) (*User, *pkgErr.ApplicationError) {
	entity, err := u.userRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToDTO(entity), nil
}

func NewUserUsecase(
	userRepository UserRepository,
) UserUsecase {
	return &UserUsecaseImpl{
		userRepository: userRepository,
	}
}
