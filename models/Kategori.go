package models

import (
	"gorm.io/gorm"
)

type Kategori struct {
	gorm.Model
	NamaKategori string
}

type KategoriRequest struct {
	NamaKategori string `json:"nama_kategori" form:"nama_kategori"`
}

type KategoriResponse struct {
	Id       uint `json:"id"`
	NamaKategori string `json:"nama_kategori"`
}

func RequestToModelKategori(dt KategoriRequest) Kategori {
	return Kategori{
		NamaKategori: dt.NamaKategori,
	}
}

func ModelToResponseKategori(dt Kategori) KategoriResponse {
	return KategoriResponse{
		Id:           dt.ID,
		NamaKategori: dt.NamaKategori,
	}
}

func ListModelToResponseKategori(dt []Kategori) []KategoriResponse {
	var responses = []KategoriResponse{}
	for _, v := range dt {
		responses = append(responses, ModelToResponseKategori(v))
	}
	return responses
}
