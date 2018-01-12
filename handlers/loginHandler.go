package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/danilopolani/gocialite.v0"
)

var gocial = gocialite.NewDispatcher()


//LoginIndex Show homepage with login URL
func LoginIndex(c *gin.Context) {
	c.Writer.Write([]byte("<html><title>Wanderlust - travel, trek, explore</title>" +
		"<head>Wanderlust - travel, trek, explore</head><body>" +
		"<a href='/auth/github'><button>Login with GitHub</button></a><br>" +
		"<a href='/auth/linkedin'><button>Login with LinkedIn</button></a><br>" +
		"<a href='/auth/facebook'><button>Login with Facebook</button></a><br>" +
		"<a href='/auth/google'><button>Login with Google</button></a><br>" +
		"<a href='/auth/bitbucket'><button>Login with Bitbucket</button></a><br>" +
		"<a href='/auth/amazon'><button>Login with Amazon</button></a><br>" +
		"<a href='/auth/amazon'><button>Login with Slack</button></a><br>" +
		"</body></html>"))
}

//LoginRedirect redirects to correct oAuth URL
func LoginRedirect(c *gin.Context) {
	// Retrieve provider from route
	provider := c.Param("provider")

	host := c.Request.Host
	fmt.Printf("%#v", host)

	// In this case we use a map to store our secrets, but you can use dotenv or your framework configuration
	// for example, in revel you could use revel.Config.StringDefault(provider + "_clientID", "") etc.
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9090/auth/github/callback",
		},
		"linkedin": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9090/auth/linkedin/callback",
		},
		"facebook": {
			"clientID":     "140361796677490",
			"clientSecret": "bdbd0ad12b644305545457c7b8532a71",
			"redirectURL":  "http://" + host + "/auth/facebook/callback",
		},
		"google": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9090/auth/google/callback",
		},
		"bitbucket": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9090/auth/bitbucket/callback",
		},
		"amazon": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9090/auth/amazon/callback",
		},
		"slack": {
			"clientID":     "xxxxxxxxxxxxxx",
			"clientSecret": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			"redirectURL":  "http://localhost:9090/auth/slack/callback",
		},
	}

	providerScopes := map[string][]string{
		"github":    []string{"public_repo"},
		"linkedin":  []string{},
		"facebook":  []string{},
		"google":    []string{},
		"bitbucket": []string{},
		"amazon":    []string{},
		"slack":     []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL)
}

//LoginCallback handles callback of the provider
func LoginCallback(c *gin.Context) {
	// Retrieve query params for state and code
	state := c.Query("state")
	code := c.Query("code")
	//provider := c.Param("provider")

	// Handle callback and check for errors
	user, token, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	// Print in terminal user information
	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)

	// If no errors, show provider name
	c.Writer.Write([]byte("Hi, " + user.FullName))
}