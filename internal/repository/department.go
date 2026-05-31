package repository

import (
	"github.com/ChernovDanila/org-api/internal/model"
	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) Create(dept *model.Department) error {
	return r.db.Create(dept).Error
}

func (r *DepartmentRepository) GetByID(id int) (*model.Department, error) {
	var dept model.Department
	err := r.db.First(&dept, id).Error
	if err != nil {
		return nil, err
	}
	return &dept, nil
}

func (r *DepartmentRepository) GetChildren(parentID int) ([]model.Department, error) {
	var children []model.Department
	err := r.db.Where("parent_id = ?", parentID).Find(&children).Error
	return children, err
}

func (r *DepartmentRepository) GetEmployees(departmentID int) ([]model.Employee, error) {
	var employees []model.Employee
	err := r.db.Where("department_id = ?", departmentID).Order("created_at").Find(&employees).Error
	return employees, err
}
