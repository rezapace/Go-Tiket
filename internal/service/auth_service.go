package service

import (
	"Ticketing/entity"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// login
type LoginUseCase interface {
	Login(ctx context.Context, email string, password string) (*entity.User, error)
}

type LoginRepository interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type loginService struct {
	repository LoginRepository
}

func NewLoginService(repository LoginRepository) *loginService {
	return &loginService{
		repository: repository,
	}
}

// func untuk melakikan pengecekan apakah semua data nya sama mulai dari email, password
func (s *loginService) Login(ctx context.Context, email string, password string) (*entity.User, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	//untuk pengecakan apakah email  ada di database?
	if user == nil {
		return nil, errors.New("user with that email not found")
	}

	//untuk pengecekan apakah password nya ada atau gaa di databse?
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect login credentials")
	}
	//ketika email dan passwerd sama maka akan mengembalikan nil
	return user, nil

}

// register
type RegistrationUseCase interface {
	Registration(ctx context.Context, user *entity.User) error
}

type RegistrationRepository interface {
	Registration(ctx context.Context, user *entity.User) error
	// GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type registrationService struct {
	repository RegistrationRepository
}

func NewRegistrationService(repository RegistrationRepository) *registrationService {
	return &registrationService{
		repository: repository,
	}
}

func (s *registrationService) Registration(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.repository.Registration(ctx, user)
}

// BuyerCreateAccount
type BuyerCreateAccountUseCase interface {
	BuyerCreateAccount(ctx context.Context, user *entity.User) error
}

type BuyerCreateAccountRepository interface {
	BuyerCreateAccount(ctx context.Context, user *entity.User) error
}

type buyercreateaccountService struct {
	repository BuyerCreateAccountRepository
}

func NewBuyerCreateAccountService(repository BuyerCreateAccountRepository) *buyercreateaccountService {
	return &buyercreateaccountService{
		repository: repository,
	}
}

// func (s *buyercreateaccountService) BuyerCreateAccount(ctx context.Context, user *entity.User) error {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	user.Password = string(hashedPassword)
// 	return s.repository.BuyerCreateAccount(ctx, user)
// }