package middle

import (
	"encoding/json"
	"growdo/src/helpers/componen"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GenerateToken(user int, inter interface{}) string {
	var (
		jwtKey = componen.GodotEnv("JWT_SECRET")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":     user,
		"userId": user,
		"user":   inter,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(jwtKey))
	return tokenString
}

type DataUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles int    `json:"roles"`
}

func IsLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var resp = map[string]interface{}{"status": false, "message": "Missing authorization token"}

		clientToken := c.Request().Header.Get("Authorization")
		if clientToken == "" {
			resp["message"] = "No Authorization header provided"
			return c.JSON(401, resp)
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			resp["message"] = "Incorrect Format of Authorization Token"
			return c.JSON(401, resp)
		}

		token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(componen.GodotEnv("JWT_SECRET")), nil
		})
		if err != nil {
			resp["message"] = "Invalid token, please login"
			return c.JSON(401, resp)
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		user := DataUser{}
		data, _ := json.Marshal(claims["user"])
		json.Unmarshal(data, &user)

		c.Set("id", user.Id)
		c.Set("name", user.Name)
		c.Set("email", user.Email)
		c.Set("roles", user.Roles)

		return next(c)
	}
}

func MiddelWareAdmin(T int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var resp = map[string]interface{}{"status": false, "message": "Missing authorization token"}

			clientToken := c.Request().Header.Get("Authorization")
			if clientToken == "" {
				resp["message"] = "No Authorization header provided"
				return c.JSON(401, resp)
			}

			extractedToken := strings.Split(clientToken, "Bearer ")

			if len(extractedToken) == 2 {
				clientToken = strings.TrimSpace(extractedToken[1])
			} else {
				resp["message"] = "Incorrect Format of Authorization Token"
				return c.JSON(401, resp)
			}

			token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
				return []byte(componen.GodotEnv("JWT_SECRET")), nil
			})
			if err != nil {
				resp["message"] = "Invalid token, please login"
				return c.JSON(401, resp)
			}

			claims, _ := token.Claims.(jwt.MapClaims)

			user := DataUser{}
			data, _ := json.Marshal(claims["user"])
			json.Unmarshal(data, &user)

			c.Set("id", user.Id)
			c.Set("name", user.Name)
			c.Set("email", user.Email)
			c.Set("roles", user.Roles)

			if T != user.Roles {
				resp["message"] = "don't have Permission in routes"
				return c.JSON(401, resp)
			}
			return next(c)
		}
	}
}
