package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	})
	<!doctype html>
	<html lang="en">
	<head>
		<!-- Required meta tags -->
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<link rel="stylesheet" href="css/styling">
	
		<title>Vending Machine</title>
	
		<!-- Bootstrap CSS -->
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css" integrity="sha384-HSMxcRTRxnN+Bdg0JdbxYKrThecOKuH5zCYotlSAcp1+c8xmyTe9GYg1l9a69psu" crossorigin="anonymous">
		</head>
		<body>
	
			<div class="container" style="height: 917px; width: 100%;">
	
				<h2 style="text-align:center;">Vending Machine</h2>      
		   
				<!-- Dynamically adding items -->
				<div class="col-md-10" style = "position:relative; left:100px; top:0px; height: 800px; width: 45%;">
					<!-- <div style="background-image: url('machine.jpg'); background-size: 800px 800px; background-repeat: no-repeat;">
						 <div  id="itemTable">
						 </div>
					</div> -->
				   <div id="itemTable" style="background-image: url('machine.jpg'); background-size: 800px 800px; background-repeat: no-repeat;">
				   </div>
				</div>
	
				<!-- Transaction Panel -->
				<div class="col-md-10" style="border: solid; position:relative; left:200px; top:25px; height: 750px; width: 19%;">
		   
					<!-- User Money -->
					<h3>Total $ In</h3>
					<div id="moneyDisplay" class="col-md-2" style="border: solid; position:relative; left:15px; height: 29px; width: 73%;">0.00</div>
					<button id="addDollar" class="button" style="margin:15px;height:50px;width:100px" onclick="getMoney('dollar')">Add Dollar</button>
					<button id="addQuarter" class="button" style="margin:15px;height:50px;width:100px" onclick="getMoney('quarter')">Add Quarter</button>
					<button id="addDime" class="button" style="margin:15px;height:50px;width:100px" onclick="getMoney('dime')">Add Dime</button>
					<button id="addNickel" class="button" style="margin:15px;height:50px;width:100px" onclick="getMoney('nickel')">Add Nickel</button>
	
					<!-- Messages -->
				   <h3>Messages</h3>
					<div id="messages" class="col-md-2" style="border: solid; position:relative; left:5px; top: -10px; height: 50px; width: 73%; margin: 10px;">
						<ul class="list-group" style="position:relative; top:1px; width: 114%; left:-14px;" id="errorMessages"></ul></div>
				   <h3>Selected Item:</h3>
				   <div class="col-md-2"id="inputId" style="border: solid; position:relative; left:5px; height: 30px; width: 73%; margin: 10px;"></div>
				   <button id="purchase" class="button" style="margin:15px;height:50px;width:235px" onclick='transaction()'>Make Purchase</button>
	
				   <!-- Change -->
				   <h3>Change</h3>
				   <div id="change" class="col-md-2" style="border: solid; position:relative; left:5px; height: 50px; width: 73%; margin: 10px;"></div>
				   <button id="makeChange" class="button" style="margin:15px;height:50px;width:235px; display:none;" onclick="changeReturn()" >Change Return</button>
				</div>
	
			</div>
	
		<!-- Scripts -->
		<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
		  
		<script src="home.js"></script>
	
		</body>
		</html>

		var user struct {
			Name  string `json:"name"`
			Phnno string `json:"phnno"`
			Email string `json:"email"`
			City  string `json:"city"`
			State string `json:"state"`
		}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		result, err := db.Exec("Insert into emply(name,phnno,email,city,state) VALUES ($1,$2,$3,$4,$5)",
			user.Name, user.Phnno, user.Email, user.City, user.State)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows affected"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%d rows affected", rowsAffected)})

	})
	router.
		Run(":8081")
}
