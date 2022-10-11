package handlers

import "github.com/AbdulahadAbduqahhorov/gin/Article/storage/inmemory"

type Handler struct{
	Im inmemory.InMemory
}