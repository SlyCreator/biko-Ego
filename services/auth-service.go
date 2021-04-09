package services

import (
	"github.com/SlyCreator/biko-Ego/dto"
	"github.com/SlyCreator/biko-Ego/entity"
	"github.com/SlyCreator/biko-Ego/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"

	"log"
)

type AuthService interface {
	CreateUser(user dto.RegisterDTO) entity.User
	IsDuplicateEmail(email string) bool
	VerifyCredential(email string,password string) interface{}
	//FindByEmail(email string) entity.User

}

type authService struct {
	userRepository repository.UserRepository
}



func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (a *authService) CreateUser(user dto.RegisterDTO)entity.User  {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate,smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map #{err}")
	}
	res := a.userRepository.CreateUser(userToCreate)
	return res
}

func (s *authService) IsDuplicateEmail(email string) bool {
	res := s.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}
func (a *authService) VerifyCredential(email string,password string) interface{}  {
	res := a.userRepository.VerifyCredential(email)
	if v ,ok := res.(entity.User); ok{
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return  res
		}
		return false
	}
	return false
}



func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
