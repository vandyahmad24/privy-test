package database

import (
	"database/sql"
	"errors"
	"vandyahmad24/privy/app/db/model"
	"vandyahmad24/privy/app/tracing"

	"github.com/opentracing/opentracing-go"
)

type Cake struct {
	db *sql.DB
}

func NewCake(db *sql.DB) *Cake {
	return &Cake{
		db: db,
	}
}

func (cl *Cake) InsertCake(span opentracing.Span, input *model.Cake) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "InsertCake")
	defer sp.Finish()
	tracing.LogRequest(sp, input)

	result, err := cl.db.Exec("INSERT INTO cake(title, description, rating, image, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?)", input.Title, input.Description, input.Rating, input.Image, input.CreatedAt, input.UpdatedAt)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	id, err := result.LastInsertId()
	input.Id = int(id)

	tracing.LogResponse(sp, result)

	return input, nil
}

func (cl *Cake) GetAll(span opentracing.Span) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetAll")
	defer sp.Finish()
	var (
		cake  model.Cake
		cakes []model.Cake
	)

	rows, err := cl.db.Query("select * from cake order by rating desc, title")
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			tracing.LogError(sp, err)
			return nil, err
		}
		cakes = append(cakes, cake)

	}

	tracing.LogResponse(sp, cakes)

	return cakes, nil
}

func (cl *Cake) Get(span opentracing.Span, id int) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "Get")
	defer sp.Finish()
	var (
		cake model.Cake
	)
	row := cl.db.QueryRow("select * from cake where id = ?", id)
	err := row.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("Record Not Found")
	} else if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	tracing.LogResponse(sp, cake)
	return cake, nil
}

func (cl *Cake) Delete(span opentracing.Span, id int) error {
	sp := tracing.CreateSubChildSpan(span, "Delete")
	defer sp.Finish()

	_, err := cl.db.Exec("DELETE FROM cake WHERE id=?", id)
	if err == sql.ErrNoRows {
		tracing.LogError(sp, err)
		return errors.New("Failed Delete")
	}
	return nil
}

func (cl *Cake) Update(span opentracing.Span, id int, input *model.Cake) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "Update")
	defer sp.Finish()

	_, err := cl.db.Exec("UPDATE cake SET title= ?, description= ? , rating= ?, image= ?, updated_at= ?  WHERE id=?", input.Title, input.Description, input.Rating, input.Image, input.UpdatedAt, id)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, errors.New("Failed Update")
	}
	res, _ := cl.Get(sp, id)

	tracing.LogResponse(sp, res)
	return res, nil
}
