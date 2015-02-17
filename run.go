package esa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

	baseUrl := a.c.ElasticSearchHost + "/"
	url := baseUrl + a.c.IndicePrefix + "*"

	// HTTP request to get all indices concerned
	res, err := http.Get(url)
	mustNotErr(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	mustNotErr(err)

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	mustNotErr(err)

	toArchive := []string{}
	toDelete := []string{}
	archiveCutoff := time.Now().Add(-a.c.ArchiveAge)
	deleteCutoff := time.Now().Add(-a.c.DeleteAge)
	for key, _ := range data {
		keyWithoutPrefix := key[len(a.c.IndicePrefix):]
		keyDate, err := time.Parse("2006.01.02", keyWithoutPrefix)
		mustNotErr(err)

		if keyDate.Before(deleteCutoff) {
			toDelete = append(toDelete, key)
		} else if keyDate.Before(archiveCutoff) {
			toArchive = append(toArchive, key)
		}
	}

	log.Printf("target: %s\n", url)
	if a.c.Delete {
		for _, keyToDelete := range toDelete {
			req, err := http.NewRequest("DELETE", baseUrl+keyToDelete, nil)
			mustNotErr(err)
			_, err = http.DefaultClient.Do(req)
			mustNotErr(err)
			log.Println("deleted: " + keyToDelete[len(a.c.IndicePrefix):])
		}
	}
	if a.c.Archive {
		for _, keyToArchive := range toArchive {
			req, err := http.NewRequest("POST", baseUrl+keyToArchive+"/_flush", nil)
			mustNotErr(err)
			_, err = http.DefaultClient.Do(req)
			mustNotErr(err)
			req, err = http.NewRequest("POST", baseUrl+keyToArchive+"/_close", nil)
			mustNotErr(err)
			_, err = http.DefaultClient.Do(req)
			mustNotErr(err)
			log.Println("archived: " + keyToArchive[len(a.c.IndicePrefix):])
		}
	}
}

func mustNotErr(err error) {
	if err != nil {
		panic(err)
	}
}
