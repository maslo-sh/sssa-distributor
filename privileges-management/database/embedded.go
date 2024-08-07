package database

import (
	"errors"
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

func UpdateToBucket(bucketName, key, value string) error {
	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	defer db.Close()
	var bucket *bolt.Bucket
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err = tx.CreateBucketIfNotExists([]byte(bucketName))
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

func GetAllItemsFromBucket(bucketName string) ([]string, error) {
	db, err := OpenDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	items := make([]string, 0, 50)

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		b.ForEach(func(k, v []byte) error {
			items = append(items, string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetNumberOfItemsFromBucket(bucketName string) (int, error) {
	db, err := OpenDatabase()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var count int
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return errors.New("bucket not found")
		}
		c := bucket.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			count++
		}
		return nil
	})
	return count, nil
}
