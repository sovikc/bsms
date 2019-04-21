package bitly

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sovikc/bsms/sms"
)

const (
	groupURL   = "https://api-ssl.bitly.com/v4/groups"
	shortenURL = "https://api-ssl.bitly.com/v4/shorten"
)

type urlShortener struct {
	token     string
	groupGUID string
	cache     *cache
}

// NewURLShortener creates a new instance of urlShortener
func NewURLShortener(token, groupGUID string) sms.URLShortener {
	u := &urlShortener{
		token:     token,
		groupGUID: groupGUID,
		cache:     newCache(),
	}

	return u
}

func (u *urlShortener) hasGroupGUID() bool {
	return len(u.groupGUID) > 0
}

func (u *urlShortener) getGroupGUID() (string, error) {
	var groupGUID string
	client := &http.Client{}
	req, err := http.NewRequest("GET", groupURL, nil)
	if err != nil {
		return groupGUID, err
	}
	req.Header.Add("Authorization", "Bearer "+u.token)
	resp, err := client.Do(req)
	if err != nil {
		return groupGUID, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return groupGUID, errors.New("error fetching group_guid")
	}

	type Group struct {
		GUID string `json:"guid"`
	}

	var groups struct {
		Groups []Group `json:"groups"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&groups); err != nil {
		return groupGUID, err
	}

	groupGUID = groups.Groups[0].GUID
	return groupGUID, nil
}

func (u *urlShortener) GetShortenedURL(longURL string) (string, error) {

	shortURL, found := u.cache.get(longURL)
	if found {
		log.Println("found " + shortURL + " in cache")
		return shortURL, nil
	}

	if !u.hasGroupGUID() {
		u.getGroupGUID()
	}

	var sb strings.Builder
	sb.WriteString(`{"long_url":"`)
	sb.WriteString(longURL)
	sb.WriteString(`","group_guid":"`)
	sb.WriteString(u.groupGUID)
	sb.WriteString(`"}`)

	var payload = []byte(sb.String())

	client := &http.Client{}
	req, err := http.NewRequest("POST", shortenURL, bytes.NewBuffer(payload))
	if err != nil {
		return shortURL, err
	}

	req.Header.Set("Authorization", "Bearer "+u.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return shortURL, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case 400:
		return shortURL, errors.New(strconv.Itoa(resp.StatusCode) + " BAD_REQUEST")
	case 403:
		return shortURL, errors.New(strconv.Itoa(resp.StatusCode) + " FORBIDDEN")
	case 404:
		return shortURL, errors.New(strconv.Itoa(resp.StatusCode) + " NOT_FOUND")
	case 422:
		return shortURL, errors.New(strconv.Itoa(resp.StatusCode) + " UNPROCESSABLE_ENTITY")
	case 500:
		return shortURL, errors.New(strconv.Itoa(resp.StatusCode) + " INTERNAL_ERROR")
	case 503:
		return shortURL, errors.New(strconv.Itoa(resp.StatusCode) + " TEMPORARILY_UNAVAILABLE")
	}

	var shortenedURL struct {
		Link    string `json:"link"`
		Message string `json:"message"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&shortenedURL); err != nil {
		return shortURL, err
	}

	shortURL = shortenedURL.Link
	u.cache.add(longURL, shortURL)
	return shortURL, nil
}
