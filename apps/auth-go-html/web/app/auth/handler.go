package auth

import (
	"net/http"
	"umsweb/app"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	app.Tmplr.Render(w, "login", getLoginData(nil, app))
}

func LoginSubmitHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	id, errs := app.Client.Authenticate(email, password)
	if errs != nil {
		app.Tmplr.Render(w, "login", getLoginData(errs, app))
		return
	}
	_, token, err := app.Jwt.Encode(map[string]interface{}{"user_id": id})
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

	http.Redirect(w, r, "/profile", http.StatusFound)
}

func APIAuthCallbackHandler(w http.ResponseWriter, r *http.Request, app *app.App) {
	action := r.URL.Query().Get("action")
	if action == "" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	})

	http.Redirect(w, r, "/"+action, http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1, // MaxAge<0 means delete cookie now
	})

	http.Redirect(w, r, "/login", http.StatusFound)
}

func getLoginData(errs map[string]string, app *app.App) any {
	return struct {
		AuthExternalURL string
		Errors          map[string]string
	}{
		AuthExternalURL: app.AuthExternalURL,
		Errors:          errs,
	}
}
