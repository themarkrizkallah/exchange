package users

import (
	"context"
	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
	"zeus/database"
	"zeus/env"
	"zeus/redis"
)

const cookieName = "exchange_userCookie"

func SignUp(c *gin.Context) {
	var payload UserPayload

	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if len(payload.Email) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Email required"})
		return
	} else if len(payload.Password) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Password required"})
		return
	}

	payload.Password, _ = argon2id.CreateHash(payload.Password, argon2id.DefaultParams)
	res, err := InsertUserPayload(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res.InsertedID})
}

func Login(c *gin.Context) {
	var (
		payload LoginPayload
		user    User
	)

	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(payload.UserName) == 0 && len(payload.Email) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "One of {user_name, email} may not be empty"})
		return
	} else if len(payload.Password) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "password may not be empty"})
		return
	}

	user, err = FindUserByLoginPayload(payload)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "User not found"})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(payload.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if !match {
		c.JSON(http.StatusForbidden, gin.H{"error": "Invalid password"})
		return
	}

	key, err := redis.NewUserSession(user.Id.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Cookie variables
	cookieValue := key
	cookieMaxAge := int(time.Hour * 24 * 30)
	cookiePath := "/"
	cookieDomain := "localhost"
	cookieSecure := false
	cookieHttpOnly := false

	c.SetCookie(cookieName, cookieValue, cookieMaxAge, cookiePath, cookieDomain, cookieSecure, cookieHttpOnly)
	c.JSON(http.StatusOK, gin.H{"response": "Login successful"})
}

func ListUsers(c *gin.Context) {
	var users []bson.M

	collection := database.GetDB(env.MongoDb).Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.MongoRetrySeconds)*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.M

		err := cur.Decode(&result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		users = append(users, result)
	}

	c.JSON(http.StatusOK, gin.H{"response": users})
}