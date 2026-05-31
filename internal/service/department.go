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

func (s *DepartmentService) GetByID(id, depth int, includeEmployees bool) (*model.Department, error) {
	dept, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if includeEmployees {
		employees, err := s.repo.GetEmployees(id)
		if err != nil {
			return nil, err
		}
		dept.Employees = employees
	}

	if depth > 0 {
		children, err := s.repo.GetChildren(id)
		if err != nil {
			return nil, err
		}
		for i := range children {
			child, err := s.GetByID(children[i].ID, depth-1, includeEmployees)
			if err != nil {
				return nil, err
			}
			children[i] = *child
		}
		dept.Children = children
	}

	return dept, nil
}
