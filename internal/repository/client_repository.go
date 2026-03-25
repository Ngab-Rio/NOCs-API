package repository

import (
	"context"

	"github.com/Ngab-Rio/NOCs-API/internal/models"
	"gorm.io/gorm"
)

type ClientRepository interface {
	Create(ctx context.Context, client *models.Client) error
	Update(ctx context.Context, client *models.Client) error
	Delete(ctx context.Context, client *models.Client) error
	FindByID(ctx context.Context, id int) (*models.Client, error)
	FindByEmail(ctx context.Context, email string) (*models.Client, error)
	FindByPhone(ctx context.Context, phone string) (*models.Client, error)
	FindAll(limit, offset int) ([]models.Client, error)
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db: db}
}

func (r *clientRepository) Create(ctx context.Context, client *models.Client) error {
	return r.db.WithContext(ctx).Create(client).Error
}

func (r *clientRepository) Update(ctx context.Context, client *models.Client) error {
	return r.db.WithContext(ctx).Model(&models.Client{}).
		Where("id = ?", client.ID).
		Updates(client).Error
}

func (r *clientRepository) Delete(ctx context.Context, client *models.Client) error {
	return r.db.WithContext(ctx).Delete(client).Error
}

func (r *clientRepository) FindByID(ctx context.Context, id int) (*models.Client, error) {
	var client models.Client
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) FindByEmail(ctx context.Context, email string) (*models.Client, error) {
	var client models.Client
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) FindByPhone(ctx context.Context, phone string) (*models.Client, error) {
	var client models.Client
	if err := r.db.WithContext(ctx).Where("phone = ?", phone).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) FindAll(limit, offset int) ([]models.Client, error) {
	var clients []models.Client

	err := r.db.Limit(limit).Offset(offset).Find(&clients).Error
	if err != nil {
		return nil, err
	}

	return clients, nil
}
