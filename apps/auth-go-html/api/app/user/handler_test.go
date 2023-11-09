package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"umsapi/app"
	"umsapi/app/user/models"
	"umsapi/internal/infra"

	"github.com/gofrs/uuid"
)

// WE CAN USE A MORE ADVANCED TEST LIBRARY LIKE GOMOCK IN CASE TESTS BECOME MORE COMPLEX

func TestSignUpHandler_Success(t *testing.T) {
	// Prepare the test data
	input := `{"email": "test@example.com", "password": "password123"}`
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(input))
	w := httptest.NewRecorder()

	// Prepare the mock dependencies
	app := &app.App{
		JSON:      infra.NewJSON(),
		Errors:    &infra.Errors{},
		Logger:    &infra.Logger{},
		Auth:      &app.AuthConfig{},
		AuthStore: nil,
		UserStore: &MockUserStore{},
	}

	// Call the handler
	SignUpHandler(w, req, app)

	// Check the response
	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.StatusCode)
	}
}

type MockUserStore struct{}

// CreateProfile is the only one needed.
func (*MockUserStore) CreateProfile(user *models.SignUpUser) (*models.Profile, error) {
	return &models.Profile{}, nil
}

// GetProfile implements store.UserStore.
func (*MockUserStore) GetProfile(id uuid.UUID) (*models.Profile, error) {
	panic("unimplemented")
}

// GetProfileByEmail implements store.UserStore.
func (*MockUserStore) GetProfileByEmail(email string) (*models.Profile, error) {
	panic("unimplemented")
}

// UpdateProfile implements store.UserStore.
func (*MockUserStore) UpdateProfile(profile *models.Profile) error {
	panic("unimplemented")
}
