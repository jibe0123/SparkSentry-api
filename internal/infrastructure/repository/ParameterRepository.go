package repository

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// NewParamRepository creates a new instance of ParameterRepository.
func NewParamRepository(db *gorm.DB) *ParameterRepository {
	return &ParameterRepository{db: db}
}

// ParameterRepository handles data operations for parameters.
type ParameterRepository struct {
	db *gorm.DB
}

// CreateParameter adds a new parameter to the database.
func (repo *ParameterRepository) CreateParameter(equipmentID uint, dto dto.ParameterCreateDTO) error {
	parameter := entities.Parameter{
		Name:        dto.Name,
		HostDevice:  dto.HostDevice,
		Device:      dto.Device,
		Log:         dto.Log,
		Point:       dto.Point,
		EquipmentID: equipmentID,
		Unit:        dto.Unit,
	}
	return repo.db.Create(&parameter).Error
}

// FindByEquipmentID retrieves all parameter associated with a specific equipment.
func (repo *ParameterRepository) FindByEquipmentID(equipmentID uint) ([]entities.Parameter, error) {
	var parameters []entities.Parameter
	err := repo.db.Preload("Equipment").Where("equipment_id = ?", equipmentID).Find(&parameters).Error
	return parameters, err
}
