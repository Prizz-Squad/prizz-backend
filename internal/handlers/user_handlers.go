package handlers

import (
	"github.com/EraldCaka/prizz-backend/internal/services"
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/EraldCaka/prizz-backend/util"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strings"
	"time"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var user *types.UserCreateRequest
	if err := ctx.BodyParser(&user); err != nil {
		return types.ErrBadRequest()
	}
	if validate := user.Validate(); len(validate) > 0 {
		return ctx.JSON(validate)
	}
	userID, err := h.userService.Create(ctx.Context(), user)
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't register the user"))
	}
	return ctx.JSON(userID)
}

func (h *UserHandler) GetUsers(ctx *fiber.Ctx) error {
	users, err := h.userService.GetAll(ctx.Context())
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get the users"))
	}
	return ctx.JSON(users)
}

func (h *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	var userID = ctx.Params("id")
	user, err := h.userService.GetByID(ctx.Context(), userID)
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't get user"))
	}
	return ctx.JSON(user)
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var (
		userID  = ctx.Params("id")
		userReq *types.UserRequest
	)
	if err := ctx.BodyParser(&userReq); err != nil {
		return types.ErrBadRequest()
	}
	if err := h.userService.Update(ctx.Context(), userID, userReq); err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't update user"))
	}
	return ctx.JSON(userID)
}

func (h *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	var userID = ctx.Params("id")
	err := h.userService.Delete(ctx.Context(), userID)
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusInternalServerError, "Couldn't delete user"))
	}
	return ctx.JSON(userID)
}

func (h *UserHandler) LoginUser(ctx *fiber.Ctx) error {
	if exist := ctx.Get("Authorization"); exist != "" {
		return ctx.JSON(types.NewError(http.StatusBadRequest, "an account has already logged in on this browser"))
	}
	var user *types.UserRequest
	if err := ctx.BodyParser(&user); err != nil {
		return types.ErrBadRequest()
	}

	userID, err := h.userService.Login(ctx.Context(), user)
	if err != nil {
		return ctx.JSON(types.NewError(http.StatusForbidden, "Unsuccessful login"))
	}

	tokenData := types.JWTToken{
		ID:       userID,
		Username: user.Username,
	}
	token, err := services.CreateJWTToken(tokenData, util.Secret, time.Hour)
	if err != nil {
		log.Println("Error while creating jwt token")
		return types.InternalServerError()
	}
	ctx.Set("Authorization", token)
	return ctx.JSON(fiber.Map{"token": token})
}

func (h *UserHandler) LogoutUser(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.JSON(types.NewError(http.StatusForbidden, "you are not logged in"))
	}
	ctx.Set("Authorization", "")
	return nil

}

func (h *UserHandler) RoleBaseMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		user, err := h.userService.GetUserByToken(ctx.Context(), token)
		if err != nil {
			return ctx.JSON(types.NewError(http.StatusInternalServerError, err.Error()))
		}
		switch user.Role {
		case types.MemberRole:
			for _, route := range util.MemberRoutes {
				if ctx.Path() == route || strings.Contains(ctx.Path(), route) {
					return ctx.Next()
				}
			}
		case types.ClientRole:
			for _, route := range util.CustomerRoutes {
				if ctx.Path() == route || strings.Contains(ctx.Path(), route) {
					return ctx.Next()
				}
			}
		case types.AdminRole:
			for _, route := range util.AdminRoutes {
				if ctx.Path() == route || strings.Contains(ctx.Path(), route) {
					return ctx.Next()
				}
			}
		}
		return nil
	}
}
