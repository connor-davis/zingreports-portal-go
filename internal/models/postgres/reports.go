package postgres

type Report struct {
	Base
	Name  string      `json:"name" gorm:"type:text;not null;"`
	Table ReportTable `json:"table" gorm:"foreignKey:ReportId;references:Id;constraint:onDelete:CASCADE;"`
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
