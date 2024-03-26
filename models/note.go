package models
import (
	"github.com/jinzhu/gorm"
	database "notetaking/database"
)
var DB *gorm.DB
type Note struct{
	gorm.Model
	Title string `json:"title"`
	Details string `json:"details"`
}
func init(){
	database.Connect()
	DB=database.GetDB()
	DB.AutoMigrate(&Note{})
}
