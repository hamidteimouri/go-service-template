package mysql

import (
	"context"
	"gorm.io/gorm"
	"laramanpurego/internal/domain/entity"
)

type mysql struct {
	db *gorm.DB
}

func NewMysql(db *gorm.DB) *mysql {
	return &mysql{db: db}
}

func InsertUser(ctx context.Context, userEntity *entity.User) {

	userModel := UserModel{}
	userModel.ConvertEntityToModel(userEntity)

}
