package service

import (
	"errors"
	"strings"
	"time"

	"github.com/ChernovDanila/org-api/internal/model"
	"github.com/ChernovDanila/org-api/internal/repository"
)

type EmployeeService struct {
	repo     *repository.EmployeeRepository
	deptRepo *repository.DepartmentRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository, deptRepo *repository.DepartmentRepository) *EmployeeService {
	return &EmployeeService{repo: repo, deptRepo: deptRepo}
}

func (s *EmployeeService) Create(departmentID int, fullName, position string, hiredAt *time.Time) (*model.Employee, error) {
	fullName = strings.TrimSpace(fullName)
	position = strings.TrimSpace(position)

	if fullName == "" {
		return nil, errors.New("full_name cannot be empty")
	}
	if position == "" {
		return nil, errors.New("position cannot be empty")
	}

	emp := &model.Employee{
		DepartmentID: departmentID,
		FullName:     fullName,
		Position:     position,
		HiredAt:      hiredAt,
	}

	if err := s.repo.Create(emp); err != nil {
		return nil, err
	}

	return emp, nil
}
