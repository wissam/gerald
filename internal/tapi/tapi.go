package tapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// TODO
// DB :( ... this needs to be stored in DB after unmarshalling... FIGURE IT
// OUT!
// yep, I did not migrate this to models...
type EmoteResponse struct {
	Data     []EmoteData `json:"data"`
	Template string      `json:"template"`
}
type EmoteImages struct {
	URL1X string `json:"url_1x"`
	URL2X string `json:"url_2x"`
	URL4X string `json:"url_4x"`
}
type EmoteData struct {
	EmoteSetID string      `json:"emote_set_id"`
	EmoteType  string      `json:"emote_type"`
	Format     []string    `json:"format"`
	ID         string      `json:"id"`
	Images     EmoteImages `json:"images"`
	Name       string      `json:"name"`
	Scale      []string    `json:"scale"`
	ThemeMode  []string    `json:"theme_mode"`
	Tier       string      `json:"tier"`
}

type UserResponse struct {
	Data []UserData `json:"data"`
}
type UserData struct {
	ID              string    `json:"id"`
	Login           string    `json:"login"`
	DisplayName     string    `json:"display_name"`
	Type            string    `json:"type"`
	BroadcasterType string    `json:"broadcaster_type"`
	Description     string    `json:"description"`
	ProfileImageURL string    `json:"profile_image_url"`
	OfflineImageURL string    `json:"offline_image_url"`
	ViewCount       int       `json:"view_count"`
	Email           string    `json:"email"`
	CreatedAt       time.Time `json:"created_at"`
}

type parampair struct {
	key   string
	value string
}

//Universal Twitch GET query
func TGQuery(turl string, pairs ...parampair) []byte {
	client := &http.Client{}
	parm := url.Values{}
	for _, pair := range pairs {
		parm.Add(pair.key, pair.value)
	}
	req, err := http.NewRequest("GET", turl, nil)
	if err != nil {
		log.Println(err.Error())
	}
	req.URL.RawQuery = parm.Encode()
	var bearer = "Bearer " + os.Getenv("CLIENTTOKEN")
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Client-Id", os.Getenv("CLIENTID"))
	resp, cerr := client.Do(req)
	if cerr != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

//Get User by name, gives back its ID. most other queries need the id as a
//query param
func GetUserID(Name string) {
	pairs := []parampair{{"login", Name}}
	var user UserResponse
	body := TGQuery("https://api.twitch.tv/helix/users", pairs...)
	json.Unmarshal(body, &user)
	var users []string
	for _, dat := range user.Data {
		users = append(users, dat.ID)
	}
}

// Getting all emotes for a channel, this is to classify if there are emotes in
// the db that "expired"
func GetAllEmotes(broadcaster_id string) []string {
	pairs := []parampair{{"broadcaster_id", broadcaster_id}}
	var emote EmoteResponse
	body := TGQuery("https://api.twitch.tv/helix/chat/emotes", pairs...)
	json.Unmarshal(body, &emote)
	var emotes []string //I need to work on the variable naming..confusing AF..later..
	for _, dat := range emote.Data {
		emotes = append(emotes, dat.ID)
	}
	return emotes
}
