package routes

import (
	"net/http"
	"strconv"
	"twittueur_api/models"
	"twittueur_api/src/utils"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// Posts length

func GlobalPostsLength(c echo.Context) error {
	// Configuration de Viper pour lire le fichier posts.json
	viper.SetConfigName("posts")
	viper.SetConfigType("json")
	viper.AddConfigPath("db")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// Convertir data en une structure go
	var data models.PostRequest
	if err := viper.Unmarshal(&data); err != nil {
		return c.JSON(500, models.Response{Message: "Une erreur s'est produite.", Success: false})
	}

	// Retourner la longueur du slice de posts
	return c.JSON(http.StatusOK, models.Response{Message: strconv.Itoa(len(data.Posts)), Success: true})
}

func GetLikesByPost(c echo.Context) error {
	// Récupérer le token du header
	authorization := c.Request().Header.Get("Authorization")

	if authorization == "" {
		return c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Vous devez renseigner un token."})
	}

	err := utils.IsTokenExists(c, authorization)
	if err != nil {
		return err
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Vous devez renseigner un id."})
	}

	var data models.PostRequest
	if err := viper.Unmarshal(&data); err != nil {
		return err
	}

	// Boucle for pour parcourir les posts
	for _, post := range data.Posts {
		if post.ID == id { // Si l'id du post est égal à l'id donné en paramètre
			return c.JSON(http.StatusOK, models.Response{Success: true, Message: "Nombre de likes récupérer avec succés", Data: strconv.Itoa(len(post.Likedby))})
		}
	}

	return c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Le post n'existe pas."})
}

func GetBookmarksByPost(c echo.Context) error {
	// Récupérer le token du header
	authorization := c.Request().Header.Get("Authorization")

	if authorization == "" {
		return c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Vous devez renseigner un token."})
	}

	err := utils.IsTokenExists(c, authorization)
	if err != nil {
		return err
	}

	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, models.Response{Success: false, Message: "Vous devez renseigner un id."})
	}

	var data models.PostRequest
	if err := viper.Unmarshal(&data); err != nil {
		return err
	}

	// Boucle for pour parcourir les posts
	for _, post := range data.Posts {
		if post.ID == id { // Si l'id du post est égal à l'id donné en paramètre
			return c.JSON(http.StatusOK, models.Response{Success: true, Message: "Nombre de bookmarks récupérer avec succés", Data: strconv.Itoa(len(post.Bookmarkedby))})
		}
	}

	return c.JSON(http.StatusNotFound, models.Response{Success: false, Message: "Le post n'existe pas."})
}
