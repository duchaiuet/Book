package Database

import (
	"log"
	"os"
)

var (
	Db             *MongoDBStore
	ErrLog         *log.Logger
	InfoLog        *log.Logger
)

func init() {
	InfoLog = log.New(os.Stdout, "Infolog ", log.Ldate|log.Ltime|log.Llongfile)
	ErrLog = log.New(os.Stderr, "Errlog", log.Ldate|log.Ltime|log.Llongfile)
	loadParameters()
	Db = NewDataStore()

}
