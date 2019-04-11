package pq

import (
	"database/sql"
	"net/url"
	"rilutham/stovia/lib/log"
	"rilutham/stovia/lib/utils"
	"sync"

	_ "github.com/lib/pq"
	ql "github.com/nleof/goyesql"
)

const (
	storeAssetTag            = "resources/sql/1.store.sql"
	batchAssetTag            = "resources/sql/2.batch.sql"
	mappingAssetTag          = "resources/sql/3.mapping.sql"
	priceLogAssetTag         = "resources/sql/4.price_log.sql"
	priceAssetTag            = "resources/sql/5.price.sql"
	markupAssetTag           = "resources/sql/6.markup.sql"
	assortmentAssetTag       = "resources/sql/7.assortment.sql"
	internalPriceLogAssetTag = "resources/sql/8.internal_price_log.sql"
	stockLocationAssetTag    = "resources/sql/9.stock_location.sql"
	priceSourceAssetTag      = "resources/sql/10.price_source.sql"
)

var (
	db     *sql.DB
	dbLock sync.Mutex
)

func tx(f func(*sql.Tx) error) error {
	t, err := db.Begin()

	if err != nil {
		log.For("PQ", "pq").Error(err)
		return err
	}

	err = f(t)

	if err != nil {
		log.For("PQ", "pq").Error(err)
		t.Rollback()
		return err
	}

	t.Commit()
	return nil
}

func Open(dsn string) (err error) {
	dbLock.Lock()
	d, _ := url.Parse(dsn)
	d.User = url.UserPassword(d.User.Username(), "-FILTERED-")
	log.For("PQ", "pq").Infof("Opening database on %s", d.String())
	db, err = sql.Open("postgres", dsn)
	dbLock.Unlock()

	if err != nil {
		log.For("PQ", "pq").Error(err)
		return err
	}

	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(5)
	return nil
}

func loadQueries(asset string) ql.Queries {
	data, err := utils.Asset(asset)
	if err != nil {
		log.For("Encoding", "zip init").Fatalf("PQ: Resources not found (forgotten import?), for %s - error %s", asset, err.Error())
	}

	queries, err := ql.ParseBytes(data)
	if err != nil {
		log.For("Encoding", "zip init").Fatalf("PQ: Resources failed on parse for %s - error %s", asset, err.Error())
	}

	return queries
}
