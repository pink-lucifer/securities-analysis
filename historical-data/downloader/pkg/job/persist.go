package job

import (
	"github.com/jmoiron/sqlx"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/conf"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
	"go.uber.org/atomic"
	"log"
	"sync"
)

var sqlDb *sqlx.DB

func initDB(config *conf.Config) {
	persist.CompleteConfig(config.DataSourceConfig)
	sqlDb = persist.GetDb()
}

func persistEodPrice(db *sqlx.DB, listed <-chan *persist.TblEodPrice, eofCh <-chan int, done chan<- bool) {
	var total, processed int
	for {
		select {
		case price := <-listed:
			processed++
			log.Println(price)
			tx, err := db.Beginx()
			if err != nil {
				log.Println(err)
				continue
			}
			id, ra, err := persist.InsertEodPrice(tx, price)
			if err != nil {
				log.Println(err)
				tx.Rollback()
				continue
			}
			log.Printf("Persisted a new eod price %s, id %d, row affected: %d", price.Symbol, id, ra)
			tx.Commit()
		case eof := <-eofCh:
			total = eof
		default:
			if total != 0 && processed != 0 && total == processed {
				done <- true
				return // differences between break and return
			}
		}
	}
}

func pooledPersistEodPrice(db *sqlx.DB, listed <-chan *persist.TblEodPrice, eofCh <-chan int, done chan<- bool) {
	var wg sync.WaitGroup
	var total, processed, skipped atomic.Int64
	for worker := 0; worker < 20; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case price := <-listed:
					tx, err := db.Beginx()
					if err != nil {
						log.Printf("pooledPersistEodPrice error: %s, symbol %s.\n", err.Error(), price.Symbol)
						processed.Inc()
						skipped.Inc()
						continue
					}
					id, ra, err := persist.InsertEodPrice(tx, price)
					if err != nil {
						log.Printf("pooledPersistEodPrice error: %s, symbol %s.\n", err.Error(), price.Symbol)
						tx.Rollback()
						processed.Inc()
						skipped.Inc()
						continue
					}
					log.Printf("Persisted a new eod price %s, id %d, row affected: %d", price.Symbol, id, ra)
					tx.Commit()
					processed.Inc()
				case eof := <-eofCh:
					total.Store(int64(eof))
				default:
					if total.Load() != 0 && processed.Load() != 0 && total.Load() == processed.Load() {
						return // differences between break and return
					}
				}
			}
		}()
	}
	wg.Wait()
	log.Printf("pooledPersistEodPrice done, total:%d, processed:%d, skipped:%d.\n", total.Load(), processed.Load(), skipped.Load())
	done <- true
}

func persistListedSymbol(db *sqlx.DB, listed <-chan *persist.TblListedSymbol, eofCh <-chan int, done chan<- bool) {
	var total, processed int
	for {
		select {
		case symbol := <-listed:
			processed++
			log.Println(symbol)
			tx, err := db.Beginx()
			if err != nil {
				log.Println(err)
				continue
			}
			id, ra, err := persist.InsertListedSymbol(tx, symbol)
			if err != nil {
				log.Println(err)
				tx.Rollback()
				continue
			}
			log.Printf("Persisted a new symbol %s, id %d, row affected: %d", symbol.Symbol, id, ra)
			tx.Commit()
		case eof := <-eofCh:
			total = eof
		default:
			if total != 0 && processed != 0 && total == processed {
				done <- true
				return // differences between break and return
			}
		}
	}
}

func pooledPersistListedSymbol(db *sqlx.DB, listed <-chan *persist.TblListedSymbol, eofCh <-chan int, done chan<- bool) {
	var wg sync.WaitGroup
	var total, processed, skipped atomic.Int64
	for worker := 0; worker < 20; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case symbol := <-listed:
					tx, err := db.Beginx()
					if err != nil {
						log.Printf("pooledPersistListedSymbol error: %s, symbol %s.\n", err.Error(), symbol.Symbol)
						processed.Inc()
						skipped.Inc()
						continue
					}
					id, ra, err := persist.InsertListedSymbol(tx, symbol)
					if err != nil {
						log.Printf("pooledPersistListedSymbol error: %s, symbol %s.\n", err.Error(), symbol.Symbol)
						tx.Rollback()
						processed.Inc()
						skipped.Inc()
						continue
					}
					log.Printf("Persisted a new symbol %s, id %d, row affected: %d", symbol.Symbol, id, ra)
					tx.Commit()
					processed.Inc()
				case eof := <-eofCh:
					total.Store(int64(eof))
				default:
					if total.Load() != 0 && processed.Load() != 0 && total.Load() == processed.Load() {
						return // differences between break and return
					}
				}
			}
		}()
	}
	wg.Wait()
	log.Printf("pooledPersistListedSymbol done, total:%d, processed:%d, skipped:%d.\n", total.Load(), processed.Load(), skipped.Load())
	done <- true
}
