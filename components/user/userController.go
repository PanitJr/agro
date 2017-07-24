package user

import (
	"encoding/json"
	"net/http"

	"agro/server/http/responseHelper"

	"github.com/gorilla/mux"
)

//APIUser userService for implement this userController.
type APIUser struct {
	userService *userServiceImpl
}

//List function will return a list of User struct as a json.
func (apiUser *APIUser) List(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := apiUser.userService.List()
	if err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrNotImplemented, w)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	responseHelper.HTTPJSONResponse(data, nil, w)
	return
}

//Create function use to create user into system then return it back as a json.
func (apiUser *APIUser) Create(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrBadRequset, w)
		return
	}
	if err := user.ValidateUser(); err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrBadRequset, w)
		return
	}
	createdUser, err := apiUser.userService.Create(user)
	if err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrNotImplemented, w)
		return
	}
	responseHelper.HTTPJSONResponse(createdUser, nil, w)
	return

}

//Update function use to update user in system then return it back as a json.
func (apiUser *APIUser) Update(w http.ResponseWriter, r *http.Request) {

	user := &User{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrBadRequset, w)
		return
	}
	if err := user.ValidateUser(); err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrBadRequset, w)
		return
	}
	updatedUser, err := apiUser.userService.Update(user)
	if err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrNotImplemented, w)
		return
	}
	responseHelper.HTTPJSONResponse(updatedUser, nil, w)
	return

}

//Delete function use to delete user in system.
func (apiUser *APIUser) Delete(w http.ResponseWriter, r *http.Request) {
	user := &User{}

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrBadRequset, w)
		return
	}

	updatedUser, err := apiUser.userService.Delete(user)
	if err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrNotImplemented, w)
		return
	}
	responseHelper.HTTPJSONResponse(updatedUser, nil, w)
	return
}

//Get function use to query user in system then retrun it as a json.
func (apiUser *APIUser) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user := &User{}

	foundUser, err := apiUser.userService.Get(user, id)
	if err != nil {
		responseHelper.HTTPJSONResponse(err.Error(), responseHelper.ErrNotFound, w)
		return
	}
	responseHelper.HTTPJSONResponse(foundUser, nil, w)
	return
}
