package gonati

import (
	"fmt"
	"math/rand"
	"strconv"
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
)

type Client struct {
	Username	string
	Password	string
	Port			int
}

func newClient(username string, password string, port int) Client {
	c := Client{}
	c.Username = username
	c.Password = password
	c.Port = port

	return c
}

func generateProxy(client Client) *url.URL {
	sessionID := strconv.Itoa(rand.Intn(1000000))
	superProxy := fmt.Sprintf("https://%s-session-%s:%s@zproxy.luminati.io:%d", client.Username, sessionID, client.Password, client.Port)

	proxyURL, err := url.Parse(superProxy)
	if err != nil {
		log.Println(err)
	}

	return proxyURL
}

func createClient(proxy *url.URL) *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	return client
}

func proxyClient(username string, password string, port int) *url.URL {
	c := newClient(username, password, port)
	sP := generateProxy(c)

	return sP
}

func CreateProxy(username string, password string, port int) *http.Client {
	c := newClient(username, password, port)
	sP := generateProxy(c)
	client := createClient(sP)

	return client
}

func TestConnection(username string, password string, port int) string {
	client := CreateProxy(username, password, port)

	request, err := http.NewRequest("GET", "https://api.ipify.org?format=json", nil)

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return string(body)
}

