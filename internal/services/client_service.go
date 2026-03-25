package services

import (
	"context"
	"errors"

	"github.com/Ngab-Rio/NOCs-API/internal/dto"
	"github.com/Ngab-Rio/NOCs-API/internal/models"
	"github.com/Ngab-Rio/NOCs-API/internal/repository"
	"gorm.io/gorm"
)

type ClientService interface {
	CreateClient(ctx context.Context, req dto.CreateClientRequest) (*dto.ClientResponse, error)
	UpdateClient(ctx context.Context, ClientID int, req dto.UpdateClientRequest) (*dto.ClientResponse, error)
}

type clientService struct {
	clientRepo repository.ClientRepository
}

func NewClientService(clientRepo repository.ClientRepository) ClientService {
	return &clientService{clientRepo: clientRepo}
}

func (s *clientService) CreateClient(ctx context.Context, req dto.CreateClientRequest) (*dto.ClientResponse, error) {
	existingClient, err := s.clientRepo.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingClient != nil {
		return nil, errors.New("email already registered")
	}

	existingPhone, err := s.clientRepo.FindByPhone(ctx, req.Phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingPhone != nil {
		return nil, errors.New("number phone already registered")
	}

	client := &models.Client{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
		Status:    "active",
	}

	if err := s.clientRepo.Create(ctx, client); err != nil {
		return nil, err
	}

	return &dto.ClientResponse{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		Address:   client.Address,
		Longitude: client.Longitude,
		Latitude:  client.Latitude,
		Status:    client.Status,
	}, nil
}

func (s *clientService) UpdateClient(ctx context.Context, clientID int, req dto.UpdateClientRequest) (*dto.ClientResponse, error) {
	client, err := s.clientRepo.FindByID(ctx, clientID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("client not found")
		}
		return nil, err
	}

	if req.Name != nil {
		client.Name = *req.Name
	}
	if req.Email != nil {
		client.Email = *req.Email
	}
	if req.Phone != nil {
		client.Phone = *req.Phone
	}
	if req.Address != nil {
		client.Address = *req.Address
	}
	if req.Longitude != nil {
		client.Longitude = *req.Longitude
	}
	if req.Latitude != nil {
		client.Latitude = *req.Latitude
	}
	if req.Status != nil {
		client.Status = *req.Status
	}

	if err := s.clientRepo.Update(ctx, client); err != nil {
		return nil, err
	}

	return &dto.ClientResponse{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		Address:   client.Address,
		Longitude: client.Longitude,
		Latitude:  client.Latitude,
		Status:    client.Status,
	}, nil
}
