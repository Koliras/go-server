package api

import (
	"strconv"

	"github.com/Koliras/go-server/model"
	"github.com/gin-gonic/gin"
)

func (s *Server) getAllUsers(ctx *gin.Context) {
	var users []model.User
	result := s.store.Find(&users)
	if result.Error != nil {
		ctx.JSON(400, gin.H{
			"message": "Could not create the user",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": users,
	})
}

func (s *Server) getUserById(ctx *gin.Context) {
    stringId := ctx.Param("id")
    id, err := strconv.Atoi(stringId)
    if err != nil {
        ctx.JSON(404, gin.H{
            "message": "ID has to be a valid number",
        })
        return
    }

    var user model.User
    result := s.store.First(&user, id)
    if result.Error != nil {
        ctx.JSON(404, gin.H{
            "message": "There is no user with such id",
        })
        return
    }

    ctx.JSON(200, gin.H{
        "data": user,
    })
}

func (s *Server) createUser(ctx *gin.Context) {
    type Body struct {
        Username string `json:"username"`
    }
    body := Body{}
    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(500, gin.H{
            "message": "Internal server error",
        })
        return
    }

    user := model.User{Username: body.Username}
    if result := s.store.Create(&user); result.Error != nil {
        ctx.JSON(400, gin.H{
            "message": "Could not create the user",
        })
        return
    }

    ctx.JSON(201, gin.H{
        "message": "Successfully created new user",
    })
}

func (s *Server) deleteUser(ctx *gin.Context) {
    stringId := ctx.Param("id")
    id, err := strconv.Atoi(stringId)
    if err != nil {
        ctx.JSON(404, gin.H{
            "message": "ID has to be a valid number",
        })
        return
    }

    if result := s.store.Delete(&model.User{}, id); result.Error != nil {
        ctx.JSON(500, gin.H{
            "message": "Could not delete the user",
        })
        return
    }

    ctx.JSON(200, gin.H{
        "message": "Successfully deleted the user",
    })
}

func (s *Server) updateUser(ctx *gin.Context) {
    type Body struct {
        Username string `json:"username"`
    }
    body := Body{}
    if err := ctx.BindJSON(&body); err != nil {
        ctx.JSON(500, gin.H{
            "message": "Internal server error",
        })
        return
    }

    stringId := ctx.Param("id")
    id, err := strconv.Atoi(stringId)
    if err != nil {
        ctx.JSON(400, gin.H{
            "message": "ID has to be a valid number",
        })
        return
    }

    var user model.User
    if result := s.store.First(&user, id); result.Error != nil {
        ctx.JSON(404, gin.H{
            "message": "Could not find the user with such id",
        })
        return
    }


    if len(body.Username) != 0 {
        user.Username = body.Username
    }
    s.store.Save(&user)

    ctx.JSON(200, gin.H{
        "message": "Successfully updated the user",
    })
}

