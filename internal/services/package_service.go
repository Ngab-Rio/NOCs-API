package services

import (
	"context"
	"errors"

	"github.com/Ngab-Rio/NOCs-API/internal/dto"
	"github.com/Ngab-Rio/NOCs-API/internal/models"
	"github.com/Ngab-Rio/NOCs-API/internal/repository"
	"gorm.io/gorm"
)

type PackageService interface {
	CreatePackage(ctx context.Context, req dto.CreatePackageRequest) (*dto.PackageResponse, error)
	UpdatePackage(ctx context.Context, packageID int, req *dto.UpdatePackageRequest) (*dto.PackageResponse, error)
	DeletePackage(ctx context.Context, packageID int) error
	FindPackageByID(ctx context.Context, packageID int) (*models.Package, error)
	FindAllPackages(ctx context.Context) ([]dto.PackageResponse, error)
}

type packageService struct {
	packageRepo repository.PackageRepository
}

func NewPackageService(packageRepo repository.PackageRepository) PackageService {
	return &packageService{packageRepo: packageRepo}
}

func (s *packageService) CreatePackage(ctx context.Context, req dto.CreatePackageRequest) (*dto.PackageResponse, error) {
	existingPkg, err := s.packageRepo.FindBySpeed(ctx, req.Speed)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		existingPkg = nil
	}

	if existingPkg != nil {
		return nil, errors.New("package with this speed already exists")
	}

	pkg := &models.Package{
		Name:   req.Name,
		Speed:  req.Speed,
		Price:  req.Price,
		Status: "active",
	}

	if err := s.packageRepo.Create(ctx, pkg); err != nil {
		return nil, err
	}

	return &dto.PackageResponse{
		ID:     pkg.ID,
		Name:   pkg.Name,
		Speed:  pkg.Speed,
		Price:  pkg.Price,
		Status: pkg.Status,
	}, nil
}
func (s *packageService) UpdatePackage(ctx context.Context, packageID int, req *dto.UpdatePackageRequest) (*dto.PackageResponse, error) {
	pkg, err := s.packageRepo.FindByID(ctx, packageID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("package not found")
		}
		return nil, err
	}
	if pkg == nil {
		return nil, errors.New("package not found")
	}

	if req.Name != nil {
		pkg.Name = *req.Name
	}

	if req.Speed != nil {
		pkg.Speed = *req.Speed
	}

	if req.Price != nil {
		pkg.Price = *req.Price
	}

	if req.Status != nil {
		pkg.Status = *req.Status
	}

	if err := s.packageRepo.Update(ctx, pkg); err != nil {
		return nil, err
	}

	return &dto.PackageResponse{
		ID:     pkg.ID,
		Name:   pkg.Name,
		Speed:  pkg.Speed,
		Price:  pkg.Price,
		Status: pkg.Status,
	}, nil
}

func (s *packageService) DeletePackage(ctx context.Context, packageID int) error {
	pkg, err := s.packageRepo.FindByID(ctx, packageID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("package not found")
		}
		return err
	}

	if pkg == nil {
		return errors.New("package not found")
	}

	if err := s.packageRepo.Delete(ctx, pkg); err != nil {
		return err
	}

	return nil
}

func (s *packageService) FindPackageByID(ctx context.Context, packageID int) (*models.Package, error) {
	pkg, err := s.packageRepo.FindByID(ctx, packageID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("package not found")
		}
		return nil, err
	}

	if pkg == nil {
		return nil, errors.New("package not found")
	}

	return pkg, nil
}

func (s *packageService) FindAllPackages(ctx context.Context) ([]dto.PackageResponse, error) {
	packages, err := s.packageRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if packages == nil {
		packages = []models.Package{}
	}

	responses := []dto.PackageResponse{}

	for _, pkg := range packages {
		responses = append(responses, dto.PackageResponse{
			ID:     pkg.ID,
			Name:   pkg.Name,
			Speed:  pkg.Speed,
			Price:  pkg.Price,
			Status: pkg.Status,
		})
	}

	return responses, nil
}
