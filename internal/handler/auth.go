package handler

import (
	"net/http"

	master_api "github.com/Korisss/skymp-master-api-go"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type requestWithAuth struct {
	Id int `uri:"id" binding:"required"`
}

// Returning user name
func (h *Handler) getUserName(ctx *gin.Context) {
	id, access := checkUserAccess(ctx)

	if !access {
		return
	}

	name, err := h.services.GetUserName(id)

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

// TODO: create with verification
func (h *Handler) resetPassword(ctx *gin.Context) {}

func checkUserAccess(ctx *gin.Context) (int, bool) {
	ctxId, _ := ctx.Get("id")
	id := ctxId.(int)

	var req requestWithAuth

	if err := ctx.ShouldBindUri(&req); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return id, false
	}

	if id != req.Id {
		newErrorResponse(ctx, http.StatusForbidden, "no access to create session for this user")
		return id, false
	}

	return id, true
}

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
