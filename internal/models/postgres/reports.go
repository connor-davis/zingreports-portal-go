package postgres

import (
	"github.com/connor-davis/zingreports-portal-go/internal/helpers"
	"gorm.io/gorm"
)

type Report struct {
	Base
	Name    string         `json:"name" gorm:"type:text;not null;" validate:"required"`
	Table   ReportTable    `json:"table" gorm:"foreignKey:ReportId;references:Id;constraint:onDelete:CASCADE;"`
	Columns []ReportColumn `json:"columns" gorm:"foreignKey:ReportId;references:Id;constraint:onDelete:CASCADE;"`
	Filters []ReportFilter `json:"filters" gorm:"foreignKey:ReportId;references:Id;constraint:onDelete:CASCADE;"`
}

func (r *Report) BeforeCreate(tx *gorm.DB) error {
	err := helpers.Validate(r)

	if err != nil {
		return err
	}

	return nil
}

type ReportTable struct {
	Base
	ReportId   string                 `json:"reportId" gorm:"type:text;not null;" validate:"required"`
	Name       string                 `json:"name" gorm:"type:text;not null;" validate:"required"`
	Columns    []ReportTableColumn    `json:"columns" gorm:"foreignKey:TableId;references:Id;constraint:onDelete:CASCADE;"`
	References []ReportTableReference `json:"references" gorm:"foreignKey:TableId;references:Id;constraint:onDelete:CASCADE;"`
}

func (rt *ReportTable) BeforeCreate(tx *gorm.DB) error {
	err := helpers.Validate(rt)

	if err != nil {
		return err
	}

	return nil
}

type ReportTableColumn struct {
	Base
	TableId string `json:"tableId" gorm:"type:text;not null;" validate:"required"`
	Name    string `json:"name" gorm:"type:text;not null;" validate:"required"`
}

func (rtc *ReportTableColumn) BeforeCreate(tx *gorm.DB) error {
	err := helpers.Validate(rtc)

	if err != nil {
		return err
	}

	return nil
}

type ReportTableReference struct {
	Base
	TableId               string `json:"tableId" gorm:"type:text;not null;" validate:"required"`
	PrimaryTable          string `json:"primaryTable" gorm:"type:text;not null;" validate:"required"`
	PrimaryTableColumn    string `json:"primaryTableColumn" gorm:"type:text;not null;" validate:"required"`
	ReferencedTable       string `json:"referencedTable" gorm:"type:text;not null;" validate:"required"`
	ReferencedTableColumn string `json:"referencedTableColumn" gorm:"type:text;not null;" validate:"required"`
}

func (rtr *ReportTableReference) BeforeCreate(tx *gorm.DB) error {
	err := helpers.Validate(rtr)

	if err != nil {
		return err
	}

	return nil
}

type ReportColumn struct {
	Base
	ReportId   string `json:"reportId" gorm:"type:text;not null;" validate:"required"`
	Name       string `json:"name" gorm:"type:text;not null;" validate:"required"`
	CustomName string `json:"customName" gorm:"type:text;not null;" validate:"required"`
}

func (rc *ReportColumn) BeforeCreate(tx *gorm.DB) error {
	err := helpers.Validate(rc)

	if err != nil {
		return err
	}

	return nil
}

type ReportFilterType string

const (
	Equal              ReportFilterType = "eq"
	NotEqual           ReportFilterType = "noteq"
	GreaterThan        ReportFilterType = "gt"
	GreaterThanOrEqual ReportFilterType = "gte"
	LessThan           ReportFilterType = "lt"
	LessThanOrEqual    ReportFilterType = "lte"
	StartsWith         ReportFilterType = "startsWith"
	EndsWith           ReportFilterType = "endsWith"
	Contains           ReportFilterType = "contains"
	In                 ReportFilterType = "in"
	NotIn              ReportFilterType = "notin"
)

type ReportFilter struct {
	Base
	ReportId   string           `json:"reportId" gorm:"type:text;not null;" validate:"required"`
	Type       ReportFilterType `json:"type" gorm:"type:text;not null;" validate:"required"`
	Value      string           `json:"value" gorm:"type:text;not null;" validate:"required"`
	ColumnName string           `json:"columnName" gorm:"type:text;not null;" validate:"required"`
	ColumnType string           `json:"columnType" gorm:"type:text;not null;" validate:"required"`
}

func (rf *ReportFilter) BeforeCreate(tx *gorm.DB) error {
	err := helpers.Validate(rf)

	if err != nil {
		return err
	}

	return nil
}
