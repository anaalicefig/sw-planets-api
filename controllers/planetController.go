package controllers

import (
	"encoding/json"
	"net/http"
	. "star-wars-api/models"
	. "star-wars-api/repositories"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var repository = PlanetRepository{}

func responseError(w http.ResponseWriter, code int, message string) {
	responseJson(w, code, map[string]string{"error": message})
}

func responseJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	name := v.Get("name")

	if name != "" {
		planet, err := repository.GetByName(name)

		if err != nil {
			responseError(w, http.StatusInternalServerError, "Invalid planet Name")
			return
		}

		responseJson(w, http.StatusOK, planet)
		return
	}

	planets, err := repository.GetAll()

	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJson(w, http.StatusOK, planets)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	planet, err := repository.GetById(params["id"])

	if err != nil {
		responseError(w, http.StatusInternalServerError, "Invalid planet ID")
		return
	}

	responseJson(w, http.StatusOK, planet)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var planet Planet

	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		responseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	planet.ID = bson.NewObjectId()

	if err := repository.Create(planet); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJson(w, http.StatusCreated, planet)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var planet Planet

	if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
		responseError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := repository.Update(params["id"], planet); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJson(w, http.StatusOK, map[string]string{"result": planet.Name + " atualizado com sucesso!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	if err := repository.Delete(params["id"]); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseJson(w, http.StatusOK, map[string]string{"result": "success"})
}
