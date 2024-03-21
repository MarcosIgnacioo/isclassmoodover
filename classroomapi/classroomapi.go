package classroomapi

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/classroom/v1"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func parseJsonToOauthToken(token *string) *oauth2.Token {
	data := oauth2.Token{}
	json.Unmarshal([]byte(*token), &data)
	return &data
}

func getClient(config *oauth2.Config, credentials *string) (*http.Client, *string) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	// tok, err := tokenFromFile(tokFile)
	var tok *oauth2.Token
	if credentials != nil {
		tok = parseJsonToOauthToken(credentials)
	} else {
		tok = getTokenFromWeb(config)
		// Aqui es donde se abriria playwright y empezaria a hacer su despapaye de si acepto si
	}
	json, err := json.Marshal(tok)
	if err != nil {
		log.Fatal(err)
	}
	tokenToStore := string(json)
	return config.Client(context.Background(), tok), &tokenToStore
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	// playwright se inicia aqui y empieza a darle si a todo hasta que llega a la ultima pagina y consigue el token
	// TODO puedo hacer una pagina en githubpages que obtenga de la url la clave y la ponga ahi bonito en un parrafito
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func GetOAuthURL(tokenJSON *string) {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, classroom.ClassroomCoursesReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client, tokenAPI := getClient(config, tokenJSON)

	fmt.Println(tokenAPI)

	srv, err := classroom.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create classroom Client %v", err)
	}
	r, err := srv.Courses.List().PageSize(10).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve courses. %v", err)
	}
	if len(r.Courses) > 0 {
		fmt.Print("Courses:\n")
		for _, c := range r.Courses {
			fmt.Printf("%s (%s)\n", c.Name, c.Id)
		}
	} else {
		fmt.Print("No courses found.")
	}
}
