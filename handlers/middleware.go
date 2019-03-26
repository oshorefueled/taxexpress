package handlers

import (
	"database/sql"
	"fmt"
	"github.com/oshorefueled/taxexpress/models"
	"net/http"
)

func needTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		h := handlerHelper{}
		err := checkUserToken(authorization)
		if err != nil {
			h.RespondWithError(w, 400, err.Error())
			return
		}
		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

/*
	Function to be replaced with JWT Authentication
 */
func checkUserToken (authorization string) (err error) {
	admin := models.Admin{Token: authorization}
	err = admin.GetAdminByToken()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Admin token does not exist")
			return err
		} else {
			fmt.Println(err.Error())
		}
	}
	return err
}