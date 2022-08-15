package migration

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gopkg.in/gormigrate.v1"
	"gorm.io/datatypes"
	"log"
)

func Migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "12082022_1",
			Migrate: func(db *gorm.DB) error {
				type Pair struct {
					DoneAt datatypes.Date `gorm:"type:timestamp unique"`
					Price  int            `gorm:"type:int;column:price;"`
					Name   string         `gorm:"type:varchar(6)"`
				}
				return db.AutoMigrate(&Pair{}).Error
			},
			Rollback: func(db *gorm.DB) error {
				return db.DropTable("pairs").Error
			},
		},
	})

	if err := m.Migrate(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Migration success")
}
