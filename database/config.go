package database
import(
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)
var (
	db *gorm.DB
)
func Connect(){
	d, err := gorm.Open("mysql", "root:hihimel2002@tcp(localhost:3306)/note?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		panic(err)
	}
	db=d
}
func GetDB() *gorm.DB{
	return db
}