package routes
import (
	"net/http"
	//"strconv"
	"example.com/myapp/models"
	"github.com/gin-gonic/gin"
)
func signup(c *gin.Context){
	var newUser models.User;

	err:=c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data you sent!!!"})
		return
	}

	err = newUser.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save User", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user_id": newUser.ID})
}

func login(c *gin.Context){
	var user models.User;

	err:=c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data you sent!!!"})
		return
	}
	err,jwtToken:=user.ValidateCredentials()
	if err!=nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
		
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": jwtToken})
}
