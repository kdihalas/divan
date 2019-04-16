package provider

import (
	"net/url"
	"time"

	"github.com/rhinoman/couchdb-go"
	log "github.com/sirupsen/logrus"
)

type Couchdb struct {
	config map[string]interface{}
	db *couchdb.Database
}

func (c *Couchdb) connect() {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection(c.config["server"].(string),c.config["port"].(int),timeout)
	if err != nil {
		log.Error(nil)
	}
	auth := couchdb.BasicAuth{Username: c.config["username"].(string), Password: c.config["password"].(string) }
	c.db = conn.SelectDB(c.config["database"].(string), &auth)
}

func (c *Couchdb) Update(docid string, doc interface{}) error {
	var old interface{}
	rev, err := c.db.Read(docid, &old, &url.Values{})
	if err != nil {
		_, err := c.db.Save(doc, docid, "")
		return err
	} else {
		_, err := c.db.Save(doc, docid, rev)
		return err
	}
}

func NewCouchDbProvider(config map[string]interface{}) *Couchdb{
	couch := &Couchdb{
		config: config,
	}
	couch.connect()
	return couch
}