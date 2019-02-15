package migrations

import (
	"github.com/jinzhu/gorm"
	"time"
	"totoval-framework/database/migration"
)

func init() {
	migration.AddMigrator(&UserVerification1548750785{})
}

type UserVerification1548750785 struct {
	migration.MigratorIdentify
	migration.MigrationUtils
}

type UserVerification struct {
	ID                       uint       `gorm:"column:uv_id;primary_key;auto_increment"`
	UserID                   uint       `gorm:"column:user_id;index;not null"`
	CountryName              string     `gorm:"column:uv_country_name;type:varchar(100);not null"`
	ProvinceName             string     `gorm:"column:uv_province_name;type:varchar(100);not null"`
	CityName                 string     `gorm:"column:uv_city_name;type:varchar(100);not null"`
	DistrictName             string     `gorm:"column:uv_district_name;type:varchar(100)"`
	Address                  string     `gorm:"column:uv_address;type:text"`
	LicenseName              string     `gorm:"column:uv_license_name;type:varchar(100)"`
	LicenseNumber            string     `gorm:"column:uv_license_number;type:varchar(100)"`
	LicensePhotoUrl          string     `gorm:"column:uv_license_photo_url;type:longtext"`
	LegalPersonName          string     `gorm:"column:uv_legal_person_name;type:varchar(100);not null"`
	LegalPersonIDCardNumber  string     `gorm:"column:uv_legal_person_id_card_number;type:varchar(100);not null"`
	LegalPersonFrontPhotoUrl string     `gorm:"column:uv_legal_person_front_photo_url;type:longtext"`
	LegalPersonBackPhotoUrl  string     `gorm:"column:uv_legal_person_back_photo_url;type:longtext"`
	UVType                   string     `gorm:"column:uv_type;type:enum('PERSONAL','COMPANY');not null"`
	Status                   string     `gorm:"column:uv_status;type:enum('PENDING','REJECTED','PASSED');default:'PENDING';not null"`
	CreatedAt                time.Time  `gorm:"column:uv_created_at"`
	UpdatedAt                time.Time  `gorm:"column:uv_updated_at"`
	DeletedAt                *time.Time `gorm:"column:uv_deleted_at"`
}

func (*UserVerification1548750785) Up(db *gorm.DB) {
	db.CreateTable(&UserVerification{})
}

func (*UserVerification1548750785) Down(db *gorm.DB) {
	db.DropTableIfExists(&UserVerification{})
}