package controller

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/stretchr/testify/assert"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) ValidateStruct(obj interface{}) error {
	if cv.validator == nil {
		return fmt.Errorf("validator is not initialized")
	}
	return cv.validator.Struct(obj)
}

func (cv *customValidator) Engine() interface{} {
	return cv.validator
}

// Wrapper function for RegisterDefaultTranslations
var registerDefaultTranslations = enTranslations.RegisterDefaultTranslations

func TestInitTrans_RegisterTranslationsFailure(t *testing.T) {
	// Create a custom validator that wraps *validator.Validate
	v := &customValidator{validator: validator.New()}
	binding.Validator = v

	// Monkey patch the wrapper function to simulate failure
	originalRegister := registerDefaultTranslations
	registerDefaultTranslations = func(v *validator.Validate, trans ut.Translator) error {
		return fmt.Errorf("mocked registration error")
	}
	defer func() { registerDefaultTranslations = originalRegister }()

	_, err := InitTrans("en")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to register translator for en")
}

func TestInitTrans_ValidLocales(t *testing.T) {
	tests := []struct {
		locale        string
		expectedError bool
	}{
		{"en", false},
		{"zh", false},
		// {"fr", false}, // fallback to en
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Locale_%s", tt.locale), func(t *testing.T) {
			trans, err := InitTrans(tt.locale)
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, trans)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, trans)
			}
		})
	}
}
