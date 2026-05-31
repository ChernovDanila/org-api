package service

import (
	"errors"
	"strings"

	"github.com/ChernovDanila/org-api/internal/model"
	"github.com/ChernovDanila/org-api/internal/repository"
)

type DepartmentService struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (s *DepartmentService) Create(name string, parentID *int) (*model.Department, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	dept := &model.Department{
		Name:     name,
		ParentID: parentID,
	}

	if err := s.repo.Create(dept); err != nil {
		return nil, err
	}

	return dept, nil
}
