package potgres

import (
	"hopeugetknowuwont/ungender/config"
	"hopeugetknowuwont/ungender/graph/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pqHandler *psqlHandler

type psqlHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) psqlHandler {
	db.AutoMigrate(&model.Garage{})
	db.AutoMigrate(&model.User{})
	return psqlHandler{DB: db}
}

func Init(config config.Config) (*gorm.DB, error) {
	dbUrl := "postgres://" + config.Psql.User + ":" + config.Psql.Password + "@" + config.Psql.Host + ":" + config.Psql.Sqlport + "/" + config.Psql.Dbname
	DB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logrus.Error("Error in connecting to DB", err)
		return nil, err
	}
	db := New(DB)
	pqHandler = &db
	return db.DB, nil
}

func (h *psqlHandler) GetAll(c *gin.Context) {
	var cars []model.Garage
	if res := h.DB.Find(&cars); res.Error != nil {
		logrus.Error("Error in getting list of cars", res.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"HeadsUp": "Error in getting the list of cars"})
		h.Return(c)
	}
	c.IndentedJSON(http.StatusOK, cars)
}

func (h *psqlHandler) Book(c *gin.Context) {

}

func (h *psqlHandler) Return(c *gin.Context) {

}

func (h *psqlHandler) PostGarage(c *gin.Context) {
	var newcar model.Garage
	if err := c.BindJSON(&newcar); err != nil {
		return
	}
	res := h.DB.Create(&newcar)
	if res.Error != nil {
		logrus.Error("Error in posting a car", res.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"HeadsUp": "Error in posting the car"})
		h.Return(c)
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"HeadsUp": "new car added to the garage"})
}

func (h *psqlHandler) PostUser(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	res := h.DB.Create(&newUser)
	if res.Error != nil {
		logrus.Error("Error in posting a user", res.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"HeadsUp": "Error in posting the user"})
		h.Return(c)
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"HeadsUp": "new user added"})
}

func GetPqHandler() *psqlHandler {
	return pqHandler
}
