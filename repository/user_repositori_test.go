package repository

import (
	"context"
	"golang-database-user/config"
	"golang-database-user/model"
	"testing"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser_Success(t *testing.T) {

    sql, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        panic(err)
    }

    newUserRepositoryImpl := NewUserRepositoryImpl(sql)
    newRoleRepositoryImpl := NewRoleRepositoryImpl(sql)

    ctx := context.Background()

    role, err := newRoleRepositoryImpl.FindMstRole(ctx, "ROLE001")

    mstUser := model.MstUser{
        IdUser:      uuid.NewString(),
        Name:        "Anuu",
        Email:       "Anu@gmail.com",
        Password:    "Anuu",
        PhoneNumber: "089580",
        Role:        role,
    }

    insertUser, err := newUserRepositoryImpl.InsertUser(ctx, mstUser)

    if err != nil {
        panic(err)
    }

    assert.NotNil(t, insertUser)
    assert.Equal(t, mstUser, insertUser)
    assert.Equal(t, mstUser.IdUser, insertUser.IdUser)
    assert.Equal(t, mstUser.Email, insertUser.Email)
    assert.Equal(t, mstUser.Role, insertUser.Role)
}

func TestInsertUser_Fail(t *testing.T) {
    sql, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        panic(err)
    }

    newUserRepositoryImpl := NewUserRepositoryImpl(sql)
    newRoleRepositoryImpl := NewRoleRepositoryImpl(sql)

    ctx := context.Background()

    role, err := newRoleRepositoryImpl.FindMstRole(ctx, "ROLE001")
    if err != nil {
        panic(err)
    }

    existingUser := model.MstUser{
        IdUser:      uuid.NewString(),
        Name:        "z",
        Email:       "Anuu@gmail.com", // email yang sudah ada didatabase
        Password:    "0",
        PhoneNumber: "0",
        Role:        role,
    }

    //insertttt
    insertUser, err := newUserRepositoryImpl.InsertUser(ctx, existingUser)

    // Verif kalau emailnya sudah ada didatabase isi err adaaaaa 
    assert.Equal(t, model.MstUser{}, insertUser, "err harusnya karna email dupli")

    assert.NotNil(t, err, "Errornya sehrusnya tidakk nil atau ada karena insert gagall atau nilai var err ada ")
}


func TestReadUser_Success(t *testing.T) {
    sql, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        panic(err)
    }

    newUserRepositoryImpl := NewUserRepositoryImpl(sql)
    ctx := context.Background()

    users, err := newUserRepositoryImpl.ReadUser(ctx)

    assert.Nil(t, err, "Tidak boleh ada error saat membaca daftar pengguna")

    // isi tabel user
    assert.NotNil(t, users, "isi tabel user harusnya ada")
    assert.True(t, len(users) > 0)

    // periksa kalau isi dari tabelnya isi kolom yang diselect valid smua atau tidak ada yang ksongg
    for _, user := range users {
        assert.NotEmpty(t, user.Name, "kol nama ada yngg kkosong")
        assert.NotEmpty(t, user.Email, "kol email ada kkosong")
        assert.NotEmpty(t, user.PhoneNumber, "kol nohhp ada yang kosonng")
    }
}

func TestReadUser_Fail(t *testing.T) {
    sql, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        panic(err)
    }

    newUserRepositoryImpl := NewUserRepositoryImpl(sql)
    ctx := context.Background()

    users, err := newUserRepositoryImpl.ReadUser(ctx)

    // periksa dari func readd user tidak ada malsalh di pembacaan data
    assert.Nil(t, err, "Tidak boleh ada error saat membaca daftar pengguna")

    // Hasil harus slice kosong karena tabel kosong
    assert.NotNil(t, users, "Hasil users tidak boleh nil")

    // untuk pengecekan apakah tabel user kosong jika kosong isi lenn atau panjang users == 0 
    // assert.Equal(t, 0, len(users), "Jumlah users harus 0 jika tabel kosong")
}

func TestUpdateUser_Success(t *testing.T) {
    sql, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        panic(err)
    }

    newUserRepositoryImpl := NewUserRepositoryImpl(sql)
    ctx := context.Background()

    // Data update
    updatedUser := model.MstUser{
        Name:        "aksan",
        Email:       "aksan@coconut.or.id",
        Password:    "passnew",
        PhoneNumber: "08987654",
    }

    userId := "8e211b35-903e-4b5a-a1a6-ab04b24d12fe"

    // coba updatee
    result, err := newUserRepositoryImpl.UpdateUser(ctx, updatedUser, userId)

    assert.Nil(t, err, "no err saat update")
    assert.NotNil(t, result, "harus ada yang diubah")
    assert.Equal(t, updatedUser.Name, result.Name, "nama ganti")
    assert.Equal(t, updatedUser.Email, result.Email, "email ganti")
    assert.Equal(t, updatedUser.Password, result.Password, "ganti pass")
    assert.Equal(t, updatedUser.PhoneNumber, result.PhoneNumber, "perbaharui noo")
}

func TestUpdateUser_Fail(t *testing.T) {
    sql, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        panic(err)
    }

    newUserRepositoryImpl := NewUserRepositoryImpl(sql)
    ctx := context.Background()

    // Data update
    updatedUser := model.MstUser{
        Name:        "aksan", 
        Email:       "aksan@coconut.or.id",
        Password:    "passnew",
        PhoneNumber: "08987654",
    }

    userId := "id salah-salah"

    // coba updatee
    result, err := newUserRepositoryImpl.UpdateUser(ctx, updatedUser, userId)

    // Asersi
    assert.NotNil(t, err, "isi var err ada karna id dalah dan update gagal")
    assert.Equal(t, model.MstUser{}, result, "update tidak adaa karena  tidak ada pengguna yang diperbarui atau diubah karena ID yang diberikan tidak valid")
}

func TestDeleteUser_Success(t *testing.T) {
    sql, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        panic(err)
    }

    newUserRepositoryImpl := NewUserRepositoryImpl(sql)
    ctx := context.Background()
    
    validUserId := "8e211b35-903e-4b5a-a1a6-ab04b24d12fe" 

    // Coba hapus user
    err = newUserRepositoryImpl.DeleteUser(ctx, validUserId)

    // Assert bahwa tidak ada error
    assert.Nil(t, err, "ad err diproses penghapusan")
}

func TestDeleteUser_Fail(t *testing.T) {
    sql, err := config.OpenConnectionPostgresSQL()
    if err != nil {
        panic(err)
    }

    newUserRepositoryImpl := NewUserRepositoryImpl(sql)
    ctx := context.Background()

    userId := "id salahh"

    // Coba hapus user
    err = newUserRepositoryImpl.DeleteUser(ctx, userId)

    // Assert bahwa error terjadi
    assert.NotNil(t, err, "Harusnya ada err")
}