package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/totoval/framework/database/migration"
	"time"
)

func init() {
	migration.AddMigrator(&UserAffiliation1553678539{})
}

type UserAffiliation struct {
	UserID    *uint      `gorm:"column:user_id;primary_key;type:int unsigned"`
	Code      *string    `gorm:"column:uaff_code;type:varchar(32);not null"`
	FromCode  *string    `gorm:"column:uaff_from_code;type:varchar(32)"`
	Root      *uint      `gorm:"column:uaff_root_id;type:int unsigned"`
	Parent    *uint      `gorm:"column:uaff_parent_id;type:int unsigned"`
	Left      *uint      `gorm:"column:uaff_left_id;type:int unsigned"`
	Right     *uint      `gorm:"column:uaff_right_id;type:int unsigned"`
	Level     *uint      `gorm:"column:uaff_level;type:int unsigned"`
	CreatedAt *time.Time `gorm:"column:user_created_at"`
	UpdatedAt time.Time  `gorm:"column:user_updated_at"`
	DeletedAt *time.Time `gorm:"column:user_deleted_at"`
}

type UserAffiliation1553678539 struct {
	migration.MigratorIdentify
	migration.MigrationUtils
}

func (*UserAffiliation1553678539) Up(db *gorm.DB) *gorm.DB {
	db = db.CreateTable(&UserAffiliation{})
	return db
}

func (*UserAffiliation1553678539) Down(db *gorm.DB) *gorm.DB {
	db = db.DropTableIfExists(&UserAffiliation{})
	return db
}