package repository

import (
	"context"

	"github.com/Ngab-Rio/NOCs-API/internal/models"
	"gorm.io/gorm"
)

type PackageRepository interface {
	Create(ctx context.Context, pkg *models.Package) error
	Update(ctx context.Context, pkg *models.Package) error
	Delete(ctx context.Context, pkg *models.Package) error
	FindByID(ctx context.Context, packageID int) (*models.Package, error)
	FindBySpeed(ctx context.Context, speed int) (*models.Package, error)
	FindAll(ctx context.Context) ([]models.Package, error)
}

type packageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(db *gorm.DB) PackageRepository {
	return &packageRepository{db: db}
}

func (r *packageRepository) Create(ctx context.Context, pkg *models.Package) error {
	return r.db.WithContext(ctx).Create(pkg).Error
}

func (r *packageRepository) Update(ctx context.Context, pkg *models.Package) error {
	return r.db.WithContext(ctx).Save(pkg).Error
}

func (r *packageRepository) Delete(ctx context.Context, pkg *models.Package) error {
	return r.db.WithContext(ctx).Delete(pkg).Error
}

func (r *packageRepository) FindByID(ctx context.Context, packageID int) (*models.Package, error) {
	var pkg models.Package
	if err := r.db.WithContext(ctx).Where("id = ?", packageID).First(&pkg).Error; err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *packageRepository) FindBySpeed(ctx context.Context, speed int) (*models.Package, error) {
	var pkg models.Package
	if err := r.db.WithContext(ctx).Where("speed = ?", speed).First(&pkg).Error; err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *packageRepository) FindAll(ctx context.Context) ([]models.Package, error) {
	packages := []models.Package{}
	if err := r.db.WithContext(ctx).Find(&packages).Error; err != nil {
		return nil, err
	}
	return packages, nil
}
