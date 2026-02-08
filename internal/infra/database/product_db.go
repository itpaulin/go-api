package database

import (
	"fmt"

	"github.com/itpaulin/go-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]*entity.Product, error) {

	var products []*entity.Product
	var err error
	offset := (page - 1) * limit
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Debug().Order("created_at " + sort).Offset(offset).Limit(limit).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}
	fmt.Println("essap orra ai", products)
	return products, err
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product

	err := p.DB.First(&product, "id = ?", id).Error

	return &product, err
}

func (p *Product) Update(product *entity.Product) error {
	// _, err := p.FindByID(id)

	// if err != nil {
	// 	return err
	// }
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	return p.DB.Delete(&entity.Product{}, "id = ?", id).Error
}
