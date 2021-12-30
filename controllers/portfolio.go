package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	// orm "github.com/go-pg/pg/v9/orm"
)

type Coin struct{
  ID string `json:"id"`
  CoinName string `json:"coinName"`
Amount int    `json:"amount"`
Price float32 `json:"price"`
TransactionFee int `json:"transactionFee"`
UserId string  `json:"userId"`
}
type User struct {
	ID     string  `json:"id"`
	UserName  string  `json:"username"`
	
}

// Create User Table
// func CreateUserTable(db *pg.DB) error {
// 	opts := &orm.CreateTableOptions{
// 		IfNotExists: true,
// 	}
// 	createError := db.CreateTable(&User{}, opts)
// 	if createError != nil {
// 		log.Printf("Error while creating todo table, Reason: %v\n", createError)
// 		return createError
// 	}
// 	log.Printf("Todo table created")
// 	return nil
// }


// func CreateCoinTable(db *pg.DB) error{
// opts := &orm.CreateTableOptions{
// 		IfNotExists: true,
// 	}
// 	createError := db.CreateTable(&Coin{}, opts)
// 	if createError != nil {
// 		log.Printf("Error while creating todo table, Reason: %v\n", createError)
// 		return createError
// 	}
// 	log.Printf("Todo table created")
// 	return nil
// }

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func getAllUserCoins(){
	var users []User
err := dbConnect.Model(&users).Relation("coins._").Select()

fmt.Println(err);

}
func GetAllUsers(c *gin.Context) {
	var users []User
	err := dbConnect.Model(&users).Select()

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data": users,
	})
	return
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	username :=user.UserName
	fmt.Println(username)
	

	insertError := dbConnect.Insert(&User{
		UserName:username,
	
		
	})
	if insertError != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created Successfully",
	})
	return
}

func CreateCoin (c *gin.Context) {
	fmt.Println("hello")
    var coin Coin
	c.BindJSON(&coin)

	// id :=c.Param("userId")
	userId:=coin.UserId
	coinName :=coin.CoinName
	amount :=coin.Amount
	price :=coin.Price
	transactionFee :=coin.TransactionFee

	
	fmt.Println(userId)

	insertError := dbConnect.Insert(&Coin{
		UserId: userId,
		CoinName:coinName,
		Amount: amount,
		Price: price,
		TransactionFee: transactionFee,
		
	})
	if insertError != nil {
		log.Printf("Error while inserting new coin into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "coin created Successfully",
	})
	return
}
// func GetSingleTodo(c *gin.Context) {
// 	userId := c.Param("userId")
// 	user := &User{ID: userId}
// 	err := dbConnect.Select(user)

// 	if err != nil {
// 		log.Printf("Error while getting a single todo, Reason: %v\n", err)
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  http.StatusNotFound,
// 			"message": "Todo not found",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Single Todo",
// 		"data": user,
// 	})
// 	return
// }

// func EditTodo(c *gin.Context) {
// 	todoId := c.Param("todoId")
// 	var user User
// 	c.BindJSON(&user)
// 	completed := user.Complete

// 	_, err := dbConnect.Model(&Todo{}).Set("completed = ?", completed).Where("id = ?", todoId).Update()
// 	if err != nil {
// 		log.Printf("Error, Reason: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status": 500,
// 			"message":  "Something went wrong",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"message": "Todo Edited Successfully",
// 	})
// 	return
// }

// func DeleteTodo(c *gin.Context) {
// 	todoId := c.Param("todoId")
// 	todo := &Todo{ID: todoId}

// 	err := dbConnect.Delete(todo)
// 	if err != nil {
// 		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"status":  http.StatusInternalServerError,
// 			"message": "Something went wrong",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Todo deleted successfully",
// 	})
// 	return
// }

