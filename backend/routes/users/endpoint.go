package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"user-management-api/domain"
)

func CreateUserHandler(userService Service) domain.Endpoint {
	return func(ctx *gin.Context, request interface{}) (interface{}, error) {
		user, ok := request.(*domain.User)
		if !ok {
			return nil, errors.New("bad request")
		}
		if err := userService.CreateUser(user); err != nil {
			return nil, err
		}

		return gin.H{"created": request}, nil
	}
}

func GetUsersHandler(userService Service) domain.Endpoint {
	return func(ctx *gin.Context, request interface{}) (interface{}, error) {
		pageStr := ctx.Query("page")
		pageSizeStr := ctx.Query("pageSize")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			//Invalid page received reverting to default
			page = 0
		}
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			//Invalid page received reverting to default
			pageSize = 0
		}

		users, err := userService.GetUsers(page, pageSize)
		if err != nil {
			return nil, err
		}
		response := gin.H{
			"users": users,
		}
		total, err := userService.Count()
		if err == nil {
			response["total"] = total
		}
		return response, nil
	}
}

func GetUserHandler(userService Service) domain.Endpoint {
	return func(ctx *gin.Context, request interface{}) (response interface{}, err error) {
		idStr := ctx.Param("id")
		userId, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		if userId <= 0 {
			return nil, errors.New("invalid id")
		}
		user, err := userService.GetUser(userId)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

func UpdateUserHandler(userService Service) domain.Endpoint {
	return func(ctx *gin.Context, request interface{}) (interface{}, error) {
		idStr := ctx.Param("id")
		userId, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}

		user, ok := request.(*domain.User)
		if !ok {
			return nil, errors.New("bad request")
		}
		err = userService.UpdateUser(userId, user)
		if err != nil {
			return nil, err
		}
		response := gin.H{"updated": request}
		total, err := userService.Count()
		if err == nil {
			response["total"] = total
		}

		return response, nil
	}
}

func DeleteUserHandler(userService Service) domain.Endpoint {
	return func(ctx *gin.Context, request interface{}) (interface{}, error) {
		idStr := ctx.Param("id")
		userId, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}

		err = userService.DeleteUser(userId)
		if err != nil {
			return nil, err
		}
		response := gin.H{}
		total, err := userService.Count()
		if err == nil {
			response["total"] = total
		}
		return response, nil
	}
}

func SearchUsersHandler(userService Service) domain.Endpoint {
	return func(ctx *gin.Context, request interface{}) (interface{}, error) {
		query := ctx.Query("query")
		pageStr := ctx.Query("page")
		pageSizeStr := ctx.Query("pageSize")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			//Invalid page received reverting to default
			page = 0
		}
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			//Invalid page received reverting to default
			pageSize = 0
		}
		users, total, err := userService.SearchUsers(query, page, pageSize)
		if err != nil {
			return nil, err
		}

		response := gin.H{
			"users": users,
			"total": total,
		}
		return response, nil
	}
}
