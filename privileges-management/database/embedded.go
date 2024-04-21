package database

import (
	bolt "go.etcd.io/bbolt"
)

const (
	DbName           = "embedded.db"
	BucketNamePrefix = "shares_"
)

func OpenDatabase() (*bolt.DB, error) {
	db, err := bolt.Open(DbName, 0600, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateBucketIfNotExists(bucketName string) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}

		return nil
	})
}

func UpdateToBucket(key, value string) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	var bucket *bolt.Bucket
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err = tx.CreateBucketIfNotExists([]byte(BucketNamePrefix))
		if err != nil {
			return err
		}
		return bucket.Put([]byte(key), []byte(value))
	})
}

func GetBucketName(requestId string) string {
	return BucketNamePrefix + requestId
}

func RemoveBucket(bucketName string) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(bucketName))
	})
}
