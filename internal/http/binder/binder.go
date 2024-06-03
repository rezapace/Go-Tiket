package binder
// folder ini digunakan untuk mengcombine data yang diinputkan dengan data yang diinginkan
import (
	internalValidator "Ticketing/internal/http/validator"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// untuk override echo.Binder , karena untuk mapping apa saja yang perlu di binding
type Binder struct {
	defaultBinder *echo.DefaultBinder
	*internalValidator.FormValidator
}

//untuk mereturn struct binder diatas
func NewBinder(
	dbr *echo.DefaultBinder,
	vdr *internalValidator.FormValidator) *Binder {
	return &Binder{dbr, vdr}
}

// untuk melakukan binding
func (b *Binder) Bind(i interface{}, c echo.Context) error {
	if err := b.defaultBinder.Bind(i, c); err != nil {
		return err
	}

	if err := defaults.Set(i); err != nil {
		return err
	}

	// untuk melakukan validasi
	if err := b.Validate(i); err != nil {
		errs := err.(validator.ValidationErrors)
		return errs
	}

	return nil
}
