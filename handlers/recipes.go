package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recipeServer/db"
	"strings"
)

// RecipeHandler handles requests for recipe resource
type RecipeHandler struct {
	Prefix string
}

func (h RecipeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.ReplaceAll(r.URL.Path, h.Prefix, "")
	realPath := strings.Replace(path, "/", "", 1)
	fmt.Println(r.Method, path, r.URL.Path, h.Prefix)
	switch r.Method {
	case http.MethodGet:
		switch realPath {
		case "":
			getRecipes(w, r)
		default:
			getRecipe(w, r, realPath)
		}
	case http.MethodPost:
		postRecipe(w, r)
	case http.MethodPatch:
		updateRecipe(w, r, false)
	case http.MethodPut:
		updateRecipe(w, r, true)
	case http.MethodDelete:
		deleteRecipe(w, r)
	}

}

func getRecipe(w http.ResponseWriter, r *http.Request, uid string) {
	fmt.Printf("Getting a recipe with id: %v\n", uid)

	recipeID := ParseID(uid, w)
	recipe, err := db.GetOne(recipeID)
	if err != nil {
		Throw("No recipe was found", http.StatusNotFound, err, w)
		return
	}
	SendJSON(recipe, w)
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all recipes")

	recipes, err := db.GetAll()
	if err != nil {
		Throw("Failed to parse body", http.StatusInternalServerError, err, w)
		return
	}

	SendJSON(recipes, w)
}

func postRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating recipe")

	decoder := json.NewDecoder(r.Body)
	data := &db.RecipeItem{}
	err := decoder.Decode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to parse body"))
		fmt.Println(err)
		return
	}

	resp, err := db.Create(&db.RecipeItem{
		Content: data.Content,
		Title:   data.Title,
	})
	if err != nil {
		Throw("Failed to create recipe", http.StatusInternalServerError, err, w)
		return
	}
	SendJSON(resp, w)
}

func updateRecipe(w http.ResponseWriter, r *http.Request, rewrite bool) {
	fragments := strings.Split(r.URL.Path, "/")
	uid := fragments[len(fragments)-1]
	if len(uid) == 0 {
		Throw("No id param provided", http.StatusBadRequest, nil, w)
		return
	}
	recipeID := ParseID(uid, w)

	decoder := json.NewDecoder(r.Body)
	data := &db.RecipeItem{}
	err := decoder.Decode(data)
	if err != nil {
		Throw("Failed to parse body", http.StatusInternalServerError, err, w)
		return
	}

	respone, updateErr := db.Update(recipeID, data, rewrite)
	if updateErr != nil {
		Throw("Failed to create recipe", http.StatusInternalServerError, updateErr, w)
		return
	}
	SendJSON(respone, w)
}

func deleteRecipe(w http.ResponseWriter, r *http.Request) {
	fragments := strings.Split(r.URL.Path, "/")
	uid := fragments[len(fragments)-1]
	if len(uid) == 0 {
		Throw("No id param provided", http.StatusBadRequest, nil, w)
		return
	}
	fmt.Printf("Deleting a recipe with id: %v\n", uid)

	recipeID := ParseID(uid, w)
	db.Delete(recipeID)
	w.Write([]byte("Recipe deleted"))
}
