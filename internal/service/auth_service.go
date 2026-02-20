package service

import (
	"blog_api/internal/model"
	"blog_api/internal/repository"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo      *repository.UserRepository
	JWTSecret string
}

func NewAuthService(repo *repository.UserRepository, secret string) *AuthService {
	return &AuthService{
		Repo:      repo,
		JWTSecret: secret,
	}
}

func (s *AuthService) Register(input model.RegisterRequest) error {
	existingUser, err := s.Repo.CheckExistingUser(input.Email)

	if err == nil && existingUser != nil {
		return errors.New("pendaftaran gagal: username, email, atau nomor telepon sudah terdaftar")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("gagal memproses password")
	}

	user := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := s.Repo.CreateUser(user); err != nil {
		return errors.New("gagal membuat akun baru")
	}

	return nil
}

func (s *AuthService) Login(input model.LoginRequest) (string, error) {
	user, err := s.Repo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("akun tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("kredensial tidak valid")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", errors.New("gagal membuat token autentikasi")
	}

	return signedToken, nil
}

func (s *AuthService) GetProfile(id uint) (*model.User, error) {
	user, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, errors.New("pengguna tidak ditemukan")
	}

	user.Password = ""

	return user, nil
}