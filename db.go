package c4

import (
	"os"

	"github.com/boltdb/bolt"
)

type DB struct {
	BoldDb *bolt.DB
}

func OpenDB(path string, mode os.FileMode, options *bolt.Options) (*DB, error) {
	bdb, err := bolt.Open(path, mode, options)
	if err != nil {
		return nil, err
	}
	err = bdb.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("assets")); err != nil {
			return err
		}
		if _, err := tx.CreateBucketIfNotExists([]byte("attributes")); err != nil {
			return err
		}
		return nil
	})

	return &DB{bdb}, err
}

func (db *DB) getIDFromBucket(key []byte, bucket string) *ID {
	var id *ID
	db.BoldDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		idbytes := b.Get(key)
		if len(idbytes) == 64 {
			id = BytesToID(idbytes)
		}
		return nil
	})
	return id
}

func (db *DB) setIDForBucket(key []byte, id *ID, bucket string) error {
	return db.BoldDb.Batch(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(key, id.RawBytes())
		return nil
	})
}

func (db *DB) SetAssetID(key []byte, id *ID) error {
	return db.setIDForBucket(key, id, "assets")
}

func (db *DB) SetAttributesID(key []byte, id *ID) error {
	return db.setIDForBucket(key, id, "attributes")
}

func (db *DB) GetAssetID(key []byte) *ID {
	return db.getIDFromBucket(key, "assets")
}

func (db *DB) GetAttributesID(key []byte) *ID {
	return db.getIDFromBucket(key, "attributes")
}

func (db *DB) ForEach(f func(key []byte, asset *ID, attributes *ID) error) {
	db.BoldDb.View(func(tx *bolt.Tx) error {
		assets_bkt := tx.Bucket([]byte([]byte("assets")))
		attributes_bkt := tx.Bucket([]byte([]byte("attributes")))
		return assets_bkt.ForEach(func(key []byte, asset_value []byte) error {
			if len(asset_value) != 64 {
				return nil
			}
			attribute_value := attributes_bkt.Get(key)
			if len(attribute_value) != 64 {
				return nil
			}
			asset_id := BytesToID(asset_value)
			attribute_id := BytesToID(attribute_value)
			return f(key, asset_id, attribute_id)
		})
	})
}

func (db *DB) Close() error {
	return db.BoldDb.Close()
}
