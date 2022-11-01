package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS author (
	id CHAR(36) PRIMARY KEY,
	firstname VARCHAR(255) NOT NULL,
	lastname VARCHAR(255) NOT NULL ,
	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP	

);

CREATE TABLE IF NOT EXISTS article (
   	id CHAR(36) PRIMARY KEY,
   	title VARCHAR(255) UNIQUE NOT NULL,
   	body TEXT NOT NULL,
   	author_id CHAR(36),
   	created_at TIMESTAMP DEFAULT now() NOT NULL,
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
   	CONSTRAINT fk_author
   	FOREIGN KEY(author_id) 
   	REFERENCES author(id)
);

`

func InitDb(config string) (*Postgres, error) {

	tempDb, err := sqlx.Connect("postgres", config)
	if err != nil {
		return nil, err
	}
	tempDb.MustExec(schema)

	tx := tempDb.MustBegin()
	tx.MustExec("INSERT INTO author (id,firstname, lastname) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING", "de9fc112-7511-4abd-9222-7019b65d1108", "John", "Smith")
	tx.MustExec("INSERT INTO author (id,firstname, lastname) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING", "9ae20328-995c-401b-90a0-9acadf63c6ec", "Abdulahad", "Abduqahhorov")

	tx.MustExec("INSERT INTO article (id,title, body,author_id) VALUES ($1, $2, $3,$4) ON CONFLICT DO NOTHING", "ead02d29-5bf9-4b9e-91c0-c6e1ab648937", "title 1", "body 1", "9ae20328-995c-401b-90a0-9acadf63c6ec")
	tx.MustExec("INSERT INTO article (id,title, body,author_id) VALUES ($1, $2, $3,$4) ON CONFLICT DO NOTHING", "0bd2034c-5283-4e59-b904-02ff9fa8ed48", "title 2", "body 2", "9ae20328-995c-401b-90a0-9acadf63c6ec")
	tx.MustExec("INSERT INTO article (id,title, body,author_id) VALUES ($1, $2, $3,$4) ON CONFLICT DO NOTHING", "cb7cfd12-6501-4e8c-8297-cb442512b9ba", "title 3", "body 3", "de9fc112-7511-4abd-9222-7019b65d1108")
	tx.NamedExec("INSERT INTO article (id,title, body,author_id) VALUES (:id, :t, :b,:a) ON CONFLICT DO NOTHING", map[string]interface{}{
		"id": "0579a273-a581-4806-8e0b-4b9188c852ec",
		"t":  "title 4",
		"b":  "body 4",
		"a":  "de9fc112-7511-4abd-9222-7019b65d1108",
	})

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &Postgres{db: tempDb}, nil

}
