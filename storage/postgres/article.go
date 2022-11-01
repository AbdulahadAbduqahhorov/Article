package postgres

import (
	"errors"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
)

func (stg Postgres) CreateArticle(id string, article models.CreateArticleModel) error {

	_, err := stg.GetAuthorById(article.AuthorId)
	if err != nil {
		return err
	}
	_, err = stg.db.Exec(`INSERT INTO 
		article (
			id,
			title,
			body,
			author_id) 
		VALUES (
			$1, 
			$2,
			$3,
			$4)`,
		id,
		article.Title,
		article.Body,
		article.AuthorId,
	)
	if err != nil {
		return err
	}
	return nil

}

func (stg Postgres) GetArticle(limit, offset int, search string) ([]models.Article, error) {
	var res []models.Article

	rows, err := stg.db.Queryx(`SELECT 
		id,
		title,
		body,
		author_id,
		created_at,
		updated_at,
		deleted_at 
		FROM article
		WHERE title ILIKE '%' || $1 || '%' AND deleted_at IS NULL
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
		var article models.Article
		err := rows.Scan(
			&article.Id,
			&article.Title,
			&article.Body,
			&article.AuthorId,
			&article.CreatedAt,
			&article.UpdatedAt,
			&article.DeletedAt,
		)
		if err != nil {
			return res, err
		}

		res = append(res, article)

	}

	return res, err

}

func (stg Postgres) GetArticleById(id string) (models.GetArticleByIdModel, error) {
	var article models.GetArticleByIdModel

	err := stg.db.QueryRow(`
	SELECT 
		ar.id,
		ar.title,
		ar.body,
		ar.created_at,
		ar.updated_at,
		ar.deleted_at,
		au.id,
		au.firstname,
		au.lastname,
		au.created_at,
		au.updated_at,
		au.deleted_at
	FROM article ar 
	JOIN author au 
	ON article.author_id=author.id 
	WHERE ar.id=$1`, id).Scan(
		&article.Id,
		&article.Title,
		&article.Body,
		&article.CreatedAt,
		&article.UpdatedAt,
		&article.DeletedAt,
		&article.Author.Id,
		&article.Author.FirstName,
		&article.Author.LastName,
		&article.Author.CreatedAt,
		&article.Author.UpdatedAt,
		&article.Author.DeletedAt,
	)
	if err != nil {
		return article, err
	}
	return article, nil
}

func (stg Postgres) UpdateArticle(article models.UpdateArticleModel) error {

	res, err := stg.db.NamedExec(`
	UPDATE  article SET 
		title=:t, 
		body=:b,
		updated_at=now() 
		WHERE id=:i AND deleted_at IS NULL )`, map[string]interface{}{
		"t": article.Title,
		"b": article.Body,
		"i": article.Id,
	})
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return errors.New("article not found")
	}
	return nil
}

func (stg Postgres) DeleteArticle(id string) error {
	res, err := stg.db.NamedExec(`UPDATE article SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL)`, id)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return errors.New("article not found")
	}
	return nil

}
