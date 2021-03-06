package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/helper"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user entity.UserEntity, username string) string
	Update(ctx context.Context, tx *sql.Tx, user entity.UserEntity) string
	Delete(ctx context.Context, tx *sql.Tx, username string) string
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) []entity.UserEntity

	FindAll(ctx context.Context, tx *sql.Tx, where string) []entity.UserEntity
	VerifyCredential(ctx context.Context, tx *sql.Tx, username string, password string) []entity.UserEntity
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user entity.UserEntity, username string) string {
	cuyNow := helper.TimePlus7(time.Now())
	user.Role = "user"
	user.CreatedTime = cuyNow
	user.CreatedBy = username

	SQL := fmt.Sprintf("insert into users (username,password,role,created_by,created_time,updated_by) values ('%s','%s','%s','%s','%s','%s') ", user.Username, user.Password, user.Role, user.CreatedBy, user.CreatedTime, user.UpdatedBy)
	fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	fmt.Println(err, row)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	fmt.Println(errs, rows)
	if rows < 1 {
		helper.PanicIfError(errs)
	}
	fmt.Println(rows)
	if rows > 0 {
		return "berhasil"
	} else {
		return "gagal"
	}

}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.UserEntity) string {
	fmt.Println("hay")
	cuyNow := helper.TimePlus7(time.Now())
	user.UpdatedTime = cuyNow
	user.UpdatedBy = user.Username
	SQL := fmt.Sprintf("update users set password='%s',updated_by='%s',updated_time='%s' where username='%s'", user.Password, user.UpdatedBy, user.UpdatedTime, user.Username)
	fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)
	if rows > 0 {
		return "berhasil"
	} else {
		return "gagal"
	}
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, username string) string {
	SQL := fmt.Sprintf("delete users where username='%s'", username)
	fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)
	if rows > 0 {
		return "berhasil"
	} else {
		return "gagal"
	}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) []entity.UserEntity {
	SQL := fmt.Sprintf("select id,username,role,created_by,created_time,updated_by,IFNULL(updated_time,'') as updated_time from users where username='%s'", username)
	var datas []entity.UserEntity
	var data entity.UserEntity
	row, err := tx.QueryContext(ctx, SQL)
	fmt.Println(err, username)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.RowId, &data.Username, &data.Role, &data.CreatedBy, &data.CreatedTime, &data.UpdatedBy, &data.UpdatedTime)
		helper.PanicIfError(err)
		datas = append(datas, data)
	}
	return datas
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, where string) []entity.UserEntity {
	SQL := fmt.Sprintf("select id,username,role,created_by,created_time,updated_by,IFNULL(updated_time,'') as updated_time from users where %s", where)
	fmt.Println(SQL)
	var datas []entity.UserEntity
	var data entity.UserEntity
	row, err := tx.QueryContext(ctx, SQL)
	fmt.Println(err)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.RowId, &data.Username, &data.Role, &data.CreatedBy, &data.CreatedTime, &data.UpdatedBy, &data.UpdatedTime)
		fmt.Println(err)

		helper.PanicIfError(err)
		// data.CreatedTime = helper.ConvertDateTime(data.CreatedTime)
		// data.UpdatedTime = helper.ConvertDateTime(data.UpdatedTime)

		datas = append(datas, data)
	}
	return datas
}
func (repository *UserRepositoryImpl) VerifyCredential(ctx context.Context, tx *sql.Tx, username string, password string) []entity.UserEntity {
	SQL := fmt.Sprintf("select username,role,created_by,created_time,updated_by,IFNULL(updated_time,'') as updated_time from users where username='%s' and password='%s'", username, password)
	fmt.Println(SQL)
	var datas []entity.UserEntity
	var data entity.UserEntity
	row, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.Username, &data.Role, &data.CreatedBy, &data.CreatedTime, &data.UpdatedBy, &data.UpdatedTime)

		helper.PanicIfError(err)
		datas = append(datas, data)
	}
	fmt.Println(datas)
	return datas

}
