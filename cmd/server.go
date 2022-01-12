package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hietkamp/norma-gw/cmd/handlers"
	authserver "github.com/hietkamp/norma-gw/internal/nuts-authserver"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	HttpServer *http.Server
}

func init() {
	// Force log's color
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}
}

// useLoggingMiddleware produces a requests_received stream in kafka
func useLoggingMiddleware(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		// Read from body and write here again.
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		handlers.HandleRequestReceived(c)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Next()
	})
}

// useAuthorizationMiddleware sets the service claims in the context
func useAuthorizationMiddleware(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if strings.HasPrefix(token, "Bearer ") {
			// Get the token from the authorization header
			token = strings.TrimPrefix(token, "Bearer ")
			nutsURL := os.Getenv("nuts_url")
			verifier := authserver.New(authserver.Options{BaseURL: nutsURL})
			claim, err := verifier.VerifyAccessToken(token)
			//log.Debug().Msgf("Result of the verification: %s, with error? %w", claim, err)
			if err != nil {
				// the token is not valid or could not be verified, lets move on
				c.Next()
			}
			// No error means that the token is valid, claim is set in the context
			c.Set("serviceClaim", claim)
		}
		c.Next()
	})
}

func New() *Server {
	router := gin.Default()
	// The order of loading the middleware matters, cors need to be first
	router.Use(cors.Default())
	useLoggingMiddleware(router)
	useAuthorizationMiddleware(router)
	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.POST("/validatedquery", handlers.HandleValidatedQuery)
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	envAddress := os.Getenv("http_server")
	server := Server{
		HttpServer: &http.Server{
			Addr: envAddress,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 60,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      router,
		},
	}
	return &server
}

func (s *Server) Serve() {
	// Initializing the httpserver in a goroutine so that
	// it won't block the graceful shutdown handling
	go func() {
		log.Info().Msgf("HTTP Server Listening on %s", s.HttpServer.Addr)
		if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("Error while listening: %s", err)
		}
	}()
}
