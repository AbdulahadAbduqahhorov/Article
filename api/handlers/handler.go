package handlers

import "github.com/AbdulahadAbduqahhorov/gin/Article/storage"

type Handler struct {
	Stg storage.StorageI
}

func NewHandler(strg storage.StorageI) Handler {
	return Handler{
		Stg: strg,
	}
}
