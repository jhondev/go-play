package auth

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"umsapi/app"
	cmodels "umsapi/app/common/models"
	"umsapi/app/user/models"

	"github.com/gofrs/uuid"
	xoauth2 "golang.org/x/oauth2"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func GoogleAuthHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	authURL := app.Auth.GoogleOauth.AuthCodeURL("", xoauth2.AccessTypeOffline)
	http.Redirect(w, r, authURL, http.StatusSeeOther)
}

func GoogleAuthCallbackHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	ctx := context.Background()
	code := r.URL.Query().Get("code")
	token, err := app.Auth.GoogleOauth.Exchange(ctx, code)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	userInfo, err := getUserInfo(ctx, app.Auth.GoogleOauth.TokenSource(ctx, token), app)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	profile, err := app.UserStore.GetProfileByEmail(userInfo.Email)
	action := "profile"
	// SignUp user if it doesn't exist
	if err == cmodels.ErrRecordNotFound {
		id, _ := uuid.NewV7()
		profile, err = app.UserStore.CreateProfile(&models.SignUpUser{
			ID:        id,
			Name:      userInfo.Name,
			Email:     userInfo.Email,
			CreatedAt: time.Now().UTC(),
		})
		action = "profile/edit"
	}
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}
	_, appToken, _ := app.Jwt.Encode(map[string]interface{}{"user_id": profile.ID})
	http.Redirect(w, r, fmt.Sprintf("%s?action=%s&token=%s", app.Auth.ClientRedirectURL, action, appToken), http.StatusSeeOther)
}

func getUserInfo(ctx context.Context, tokenSource xoauth2.TokenSource, app *app.App) (*oauth2.Userinfo, error) {
	// Create a new oauth2 service
	service, err := oauth2.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		return nil, err
	}

	// Fetch the user info
	userinfo, err := service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		return nil, err
	}

	return userinfo, nil
}
