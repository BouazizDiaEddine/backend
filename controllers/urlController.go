package controllers

import (
	"backend/initialazers"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
	"strings"
)

type bodyUrl struct {
	Url       string
	Operation models.Operation
}

type resultToReturn struct {
	Processed_url string
}

func isURL(u *url.URL) string {

	if u.Scheme == "" {
		return "this url has no Scheme"
	} else if u.Host == "" {
		return "this url has no host"
	}
	return ""
}

func UrlProcess(c *gin.Context) {
	var body bodyUrl
	var treatmentResult models.Url
	var toReturn resultToReturn
	err := c.Bind(&body)
	if err != nil {
		log.Fatal("could not bind")
		return
	}

	treatmentResult.UrlGiven = body.Url

	parsedUrl, err := url.Parse(body.Url)

	urlError := isURL(parsedUrl)
	if urlError != "" {
		c.JSON(200, gin.H{
			"result": urlError,
		})
		return
	}
	if err != nil {
		log.Fatal("error parsing the Url")
		return
	}

	switch body.Operation {
	case models.Redirected:
		log.Println("redirect " + body.Url)
		treatmentResult.Operation = models.Redirected
		treatmentResult.UrlTreated = canonicalURL(parsedUrl).String()
	case models.Canonical:
		log.Println("Canonicalize " + body.Url)
		treatmentResult.Operation = models.Canonical
		treatmentResult.UrlTreated = redirectedURL(parsedUrl).String()
	case models.All:
		log.Println("all operations " + body.Url)
		treatmentResult.Operation = models.All
		treatmentResult.UrlTreated = processURL(parsedUrl).String()
	default:
		fmt.Println("Operation unknown")
		return
	}

	result := initialazers.DB.Create(&treatmentResult)

	if result.Error != nil {
		log.Fatal("could not create post")
		c.Status(400)
	}

	toReturn.Processed_url = treatmentResult.UrlTreated

	c.JSON(200, gin.H{
		"result": toReturn,
	})

}

func canonicalURL(u *url.URL) *url.URL {
	u.RawQuery = ""
	u.Path = strings.TrimSuffix(u.Path, "/")
	return u
}

func redirectedURL(u *url.URL) *url.URL {
	u.Host = "www.byfood.com"
	u.Path = strings.ToLower(u.Path)
	return u
}

func processURL(u *url.URL) *url.URL {
	return redirectedURL(canonicalURL(u))
}
