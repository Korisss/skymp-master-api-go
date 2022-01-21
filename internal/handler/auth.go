package handler

import (
	"net/http"

	master_api "github.com/Korisss/skymp-master-api-go"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userNameRequest struct {
	Id int `uri:"id" binding:"required"`
}

// Returning user name
// TODO: check id with token match
func (h *Handler) getUserName(ctx *gin.Context) {
	ctxId, _ := ctx.Get("id")
	id := ctxId.(int)

	var req userNameRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if id != req.Id {
		newErrorResponse(ctx, http.StatusForbidden, "no access to requested information")
		return
	}

	name, err := h.services.GetUserName(req.Id)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"name": name,
	})
}

func (h *Handler) login(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, token, err := h.services.GenerateToken(req.Email, req.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id":    id,
		"token": token,
	})
}

func (h *Handler) register(ctx *gin.Context) {
	var req master_api.User

	if err := ctx.BindJSON(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(req)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) resetPassword(ctx *gin.Context) {}

// TODO: check id with token match
func (h *Handler) createSession(ctx *gin.Context) {}

// static async play(ctx: Context | Router.RouterContext): Promise<void> {
//     const user = (ctx as Record<string, User>).user;
//     if (!(await UserController.ensureTokenMatchesId(ctx))) return;
//     user.currentServerAddress = ctx.params.serverAddress;
//     user.currentSession = randomString(32);
//     await UserController.getRepository(ctx).save(user);
//     ctx.status = 200;
//     ctx.body = {
//       session: user.currentSession
//     };
//   }
