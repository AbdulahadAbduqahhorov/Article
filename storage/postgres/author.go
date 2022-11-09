package postgres

import (
	"errors"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage"
	"github.com/jmoiron/sqlx"
)

type authorRepo struct {
	db *sqlx.DB
}

func NewAuthorRepo(db *sqlx.DB) storage.AuthorRepoI {
	return authorRepo{
		db: db,
	}
}
func (stg authorRepo) CreateAuthor(id string, author models.CreateAuthorModel) error {

	_, err := stg.db.Exec(`INSERT INTO 
		author (
			id,
			firstname,
			lastname
			) 
		VALUES (
			$1, 
			$2,
			$3
			)`,
		id,
		author.FirstName,
		author.LastName,
	)
	if err != nil {
		return err
	}
	return nil

}

func (stg authorRepo) GetAuthor(limit, offset int, search string) ([]models.Author, error) {
	var res []models.Author

	rows, err := stg.db.Queryx(`SELECT 
		id,
		firstname,
		lastname,
		created_at,
		updated_at,
		deleted_at 
		FROM author
		WHERE ((firstname ILIKE '%' || $1 || '%') or (lastname ILIKE '%' || $1 || '%') ) AND deleted_at IS NULL
		LIMIT $2
		OFFSET $3
	`,
		search,
		limit,
		offset,
	)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var author models.Author
		err := rows.Scan(
			&author.Id,
			&author.FirstName,
			&author.LastName,
			&author.CreatedAt,
			&author.UpdatedAt,
			&author.DeletedAt,
		)
		if err != nil {
			return res, err
		}

		res = append(res, author)

	}

	return res, err

}

func (stg authorRepo) GetAuthorById(id string) (models.Author, error) {
	var author models.Author

	err := stg.db.QueryRow(`
	SELECT 
		id,
		firstname,
		lastname,
		created_at,
		updated_at,
		deleted_at
	FROM author  
	WHERE id=$1 AND deleted_at is NULL`, id).Scan(
		&author.Id,
		&author.FirstName,
		&author.LastName,
		&author.CreatedAt,
		&author.UpdatedAt,
		&author.DeletedAt,
	)
	if err != nil {
		return author, err
	}
	return author, nil
}

func (stg authorRepo) UpdateAuthor(author models.UpdateAuthorModel) error {
	res, err := stg.db.NamedExec(`
	UPDATE  author SET 
		firstname=:f, 
		lastname=:l,
		updated_at=now() 
		WHERE id=:i AND deleted_at IS NULL `, map[string]interface{}{
		"f": author.FirstName,
		"l": author.LastName,
		"i": author.Id,
	})
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n > 0 {
		return nil
	}
	return errors.New("author not found")
}

func (stg authorRepo) DeleteAuthor(id string) error {

	res, err := stg.db.Exec(`UPDATE author SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL`, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return errors.New("author not found")
	}
	return nil
}
