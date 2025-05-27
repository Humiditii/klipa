package db

import (
	"strconv"

	"github.com/atotto/clipboard"
)

type KlipaModel struct {
	ID int64
	Key string `binding:"required"`
	Value string `binding:"required"`
}

func (k *KlipaModel) Save() {
	query := `INSERT INTO klipa (key, value) VALUES (?, ?)`

	stmt, err := Db.Prepare(query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(k.Key, k.Value)

	if err != nil {
		panic(err)
	}
}

func DeleteOne(k string){
	query := `DELETE FROM klipa WHERE key = ?`

	stmt, err := Db.Prepare(query)

	if err!=nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(k)

	if err != nil {
		panic(err)
	}
}

func FlushMem(){
	query := `DELETE FROM klipa`

	stmt, err := Db.Prepare(query)

	if err!=nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

func GetValueByKey( k string) string {
	query := `SELECT value FROM klipa WHERE key = ?`

	var value string

	result := Db.QueryRow(query, k)

	result.Scan(&value)

	return value

	
}

func GetKeys() *[]string {

	query := `SELECT key FROM klipa`

	rows, err := Db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var keys []string

	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			panic(err)
		}
		keys = append(keys, key)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return &keys
}

func GetLastIdFromDb() string {
	query := `SELECT id FROM klipa ORDER BY id DESC LIMIT 1`
	row := Db.QueryRow(query)

	var id int
	_ = row.Scan(&id)

	return "klipa_"+strconv.Itoa(id + 1)
}

func CopyFromKlipaToMem(k string){
	clipboard.WriteAll(GetValueByKey(k))
}