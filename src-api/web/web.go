package web

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/MrMelon54/ktane-mod-ideas/src-api/structure"
	"github.com/MrMelon54/ktane-mod-ideas/src-api/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"xorm.io/xorm"
)

const (
	LoginFrameStart = "<!DOCTYPE html><html><head><script>window.opener.postMessage({user:"
	LoginFrameEnd   = "},\"%s\");window.close();</script></head></html>"
	CheckFrameStart = "<!DOCTYPE html><html><head><script>window.onload=function(){window.parent.postMessage({user:"
	CheckFrameEnd   = "},\"%s\");}</script></head></html>"
)

type Web struct {
	oauthClient  *oauth2.Config
	stateManager *utils.StateManager
	pageTitle    string
	engine       *xorm.Engine
	ownerSub     string
	originUrl    string
	resourceUrl  string
}

func New(engine *xorm.Engine) *Web {
	sessionStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	return &Web{
		oauthClient: &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			Scopes:       []string{"identify"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://discord.com/oauth2/authorize",
				TokenURL: "https://discord.com/api/oauth2/token",
			},
			RedirectURL: os.Getenv("REDIRECT_URL"),
		},
		stateManager: utils.NewStateManager(sessionStore),
		pageTitle:    os.Getenv("TITLE"),
		engine:       engine,
		ownerSub:     os.Getenv("OWNER"),
		originUrl:    os.Getenv("ORIGIN_URL"),
		resourceUrl:  discordgo.EndpointUser("@me"),
	}
}

func (w *Web) SetupWeb() *http.Server {
	gob.Register(new(utils.WebServiceKeyType))

	router := mux.NewRouter()
	router.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(rw, "KTaNE Mod Ideas API endpoint")
	})
	router.HandleFunc("/login", w.stateManager.SessionWrapper(w.loginPage))
	router.HandleFunc("/logout", w.stateManager.ClearWrapper(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintln(rw, "Logged out")
	}))
	router.HandleFunc("/check", w.stateManager.SessionWrapper(w.checkPage))
	router.HandleFunc("/admin", w.stateManager.SessionWrapper(w.loginWrapper(func(rw http.ResponseWriter, req *http.Request, state *utils.State, loggedIn bool, user structure.User) {
		if user.Admin {
			_, _ = rw.Write([]byte("is admin\n"))
		} else {
			_, _ = rw.Write([]byte("not admin\n"))
		}
	})))
	router.HandleFunc("/home", w.stateManager.SessionWrapper(w.loginWrapper(w.homeIdeas))).Methods(http.MethodGet)
	router.HandleFunc("/search", w.stateManager.SessionWrapper(w.loginWrapper(w.searchIdeas))).Methods(http.MethodPost)
	router.HandleFunc("/idea", w.stateManager.SessionWrapper(w.loginWrapper(w.getIdea))).Methods(http.MethodGet)
	router.HandleFunc("/idea", w.stateManager.SessionWrapper(w.loginWrapper(w.makeIdea))).Methods(http.MethodPost)
	router.HandleFunc("/idea", w.stateManager.SessionWrapper(w.loginWrapper(w.modifyIdea))).Methods(http.MethodPatch, http.MethodDelete)

	return &http.Server{
		Addr:    os.Getenv("LISTEN"),
		Handler: router,
	}
}

func (w *Web) loginWrapper(cb func(rw http.ResponseWriter, req *http.Request, state *utils.State, loggedIn bool, user structure.User)) func(rw http.ResponseWriter, req *http.Request, state *utils.State) {
	return func(rw http.ResponseWriter, req *http.Request, state *utils.State) {
		if myUser, ok := utils.GetStateValue[*structure.User](state, utils.KeyUser); ok {
			if myUser == nil {
				cb(rw, req, state, false, structure.User{})
				return
			}
			cb(rw, req, state, true, *myUser)
			return
		}
		cb(rw, req, state, false, structure.User{})
	}
}

func (w *Web) loginPage(rw http.ResponseWriter, req *http.Request, state *utils.State) {
	q := req.URL.Query()
	if q.Has("in_popup") {
		state.Put("login-in-popup", true)
	}
	if myUser, ok := utils.GetStateValue[*structure.User](state, utils.KeyUser); ok {
		if myUser != nil {
			if doLoginPopup(rw, w.originUrl, state, myUser) {
				return
			}
			http.Redirect(rw, req, "/", http.StatusTemporaryRedirect)
			return
		}
	}
	if flowState, ok := utils.GetStateValue[uuid.UUID](state, utils.KeyState); ok {
		q := req.URL.Query()
		if q.Has("code") && q.Has("state") {
			if q.Get("state") == flowState.String() {
				exchange, err := w.oauthClient.Exchange(context.Background(), q.Get("code"))
				if err != nil {
					fmt.Println("Exchange token error:", err)
					return
				}
				state.Put(utils.KeyAccessToken, exchange.AccessToken)
				state.Put(utils.KeyRefreshToken, exchange.RefreshToken)

				c, err := discordgo.New("Bearer " + exchange.AccessToken)
				if err != nil {
					fmt.Println("Create client error:", err)
					return
				}

				user, err := c.User("@me")
				if err != nil {
					return
				}
				var meta structure.User
				b, err := w.engine.Where("discord_id").Get(&meta)
				if err != nil {
					return
				}
				if !b {
					meta = structure.User{
						DiscordId:            user.ID,
						DiscordName:          user.Username,
						DiscordDiscriminator: user.Discriminator,
						Picture:              user.AvatarURL("256"),
						Banned:               false,
						Admin:                false,
					}
					_, err := w.engine.Insert(&meta)
					if err != nil {
						http.Error(rw, "Failed to insert new user into database", http.StatusInternalServerError)
						return
					}
				}
				if meta.DiscordName != user.Username || meta.DiscordDiscriminator != user.Discriminator {
					meta.DiscordName = user.Username
					meta.DiscordDiscriminator = user.Discriminator
					_, err := w.engine.Update(meta, structure.User{DiscordId: user.ID})
					if err != nil {
						http.Error(rw, "Failed to update Discord Tag in database", http.StatusInternalServerError)
						return
					}
				}
				state.Put(utils.KeyUser, &meta)

				if doLoginPopup(rw, w.originUrl, state, &meta) {
					return
				}
				http.Redirect(rw, req, "/", http.StatusTemporaryRedirect)
				return
			}
			http.Error(rw, "OAuth flow state doesn't match\n", http.StatusBadRequest)
			return
		}
	}
	flowState := uuid.New()
	state.Put(utils.KeyState, flowState)
	http.Redirect(rw, req, w.oauthClient.AuthCodeURL(flowState.String(), oauth2.AccessTypeOffline), http.StatusTemporaryRedirect)
}

func (w *Web) checkPage(rw http.ResponseWriter, _ *http.Request, state *utils.State) {
	if myUser, ok := utils.GetStateValue[*structure.User](state, utils.KeyUser); ok {
		if myUser != nil {
			exportUserDataAsJson(rw, w.originUrl, myUser, true)
			return
		}
	}
	rw.WriteHeader(http.StatusBadRequest)
}

func doLoginPopup(rw http.ResponseWriter, originUrl string, state *utils.State, meta *structure.User) bool {
	if b, ok := utils.GetStateValue[bool](state, "login-in-popup"); ok {
		if b {
			exportUserDataAsJson(rw, originUrl, meta, false)
			return true
		}
	}
	return false
}

func exportUserDataAsJson(rw http.ResponseWriter, originUrl string, meta *structure.User, checkMode bool) {
	start := LoginFrameStart
	end := LoginFrameEnd
	if checkMode {
		start = CheckFrameStart
		end = CheckFrameEnd
	}
	_, _ = rw.Write([]byte(start))
	encoder := json.NewEncoder(rw)
	_ = encoder.Encode(meta)
	_, _ = rw.Write([]byte(fmt.Sprintf(end, originUrl)))
}
