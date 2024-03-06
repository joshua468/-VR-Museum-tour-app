package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Exhibit struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Rotation string `json:"rotation"`
	Model    string `json:"model"`
	Scale    string `json:"scale"`
}

var exhibits = []Exhibit{
	{Name: "Ancient Sculpture", Position: "-5 0 -10", Rotation: "0 180 0", Model: "models/sculpture.glb", Scale: "0.5 0.5 0.5"},
	{Name: "Modern Art", Position: "5 0 -10", Rotation: "0 0 0", Model: "models/modern_art.glb", Scale: "0.3 0.3 0.3"},
	{Name: "Olumorock", Position: "10 0 -15", Rotation: "0 270 0", Model: "models/olumorock.glb", Scale: "0.6 0.6 0.6"},
	{Name: "Cape Coast Castle", Position: "15 0 -20", Rotation: "0 90 0", Model: "models/cape_coast_castle.glb", Scale: "0.4 0.4 0.4"},
	{Name: "Table Mountain", Position: "20 0 -25", Rotation: "0 180 0", Model: "models/table_mountain.glb", Scale: "0.8 0.8 0.8"},
	{Name: "Maasai Mara", Position: "25 0 -30", Rotation: "0 90 0", Model: "models/maasai_mara.glb", Scale: "0.7 0.7 0.7"},
	{Name: "Pyramids of Giza", Position: "30 0 -35", Rotation: "0 0 0", Model: "models/pyramids.glb", Scale: "1.0 1.0 1.0"},
	{Name: "Hassan II Mosque", Position: "35 0 -40", Rotation: "0 270 0", Model: "models/hassan_ii_mosque.glb", Scale: "0.5 0.5 0.5"},
}

var JWTSecret = []byte("bRBKbsR4NE05oYDFoxu8M1JWCZ4xvsCofhcsd1EReFs")

func main() {
	r := gin.Default()

	r.POST("/login", loginHandler)
	r.GET("/exhibits", authMiddleware(), getExhibitsHandler)

	r.StaticFS("/static", http.Dir("static"))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func loginHandler(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "user" && password == "password" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = username
		tokenString, _ := token.SignedString(JWTSecret)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func getExhibitsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"exhibits": exhibits})
}
