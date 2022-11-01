package postgres

import (
	"errors"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
)

func (stg Postgres) CreateAuthor(id string, author models.CreateAuthorModel) error {

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

func (stg Postgres) GetAuthor(limit, offset int, search string) ([]models.Author, error) {
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

func (stg Postgres) GetAuthorById(id string) (models.Author, error) {
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
	WHERE id=$1`, id).Scan(
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

func (stg Postgres) UpdateAuthor(author models.UpdateAuthorModel) error {
	res, err := stg.db.NamedExec(`
	UPDATE  author SET 
		firstname=:f, 
		lastname=:l,
		updated_at=now() 
		WHERE id=:i AND deleted_at IS NULL )`, map[string]interface{}{
		"f": author.FirstName,
		"l": author.LastName,
		"i": author.Id,
	})
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return errors.New("author not found")
	}
	return nil
}

func (stg Postgres) DeleteAuthor(id string) error {

	res, err := stg.db.NamedExec(`UPDATE author SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL)`, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return errors.New("author not found")
	}
	return nil
}
