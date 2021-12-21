package transport

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/kott/go-service-example/pkg/errors"
	"github.com/kott/go-service-example/pkg/services/users"
	"github.com/kott/go-service-example/pkg/services/users/store"
	"github.com/kott/go-service-example/pkg/utils/context"
	"github.com/kott/go-service-example/pkg/utils/log"
	"net/http"
)

type handler struct {
	us users.Service
}

func newHandler(us users.Service) *handler {
	return &handler{us: us}
}

func Activate(r *gin.Engine, db *sql.DB) {
	h := newHandler(users.New(store.New(db)))
	r.GET("/users/:id", h.Get)
	r.GET("/users", h.GetAll)
	r.POST("/users", h.Create)
	r.POST("/users/:id", h.Update)
}

func (h *handler) Get(c *gin.Context) {
	ctx := context.GetReqCtx(c)

	id := c.Param("id")
	log.Info(ctx, "Querying user id=%s", id)
	user, err := h.us.Get(ctx, id)
	if err != nil {
		code, appErr := handleError(err)
		c.IndentedJSON(code, appErr)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *handler) GetAll(c *gin.Context) {
	ctx := context.GetReqCtx(c)

	var param struct {
		Limit int `form:"limit,default=10"`
		Offset int `form:"offset,default=0"`
	}
	if err := c.BindQuery(&param); err != nil {
		log.Error(ctx, "limit or offset not provided in query, %v", param)
		c.IndentedJSON(http.StatusBadRequest, errors.NewAppError(errors.BadRequest, errors.Descriptions[errors.BadRequest], ""))
		return
	}

	log.Info(ctx, "get all users with limit=%d, offset=%d", param.Limit, param.Offset)
	ul, err := h.us.GetAll(ctx, param.Limit, param.Offset)
	if err != nil {
		code, appErr := handleError(err)
		c.IndentedJSON(code, appErr)
		return
	}
	// 这里在JSON中再包一层key users
	c.IndentedJSON(http.StatusOK, users.Users{Users: ul})
}

func (h *handler) Create(c *gin.Context) {
	ctx := context.GetReqCtx(c)

	var uc users.UserCreateUpdate
	if err := c.ShouldBindJSON(&uc); err != nil {
		log.Error(ctx, "request parse error: %s", err.Error())
		c.IndentedJSON(http.StatusBadRequest, errors.NewAppError(errors.BadRequest, errors.Descriptions[errors.BadRequest], ""))
		return
	}

	log.Info(ctx, "creating user %v", uc)
	user, err := h.us.Create(ctx, uc)
	if err != nil {
		code, appErr := handleError(err)
		c.IndentedJSON(code, appErr)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *handler) Update(c *gin.Context) {

}

func handleError(e error) (int, error) {
	switch e {
	case users.ErrUserNotFound:
		return http.StatusNotFound, errors.NewAppError(errors.NotFound, "User could not be found", "id")
	case users.ErrUserCreate:
		fallthrough
	case users.ErrUserUpdate:
		return http.StatusInternalServerError, errors.NewAppError(errors.InternalServerError, "unable to create / update user", "")
	default:
		return http.StatusInternalServerError, errors.NewAppError(errors.InternalServerError, e.Error(), "unknown")
	}
}