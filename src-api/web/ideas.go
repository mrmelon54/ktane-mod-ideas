package web

import (
	"encoding/json"
	"fmt"
	"github.com/MrMelon54/ktane-mod-ideas/src-api/structure"
	"github.com/MrMelon54/ktane-mod-ideas/src-api/utils"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (w *Web) homeIdeas(rw http.ResponseWriter, _ *http.Request, _ *utils.State, _ bool, _ structure.User) {
	// Find ideas
	var ideas []structure.Idea
	err := w.engine.Limit(10).Asc("id").Find(&ideas)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Failed to read from database", http.StatusInternalServerError)
		return
	}
	if ideas == nil {
		ideas = []structure.Idea{}
	}
	w.addAuthorToIdeaArray(ideas)

	// Encode output
	encoder := json.NewEncoder(rw)
	err = encoder.Encode(ideas)
	if err != nil {
		http.Error(rw, "Failed to encode data", http.StatusInternalServerError)
		return
	}
}

func (w *Web) searchIdeas(rw http.ResponseWriter, req *http.Request, _ *utils.State, _ bool, _ structure.User) {
	// Parse input
	q := req.URL.Query()
	searchQuery := q.Get("q")

	// Find ideas
	var ideas []structure.Idea
	b, err := w.engine.Where("name = ?", fmt.Sprintf("%%%s%%", searchQuery)).Asc("name").Get(&ideas)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Failed to read from database", http.StatusInternalServerError)
		return
	}
	if !b {
		http.Error(rw, "404 Not Found", http.StatusNotFound)
		return
	}
	if ideas == nil {
		ideas = []structure.Idea{}
	}
	w.addAuthorToIdeaArray(ideas)

	// Encode output
	encoder := json.NewEncoder(rw)
	err = encoder.Encode(ideas)
	if err != nil {
		http.Error(rw, "Failed to encode data", http.StatusInternalServerError)
		return
	}
}

func (w *Web) getIdea(rw http.ResponseWriter, req *http.Request, _ *utils.State, _ bool, _ structure.User) {
	// Parse input
	q := req.URL.Query()
	ideaId, err := strconv.Atoi(q.Get("idea"))
	if err != nil {
		http.Error(rw, "Invalid idea parameter", http.StatusBadRequest)
		return
	}

	// Get idea
	var idea structure.Idea
	foundIdea, err := w.engine.Where("id = ?", ideaId).Get(&idea)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Failed to read from database", http.StatusInternalServerError)
		return
	}
	if !foundIdea {
		http.Error(rw, "404 Not Found", http.StatusNotFound)
		return
	}
	w.addAuthorToIdea(&idea)

	// Encode output
	encoder := json.NewEncoder(rw)
	err = encoder.Encode(idea)
	if err != nil {
		http.Error(rw, "Failed to encode data", http.StatusInternalServerError)
		return
	}
}

func (w *Web) makeIdea(rw http.ResponseWriter, req *http.Request, _ *utils.State, loggedIn bool, user structure.User) {
	// Check permissions
	if !loggedIn {
		http.Error(rw, "401 Forbidden: User must be logged in", http.StatusForbidden)
		return
	}
	if user.Banned {
		http.Error(rw, "401 Forbidden: User is unable to make new ideas", http.StatusForbidden)
		return
	}

	var idea structure.Idea
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&idea)
	if err != nil {
		http.Error(rw, "400 Bad Request: Unable to decode body", http.StatusBadRequest)
		return
	}

	idea.Id = 0
	idea.Owner = user.Id
	idea.CreatedAt = time.Time{}
	idea.UpdatedAt = time.Time{}

	switch idea.State {
	case structure.IdeaStateRed, structure.IdeaStateYellow, structure.IdeaStateGreen:
	default:
		idea.State = structure.IdeaStateUnknown
	}

	_, err = w.engine.Insert(idea)
	if err != nil {
		http.Error(rw, "507 Insufficient Storage", http.StatusInsufficientStorage)
		return
	}
}

func (w *Web) modifyIdea(rw http.ResponseWriter, req *http.Request, _ *utils.State, loggedIn bool, user structure.User) {
	// Check permissions
	if !loggedIn {
		http.Error(rw, "401 Forbidden: User must be logged in", http.StatusForbidden)
		return
	}
	if user.Banned {
		http.Error(rw, "401 Forbidden: User is unable to make new ideas", http.StatusForbidden)
		return
	}

	var idea structure.Idea
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&idea)
	if err != nil {
		http.Error(rw, "400 Bad Request: Unable to decode body", http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodPatch:
		_, err := w.engine.Where("idea = ?", idea.Id).Update(idea)
		if err != nil {
			http.Error(rw, "507 Insufficient Storage", http.StatusInsufficientStorage)
			return
		}
	case http.MethodDelete:
		_, err := w.engine.Where("idea = ?", idea.Id).Delete(&structure.Idea{})
		if err != nil {
			http.Error(rw, "507 Insufficient Storage", http.StatusInsufficientStorage)
			return
		}
	}
}

func (w *Web) addAuthorToIdea(idea *structure.Idea) {
	var user structure.User
	b, err := w.engine.Where("id = ?", idea.Owner).Get(&user)
	if err != nil || !b {
		idea.Author = "<Unknown>"
		return
	}
	idea.Author = user.DiscordName
}

func (w *Web) addAuthorToIdeaArray(idea []structure.Idea) {
	authors := make([]int64, len(idea))
	for i := range idea {
		authors[i] = idea[i].Owner
	}
	var users []structure.User
	err := w.engine.In("idea", authors).Find(&users)
	if err != nil {
		for i := range idea {
			idea[i].Author = "<Unknown>"
		}
		return
	}

outer:
	for i := range idea {
		o := idea[i].Owner
		for _, u := range users {
			if u.Id == o {
				idea[i].Author = u.DiscordName
				continue outer
			}
		}
	}
}
