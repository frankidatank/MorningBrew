package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type post struct {
	Source   string `json:"source"`
	Title    string `json:"title"`
	PostUrl  string `json:"url"`
	ImageUrl string `json:"image_url"`
}

func getDevNews() []post {
	var posts []post
	morningBrew := getMorningBrew()
	posts = append(posts, morningBrew...)
	hashNode := getHashNode()
	posts = append(posts, hashNode...)
	//TODO: finish this function
	//getDailyDev()
	//g.IndentedJSON(http.StatusOK, posts)
	return posts
}

func getHashNode() []post {
	c := colly.NewCollector()
	var posts []post
	c.OnHTML("body div[id=__next] div[class=css-icgnyh] h1 a", func(e *colly.HTMLElement) {
		post := post{
			Source:  "HashNode",
			Title:   e.Text,
			PostUrl: e.Attr("href"),
		}
		posts = append(posts, post)
	})
	c.Visit("https://hashnode.com/")
	return posts
}

func getMorningBrew() []post {
	c := colly.NewCollector()

	var posts []post

	c.OnHTML("div[id=content] li", func(e *colly.HTMLElement) {
		post := post{
			Source:  "MorningBrew",
			Title:   e.Text,
			PostUrl: e.ChildAttr("a", "href"),
		}
		posts = append(posts, post)
	})
	c.Visit("https://blog.cwa.me.uk/2022/11/08/the-morning-brew-3581/")
	return posts
}

// TODO: find out the query to select the posts
// func getDailyDev() []post {
// 	c := colly.NewCollector()

// 	var posts []post
// 	c.OnHTML(".Card_link__RyP56 absolute inset-0 w-full h-full focus-outline", func(e *colly.HTMLElement) {
// 		fmt.Println(e.Text)
// 	})
// 	c.Visit("https://app.daily.dev/")

// 	return posts
// }

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(g *gin.Context) {
		posts := getDevNews()
		// jsonPosts, err := json.Marshal(posts)
		// if err != nil {
		// 	fmt.Printf("Error: %s", err.Error())
		// }
		//fmt.Println((jsonPosts))
		g.HTML(http.StatusOK, "index.tmpl", gin.H{
			"posts": posts,
		})
	})
	//router.GET("/morning-dev", getDevNews)
	router.Run("localhost:8080")

}
