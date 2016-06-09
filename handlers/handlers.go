package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiborv/hilbert-geohash/geohash/hash"
	"github.com/tiborv/hilbert-geohash/geohash/point"
)

func Register(r *gin.Engine) {
	api(r.Group("/api"))

	r.GET("/", showIndex)
	r.GET("/h/:_", showIndex)
	r.GET("/p/:_", showIndex)

}

func showIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Hilbert GeoHash",
	})
}

func api(api *gin.RouterGroup) {
	api.POST("/point", func(c *gin.Context) {
		p := point.Point{}
		err := c.Bind(&p)
		hash, err := hash.NewHashPoint(p)
		if err != nil || !p.IsValid() {
			c.JSON(http.StatusBadRequest, "Invalid point")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"point":  p,
				"hash":   reverseString(hash.String),
				"zorder": reverseString(hash.GetZorder().HashString(32)),
			})
		}
	})

	api.POST("/hash", func(c *gin.Context) {
		hash := &hash.Hash{}
		err := c.Bind(&hash)
		hash.String = reverseString(hash.String)
		if err != nil || hash.InitFromString() != nil {
			c.JSON(http.StatusBadRequest, "Invalid hash")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"hash":   reverseString(hash.String),
				"point":  hash.GenPoint(),
				"zorder": reverseString(hash.GetZorder().HashString(32)),
			})
		}
	})
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
