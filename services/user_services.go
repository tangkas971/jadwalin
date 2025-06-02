package services

import (
	"jadwalin/dto"
	"jadwalin/repository"
)

type UserService interface {
	FindAll()([]dto.UserResponseDTO, error)
	// FindById(id int)(dto.UserResponseDTO, error)
	// FindByRole(roleUser string)([]dto.UserResponseDTO, error)
	// Delete(roleUser string, id int) error
	// CreateUser(input dto.RegisterUserRequest)(dto.UserResponseDTO, error)
	// LoginUser(input dto.LoginUserRequest)(dto.LoginResponseDTO, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) FindAll()([]dto.UserResponseDTO, error){
	users, err := s.repo.FindAll()
	if err != nil {
		return []dto.UserResponseDTO{}, err
	}
	var userDTOs []dto.UserResponseDTO
	for _, user := range users{
		userDTO := dto.UserResponseDTO{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			Role: user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		userDTOs = append(userDTOs, userDTO)
	}

	return userDTOs, nil 
}

// func (s *userService) FindByRole(roleUser string)([]dto.UserResponseDTO, error){
// 	users, err := s.repo.FindByRole(roleUser)
// 	if err != nil {
// 		return []dto.UserResponseDTO{}, err
// 	}

// 	var userDTOs []dto.UserResponseDTO
// 	for _, user := range users{
// 		userDTO := dto.UserResponseDTO{
// 			Id: user.Id,
// 			Name: user.Name,
// 			NIM: user.NIM,
// 			Email: user.Email,
// 		}
// 		userDTOs = append(userDTOs, userDTO)
// 	}
// 	return userDTOs, nil 
// }

// func (s *userService) FindById(id int)(dto.UserResponseDTO, error){
// 	user, err := s.repo.FindById(id)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound{
// 			return dto.UserResponseDTO{}, fmt.Errorf("user dengan id %d tidak ditemukan", id)
// 		}
// 		return dto.UserResponseDTO{}, err
// 	}

// 	return dto.UserResponseDTO{
// 		Id: user.Id,
// 		Name: user.Name,
// 		NIM: user.NIM,
// 		Email: user.Email,
// 		Role: user.Role,
// 		CreatedAt: user.CreatedAt,
// 		UpdatedAt: user.UpdatedAt,
// 	}, nil 
// }

// func (s *userService) CreateUser(input dto.RegisterUserRequest)(dto.UserResponseDTO, error){
// 	// CEK NIM UNTUK MAHASISWA
// 	if input.Role == "mahasiswa" && input.NIM == nil {
// 		return dto.UserResponseDTO{}, fmt.Errorf("NIM wajib diisi oleh Mahasiswa")
// 	}

// 	// CEK PASSWORD
// 	if input.Password != input.ConfirmPassword {
// 		return dto.UserResponseDTO{}, fmt.Errorf("password dan confirm password tidak cocok")
// 	}
 	
// 	// CEK DUPLIKASI EMAIL
// 	existingUser, err := s.repo.FindByEmail(input.Email)
// 	if err != nil {
// 		return dto.UserResponseDTO{}, err
// 	}

// 	if existingUser != nil {
// 		return dto.UserResponseDTO{}, fmt.Errorf("email already exist")
// 	}

// 	// HASH PASSWORD
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
//     if err != nil {
//         return dto.UserResponseDTO{}, err
//     }

// 	// MENGUBAH DARI DTO => MODEL
// 	user := model.User{
// 		Name: input.Name,
// 		NIM: input.NIM,
// 		Email: input.Email,
// 		Password: string(hashedPassword),
// 		Role: input.Role,
// 	}

// 	// MENGIRIM KE REPOSITORY
// 	err = s.repo.Create(&user)
// 	if err != nil {
// 		return dto.UserResponseDTO{}, err
// 	}

// 	// MENGEMBALIKAN RESPONSE
// 	return dto.UserResponseDTO{
// 		Id: user.Id,
// 		Name: user.Name,
// 		NIM: user.NIM,
// 		Email: user.Email,
// 		Role: user.Role,
// 		CreatedAt: user.CreatedAt,
// 		UpdatedAt: user.UpdatedAt,
// 	}, nil 
// }

// func (s *userService) LoginUser(input dto.LoginUserRequest)(dto.LoginResponseDTO, error){
// 	user, err := s.repo.FindByEmail(input.Email)
// 	if err != nil {
// 		return dto.LoginResponseDTO{}, fmt.Errorf("email tidak terdaftar")
// 	}

// 	ok := utils.CheckPasswordHash(input.Password, user.Password)
// 	if !ok {
// 		return dto.LoginResponseDTO{}, fmt.Errorf("email atau password salah")
// 	}

// 	// membuat jwt token
// 	token, err := utils.GenerateJWT(uint(user.Id), user.Email, user.Role)
// 	if err != nil {
// 		return dto.LoginResponseDTO{}, fmt.Errorf("gagal generate JWT")
// 	}

// 	return dto.LoginResponseDTO{
// 		Id: user.Id,
// 		Name: user.Name,
// 		Email: user.Email,
// 		Role: user.Role,
// 		Token: token,
// 	}, nil 
// } 

// func (s *userService) Delete(roleUser string, id int) error {
// 	if roleUser != "admin" {
// 		return fmt.Errorf("hanya admin yang dapat menghapus user")
// 	}
// 	err := s.repo.Delete(id)
// 	return err
// }