package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "123456"
  dbname   = "gotest"
)
type Coin struct{
  ID string `json:"id"`
  CoinName string `json:"coinName"`
Amount int    `json:"amount"`
Price float32 `json:"price"`
TransactionFee int `json:"transactionFee"`
}
type User struct {
	ID     string  `json:"id"`
	UserName  string  `json:"username"`
	
}

var Users = []User{
	{ID: "1",  UserName: "John Coltrane"},
	{ID: "2",  UserName: "Gerry Mulligan"},
	{ID: "3",  UserName: "Sarah Vaughan"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Users)
}


func postCoin(c *gin.Context) {
    var newCoin Coin

    // Call BindJSON to bind the received JSON to
    // newCoin.
    if err := c.BindJSON(&newCoin); err != nil {
        return
    }

    // Add the new newCoin to the slice.
    // Users = append(Users, newCoin)
    c.IndentedJSON(http.StatusCreated, newCoin)
}

func postPortfolio(c *gin.Context){
  var newUser User

   sqlStatement := `
INSERT INTO users (age, email, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING id`
  id := 0
  err = db.QueryRow(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun").Scan(&id)
  if err != nil {
    panic(err)
  }
  fmt.Println("New record ID is:", id)
 if err := c.BindJSON(&newUser); err != nil {
        return
    }
    Users = append(Users, newUser)
      c.IndentedJSON(http.StatusCreated, newUser)
}
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range Users {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
func main() {
	 psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
    router := gin.Default()
    router.GET("/portfolios", getUsers)
    router.POST("/portfolio",postPortfolio)
	 router.GET("/portfolio/:id", getAlbumByID)
 router.POST("/portfolio/:id/entry", postCoin)
    router.Run("localhost:8080")
}

