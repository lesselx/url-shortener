package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	data "url.shortener.agnes/internal/model"
)

type GenericResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type InputRequest struct {
	OriginUrl *string `json:"origin_url" binding:"required"`
}

type LinkResponse struct {
	ShortUrl *string   `json:"short_url"`
	Created  time.Time `json:"created_at"`
}

type RedirectLink struct {
	Alias *string `uri:"alias" binding:"required"`
}

const (
	keyLength = 6
	charset   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func (app *application) GenerateLink(c *gin.Context) {
	var json InputRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alias, err := app.generateAlias()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	link := data.Link{
		OriginUrl: json.OriginUrl,
		Alias:     &alias,
	}

	tx := app.db.Create(&link)

	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}

	shortenUrl := fmt.Sprintf("%s:%d%s/%s", app.config.host, app.config.port, generateLink, alias)

	c.JSON(http.StatusCreated,
		GenericResponse{
			StatusCode: http.StatusCreated,
			Data: LinkResponse{
				ShortUrl: &shortenUrl,
				Created:  link.CreatedAt,
			},
		},
	)

}

func (app *application) RedirectLink(c *gin.Context) {

	var alias RedirectLink
	if err := c.ShouldBindUri(&alias); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	var link data.Link

	res := app.db.First(&link, "alias = ?", alias.Alias)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		}
		return
	}

	c.Redirect(http.StatusMovedPermanently, *link.OriginUrl)

}

func (app *application) generateAlias() (string, error) {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var link data.Link

	for {
		shortKey := make([]byte, keyLength)
		for i := 0; i < keyLength; i++ {
			randomIdx := seededRand.Intn(len(charset))
			shortKey[i] = charset[randomIdx]
		}

		err := app.db.First(&link, "alias = ?", shortKey).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return string(shortKey), nil

		} else if err != nil {
			log.Printf("database error: %v", err)
			return "", err
		}

	}

}
