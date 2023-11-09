package user

import (
	"net/http"
	"umsweb/app"
	"umsweb/internal/client"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	app.Tmplr.Render(w, "signup", getSignUpData(nil, app))
}

func SignUpFormHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	profile := app.Client.SignUp(email, password)
	if profile.Errors != nil {
		app.Tmplr.Render(w, "signup", getSignUpData(profile.Errors, app))
		return
	}
	_, token, err := app.Jwt.Encode(map[string]interface{}{"user_id": profile.ID})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	})

	http.Redirect(w, r, "/profile/edit", http.StatusFound)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	token := getTokenString(r)
	profile := app.Client.GetProfile(token)
	app.Tmplr.Render(w, "profile", profile)
}

func ProfileEditHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	token := getTokenString(r)
	profile := app.Client.GetProfile(token)
	app.Tmplr.Render(w, "profile-edit", profile)
}

func ProfileUpdateHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	profile := &client.PatchProfile{Email: r.FormValue("email"), External: r.FormValue("external")}
	name := r.FormValue("name")
	if name != "" {
		profile.Name = name
	}
	tel := r.FormValue("telephone")
	if tel != "" {
		profile.Telephone = tel
	}

	errs := app.Client.UpdateProfile(profile, getTokenString(r))
	if errs != nil {
		profile.Errors = errs
		app.Tmplr.Render(w, "profile-edit", profile)
		return
	}

	http.Redirect(w, r, "/profile", http.StatusFound)
}

func getTokenString(r *http.Request) string {
	c, _ := r.Cookie("jwt")
	return c.Value
}

func getSignUpData(errs map[string]string, app *app.App) any {
	return struct {
		AuthExternalURL string
		Errors          map[string]string
	}{
		AuthExternalURL: app.AuthExternalURL,
		Errors:          errs,
	}
}
