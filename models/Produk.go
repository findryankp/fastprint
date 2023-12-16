package models

import (
	"gorm.io/gorm"
)

type Produk struct {
	gorm.Model
	NamaProduk string
	Harga      int64
	KategoriId uint
	Kategori   Kategori `gorm:"foreignKey:KategoriId"`
	StatusId   uint
	Status     Status `gorm:"foreignKey:StatusId"`
}

type ProdukRequest struct {
	NamaProduk string `json:"nama_produk" form:"nama_produk"`
	Harga      int64  `json:"harga" form:"harga"`
	KategoriId uint   `json:"kategori_id" form:"kategori_id"`
	StatusId   uint   `json:"status_id" form:"status_id"`
}

type ProdukResponse struct {
	Id         uint   `json:"id"`
	NamaProduk string `json:"nama_produk"`
	Harga      int64  `json:"harga"`
	KategoriId uint   `json:"kategori_id"`
	Kategori   string `json:"kategori"`
	StatusId   uint   `json:"status_id"`
	Status     string `json:"status"`
}

func RequestToModelProduk(dt ProdukRequest) Produk {
	return Produk{
		NamaProduk: dt.NamaProduk,
		Harga:      dt.Harga,
		KategoriId: dt.KategoriId,
		StatusId:   dt.StatusId,
	}
}

func ModelToResponseProduk(dt Produk) ProdukResponse {
	result := ProdukResponse{
		Id:         dt.ID,
		NamaProduk: dt.NamaProduk,
		Harga:      dt.Harga,
		KategoriId: dt.KategoriId,
		Kategori:   dt.Kategori.NamaKategori,
		StatusId:   dt.StatusId,
		Status:     dt.Status.NamaStatus,
	}

	return result
}

func ListModelToResponseProduk(dt []Produk) []ProdukResponse {
	var responses = []ProdukResponse{}
	for _, v := range dt {
		responses = append(responses, ModelToResponseProduk(v))
	}
	return responses
}
