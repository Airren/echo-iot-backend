package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/airren/echo-iot-backend/model"
)

var db *gorm.DB

func InitMySQL() error {
	var (
		err error
	)
	dsn := "root:1q2w3e4r5t@tcp(www.bytegopher.com:3307)/echo_iot_cloud?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	//db.SingularTable(true)
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Order{})
	if err != nil {
		return err
	}

	//db.Model(&model.Order{}).AddIndex("idx_task_name", "name")
	//db.Model(&model.Order{}).AddIndex("idx_task_name_claimer_operator", "name", "claimer", "operator")
	//db.Model(&model.Order{}).AddIndex("idx_task_priority", "priority")
	//db.Model(&model.Order{}).AddIndex("idx_task_status", "status")
	//db.Model(&model.Order{}).AddIndex("idx_task_group_id", "group_id")
	//db.Model(&model.Order{}).AddIndex("idx_task_created_at", "created_at")

	db.Debug()

	return nil
}
