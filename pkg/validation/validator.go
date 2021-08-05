package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type AlbumSchema struct {
	Title  string  `validate:"required"`
	Artist string  `validate:"required"`
	Price  float64 `validate:"required"`

}

func ValidateAlbumSchema(s *AlbumSchema)(err error){
	validate = validator.New()

	err = validate.Struct(s)
	if err != nil{
		return err
	}
	return nil
}