package services

import (
	"errors"
	"shop-api/models"
	"shop-api/utils"
	"sync"
	"time"
)

type UserService interface {
	Register(name, email, password string, role models.Role, shopID int) (*models.User, error)
	Login(email, password string) (*models.User, string, error)
	GetByShopID(shopID int) []models.User
	GetByID(id int) (*models.User, error)
}

type UserServiceImpl struct {
	users  []models.User
	nextID int
	mu     sync.RWMutex
}

func NewUserService() UserService {
	// Create some initial users with hashed passwords
	hashedPassword1, _ := utils.HashPassword("admin123")
	hashedPassword2, _ := utils.HashPassword("admin123")

	return &UserServiceImpl{
		users: []models.User{
			{
				ID:        1,
				Name:      "Super Admin 1",
				Email:     "super@shop1.com",
				Password:  hashedPassword1,
				Role:      models.RoleSuperAdmin,
				ShopID:    1,
				CreatedAt: time.Now(),
			},
			{
				ID:        2,
				Name:      "Admin 1",
				Email:     "admin@shop1.com",
				Password:  hashedPassword2,
				Role:      models.RoleAdmin,
				ShopID:    1,
				CreatedAt: time.Now(),
			},
		},
		nextID: 3,
	}
}

func (s *UserServiceImpl) Register(name, email, password string, role models.Role, shopID int) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if email already exists
	for _, user := range s.users {
		if user.Email == email {
			return nil, errors.New("email already exists")
		}
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:        s.nextID,
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		Role:      role,
		ShopID:    shopID,
		CreatedAt: time.Now(),
	}

	s.nextID++
	s.users = append(s.users, user)

	return &user, nil
}

func (s *UserServiceImpl) Login(email, password string) (*models.User, string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Find user by email
	var user *models.User
	for i := range s.users {
		if s.users[i].Email == email {
			user = &s.users[i]
			break
		}
	}

	if user == nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Check password
	if err := utils.CheckPassword(user.Password, password); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Generate token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *UserServiceImpl) GetByShopID(shopID int) []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var users []models.User
	for _, user := range s.users {
		if user.ShopID == shopID {
			users = append(users, user)
		}
	}
	return users
}

func (s *UserServiceImpl) GetByID(id int) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
