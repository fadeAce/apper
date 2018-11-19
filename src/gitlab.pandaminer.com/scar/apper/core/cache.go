package core

type cache struct {
}

func Caching(key, txnID, typ string, value interface{}) {
	// txID - fragmentID - data
	// caching for consumer to persis as a unit

	// persis to postgres db
}


func CachingFailure(s string, s2 string, s3 string) {

}