package repository

import (
	"context"

	"github.com/Ngab-Rio/NOCs-API/internal/dto"
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
	FindAll(ctx context.Context, req dto.GetClientsRequest) ([]models.Client, int64, error)
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

func (r *clientRepository) FindAll(ctx context.Context, req dto.GetClientsRequest) ([]models.Client, int64, error) {
	var clients []models.Client
	var total int64

	baseQuery := r.db.WithContext(ctx).Model(&models.Client{})

	if req.Search != "" {
		search := "%" + req.Search + "%"
		baseQuery = baseQuery.Where(
			"name ILIKE ? OR email ILIKE ? OR phone ILIKE ? OR address ILIKE ?",
			search, search, search, search,
		)
	}

	if req.Status != "" {
		baseQuery = baseQuery.Where("status = ?", req.Status)
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	allowedSortBy := map[string]bool{
		"name":       true,
		"email":      true,
		"phone":      true,
		"address":    true,
		"created_at": true,
	}

	sortBy := "created_at"
	if allowedSortBy[req.SortBy] {
		sortBy = req.SortBy
	}

	order := "desc"
	if req.Order == "asc" {
		order = "asc"
	}

	query := baseQuery.Order(sortBy + " " + order)

	if req.Limit > 0 {
		query = query.Limit(req.Limit)
	}

	if req.Offset > 0 {
		query = query.Offset(req.Offset)
	}

	if err := query.Find(&clients).Error; err != nil {
		return nil, 0, err
	}

	return clients, total, nil

}
