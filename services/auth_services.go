package services

import (
	"fmt"
	"jadwalin/dto"
	"jadwalin/model"
	"jadwalin/repository"
	"jadwalin/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CreateStudent(userRole string, input dto.StudentRegisterRequest) error 
	CreateLecturer(userRole string, input dto.LecturerRegisterRequest) error
	CreateAdmin(userRole string, input dto.AdminRegisterRequest) error
	Login(input dto.LoginUserRequest)(dto.LoginResponseDTO, error)		
}

type authService struct {
	authRepo repository.AuthRepository
	userRepo repository.UserRepository
}

func NewAuthService(authRepo repository.AuthRepository, userRepo repository.UserRepository) AuthService {
	return &authService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (s *authService) CreateStudent(userRole string, input dto.StudentRegisterRequest) error {
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}

	existingUserWithEmail, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return err
	}

	if existingUserWithEmail != nil {
		return fmt.Errorf("email already exist")
	}

	existingUserWithNim, err := s.userRepo.FindByNim(input.NIM)
	if err != nil {
		return err
	}
	if existingUserWithNim != nil {
		return fmt.Errorf("nim already exiist")
	}

	if input.Password != input.ConfirmPassword {
		return fmt.Errorf("password dan confirm password tidak cocok")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	User := model.User{
		Name: input.Name,
		Email: input.Email,
		Password: string(hashPassword),
		Role: input.Role,
	}

	err = s.authRepo.CreateUser(&User)
	if err != nil {
		return err
	}

	Student := model.Student{
		UserID: User.Id,
		Nim: input.NIM,
		GradeId: input.GradeId,
		ProdiId: input.ProdiId,
	}

	err = s.authRepo.CreateStudent(&Student)
	if err != nil {
		return err
	}

	return nil 
}

func (s *authService) CreateLecturer(userRole string, input dto.LecturerRegisterRequest) error {
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}
	existingUserWithEmail, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return err
	}

	if existingUserWithEmail != nil {
		return fmt.Errorf("email already exist")
	}

	existingUserWithNip, err := s.userRepo.FindByNip(input.NIP)
	if err != nil {
		return err
	}
	if existingUserWithNip != nil {
		return fmt.Errorf("nip already exiist")
	}

	if input.Password != input.ConfirmPassword {
		return fmt.Errorf("password dan confirm password tidak cocok")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	User := model.User{
		Name: input.Name,
		Email: input.Email,
		Password: string(hashPassword),
		Role: input.Role,
	}

	err = s.authRepo.CreateUser(&User)
	if err != nil {
		return err
	}

	Lecturer := model.Lecturer{
		UserID: User.Id,
		Nip: input.NIP,
		Prodi: input.Prodi,
	}

	err = s.authRepo.CreateLecturer(&Lecturer)
	if err != nil {
		return err
	}

	return nil 
}

func (s *authService) CreateAdmin(userRole string, input dto.AdminRegisterRequest) error {
	err := utils.RoleCheck(userRole, "admin")
	if err != nil {
		return err
	}
	existingUserWithEmail, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return err
	}

	if existingUserWithEmail != nil {
		return fmt.Errorf("email already exist")
	}

	if input.Password != input.ConfirmPassword {
		return fmt.Errorf("password dan confirm password tidak cocok")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Name: input.Name,
		Email: input.Email,
		Password: string(hashPassword),
		Role: input.Role,
	}

	err = s.authRepo.CreateUser(&user)
	if err != nil {
		return err
	}

	return nil 

}

func (s *authService) Login(input dto.LoginUserRequest)(dto.LoginResponseDTO, error) {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return dto.LoginResponseDTO{}, err
	}

	if user == nil {
		return dto.LoginResponseDTO{}, fmt.Errorf("email tidak terdaftar")
	}

	ok := utils.CheckPasswordHash(input.Password, user.Password)
	if !ok {
		return dto.LoginResponseDTO{}, fmt.Errorf("password salah")
	}

	token, err := utils.GenerateJWT(uint(user.Id), user.Email, user.Role)
	if err != nil {
		return dto.LoginResponseDTO{}, fmt.Errorf("gagal mengenerate token")
	}

	return dto.LoginResponseDTO{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		Token: token,
	}, nil 
}