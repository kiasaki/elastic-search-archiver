package esa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ElasticSearchArchiver struct {
	c Config
}

func New(c Config) ElasticSearchArchiver {
	return ElasticSearchArchiver{c}
}

// Calls ElasticSearch with the prefix to figure out all collections concerned
// then will filter the one that will be deleted and the ones that will be
// archived then call the appropriate ElasticSearch endpoints to make that
// happen while logging what it does.
func (a ElasticSearchArchiver) Run() {
	err := a.c.Validate()
	if err != nil {
		log.Fatal(err)
	}

	url := a.c.ElasticSearchHost + "/" + a.c.IndicePrefix + "*"

	// HTTP request to get all indices concerned
	res, err := http.Get(url)
	mustNotErr(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	mustNotErr(err)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	mustNotErr(err)

	for key, _ := range data {
		log.Println(key)
	}
}

func mustNotErr(err error) {
	if err != nil {
		panic(err)
	}
}
