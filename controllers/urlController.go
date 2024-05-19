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
	if err != nil {
		log.Fatal("error parsing the Url")
		return
	}

	switch body.Operation {
	case models.Redirected:
		log.Println("redirect " + body.Url)
		treatmentResult.Operation = models.Redirected
		canonicalURL(parsedUrl)
	case models.Canonical:
		fmt.Println("Canonicalize " + body.Url)
		treatmentResult.Operation = models.Canonical
		treatmentResult.UrlTreated = redirectedURL(parsedUrl)
	case models.All:
		fmt.Println("all operations " + body.Url)
		treatmentResult.Operation = models.All
		treatmentResult.UrlTreated = processURL(parsedUrl)
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

func canonicalURL(u *url.URL) string {
	u.RawQuery = ""                          // Remove query parameters
	u.Path = strings.TrimSuffix(u.Path, "/") // Remove trailing slash
	return u.String()
}

// Method for ensuring the domain is www.byfood.com and converting the entire URL to lowercase
func redirectedURL(u *url.URL) string {
	u.Host = "www.byfood.com"        // Ensure domain is www.byfood.com
	u.Path = strings.ToLower(u.Path) // Convert entire URL path to lowercase
	return u.String()
}

// Method for conducting both cleaning up and redirection operations on the URL
func processURL(u *url.URL) string {
	canonicalURL(u)
	return redirectedURL(u)
}
