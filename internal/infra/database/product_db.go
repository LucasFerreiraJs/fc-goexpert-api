package database

import (
	"fmt"

	"github.com/lucasferreirajs/fc-goexpert/apis/internal/entity"
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

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	var err error
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
		if err != nil {

		}
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindById(product.ID.String())
	if err != nil {
		return err
	}
	fmt.Printf("product recebido %v", product)
	print("\n")
	print("\n")
	err = p.DB.Save(product).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}
