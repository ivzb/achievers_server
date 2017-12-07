package controller

import (
	"net/http"

	"app/model/auth"
	"app/model/token"
	"app/shared/crypto"
	"app/shared/form"
	"app/shared/response"
	"app/shared/router"
	mtoken "app/shared/token"

	"github.com/justinas/alice"
)

const (
	Unauthorized = "unauthorized"
	Authorized   = "authorized"
)

// Routes
func init() {
	router.Post("/auth", alice.
		New( /*mauth.Handler*/ ).
		ThenFunc(AuthPOST))
}

// *****************************************************************************
// Auth
// *****************************************************************************
func AuthPOST(w http.ResponseWriter, r *http.Request) {
	auth := auth.New()
	err, errMsg := form.Validate(r, auth)

	if err == form.ErrRequiredMissing || err == form.ErrWrongContentType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	}

	if err == form.ErrBadStruct || err == form.ErrNotStruct {
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	err, errMsg = form.StructCopy(r, auth)
	if err == form.ErrWrongType {
		response.SendError(w, http.StatusBadRequest, errMsg)
		return
	}

	if err == form.ErrNotSupported || err == form.ErrNotStruct {
		response.SendError(w, http.StatusInternalServerError, FriendlyError)
		return
	}

	// Get an item
	user, err := auth.Auth()

	if err != nil {
		response.Send(w, http.StatusUnauthorized, Unauthorized, 0, nil)
		return
	}

	priv := mtoken.ReadConfig().Priv

	t, err := crypto.Encrypt([]byte(user.ID), priv.PublicKey)

	if err != nil {
		response.Send(w, http.StatusInternalServerError, FriendlyError, 0, nil)
		return
	}

	tr := &token.Entity{AuthToken: t}

	response.Send(w, http.StatusOK, Authorized, 1, tr)
}
