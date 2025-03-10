package postgres

type Report struct {
	Base
	Name    string         `json:"name" gorm:"type:text;not null;"`
	Table   ReportTable    `json:"table" gorm:"foreignKey:ReportId;references:Id;constraint:onDelete:CASCADE;"`
	Columns []ReportColumn `json:"columns" gorm:"foreignKey:ReportId;references:Id;constraint:onDelete:CASCADE;"`
	Filters []ReportFilter `json:"filters" gorm:"foreignKey:ReportId;references:Id;constraint:onDelete:CASCADE;"`
}

type ReportTable struct {
	Base
	ReportId   string                 `json:"reportId" gorm:"type:text;not null;"`
	Name       string                 `json:"name" gorm:"type:text;not null;"`
	Columns    []ReportTableColumn    `json:"columns" gorm:"foreignKey:TableId;references:Id;constraint:onDelete:CASCADE;"`
	References []ReportTableReference `json:"references" gorm:"foreignKey:TableId;references:Id;constraint:onDelete:CASCADE;"`
}

type ReportTableColumn struct {
	Base
	TableId string `json:"tableId" gorm:"type:text;not null;"`
	Name    string `json:"name" gorm:"type:text;not null;"`
}

type ReportTableReference struct {
	Base
	TableId               string `json:"tableId" gorm:"type:text;not null;"`
	PrimaryTable          string `json:"primaryTable" gorm:"type:text;not null;"`
	PrimaryTableColumn    string `json:"primaryTableColumn" gorm:"type:text;not null;"`
	ReferencedTable       string `json:"referencedTable" gorm:"type:text;not null;"`
	ReferencedTableColumn string `json:"referencedTableColumn" gorm:"type:text;not null;"`
}

type ReportColumn struct {
	Base
	ReportId   string `json:"reportId" gorm:"type:text;not null;"`
	Name       string `json:"name" gorm:"type:text;not null;"`
	CustomName string `json:"customName" gorm:"type:text;not null;"`
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
	ReportId   string           `json:"reportId" gorm:"type:text;not null;"`
	Type       ReportFilterType `json:"type" gorm:"type:text;not null;"`
	Value      string           `json:"value" gorm:"type:text;not null;"`
	ColumnName string           `json:"columnName" gorm:"type:text;not null;"`
	ColumnType string           `json:"columnType" gorm:"type:text;not null;"`
}
