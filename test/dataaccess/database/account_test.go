package database_test

import (
	"context"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/Fiagram/account_service/internal/dataaccess/database"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := database.Account{
		Username:    randomString(10),
		Fullname:    randomVnPersonName(),
		Email:       randomString(10) + "@gmail.com",
		PhoneNumber: "+84 123456789",
		RoleId:      1,
	}
	id, err := asor.CreateAccount(context.Background(), input)

	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotZero(t, id)

	errD := asor.DeleteAccountByUsername(context.Background(), input.Username)
	require.NoError(t, errD)
}

func TestGetAccountById(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := database.Account{
		Username:    randomString(10),
		Fullname:    randomVnPersonName(),
		Email:       randomString(10) + "@gmail.com",
		PhoneNumber: randomVnPhoneNum(),
		RoleId:      2,
	}
	id, err := asor.CreateAccount(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	acc, err := asor.GetAccountById(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, input.Username, acc.Username)
	require.Equal(t, input.Fullname, acc.Fullname)
	require.Equal(t, input.Email, acc.Email)
	require.Equal(t, input.PhoneNumber, acc.PhoneNumber)
	require.Equal(t, input.RoleId, acc.RoleId)

	errD := asor.DeleteAccountByUsername(context.Background(), input.Username)
	require.NoError(t, errD)
}

func TestGetAccountByUsername(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := database.Account{
		Username:    randomString(10),
		Fullname:    randomVnPersonName(),
		Email:       randomString(10) + "@gmail.com",
		PhoneNumber: randomVnPhoneNum(),
		RoleId:      2,
	}
	id, err := asor.CreateAccount(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	acc, err := asor.GetAccountByUsername(context.Background(), input.Username)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, input.Username, acc.Username)
	require.Equal(t, input.Fullname, acc.Fullname)
	require.Equal(t, input.Email, acc.Email)
	require.Equal(t, input.PhoneNumber, acc.PhoneNumber)
	require.Equal(t, input.RoleId, acc.RoleId)

	errD := asor.DeleteAccountByUsername(context.Background(), input.Username)
	require.NoError(t, errD)
}

func TestDeleteAccountById(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := database.Account{
		Username:    randomString(10),
		Fullname:    randomVnPersonName(),
		Email:       randomString(10) + "@gmail.com",
		PhoneNumber: randomVnPhoneNum(),
		RoleId:      2,
	}
	id, err := asor.CreateAccount(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	errD := asor.DeleteAccountById(context.Background(), id)
	require.NoError(t, errD)
}

func TestDeleteAccountByUsername(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := database.Account{
		Username:    randomString(10),
		Fullname:    randomVnPersonName(),
		Email:       randomString(10) + "@gmail.com",
		PhoneNumber: randomVnPhoneNum(),
		RoleId:      2,
	}
	id, err := asor.CreateAccount(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	errD := asor.DeleteAccountByUsername(context.Background(), input.Username)
	require.NoError(t, errD)
}

// Following functions serve account_test suit
func randomString(length uint) string {
	g := rand.New(rand.NewSource(time.Now().UnixNano()))
	const alphabet = "qazwsxedcrfvtgbyhnujmikolp"
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < int(length); i++ {
		c := alphabet[g.Intn(k)]
		sb.WriteByte(c)
	}
	return strings.TrimSpace(sb.String())
}

func randomVnPersonName() string {
	lastnames := []string{
		"Nguyễn", "Vũ", "Trần", "Huỳnh", "Lê", "Phạm",
		"Phan", "Hoàng", "Phùng", "Tô", "Mai", "Trương"}
	middles := []string{
		"Đỗ", "Đức", "Mạnh", "Thị", "Uyển", "Lâm",
		"Văn", "Hàn", "Thùy", "Anh", "Duy", "Khánh",
	}
	firstnames := []string{
		"Thế", "Tuấn", "Trung", "Hùng", "Dũng", "Tân",
		"Hà", "Trí", "Hiếu", "Thái", "Tiến", "Ngọc",
	}
	g := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]string, 4)
	out = append(out, lastnames[g.Intn(len(lastnames))])
	out = append(out, middles[g.Intn(len(middles))])
	out = append(out, middles[g.Intn(len(middles))])
	out = append(out, firstnames[g.Intn(len(firstnames))])
	return strings.TrimSpace(strings.Join(out, " "))
}

func randomVnPhoneNum() string {
	g := rand.New(rand.NewSource(time.Now().UnixNano()))
	const nums = "0123456789"
	var sb strings.Builder
	k := len(nums)
	for i := range 9 {
		c := nums[rand.Intn(k)]
		for c == '0' && i == 0 {
			c = nums[g.Intn(k)]
		}
		sb.WriteByte(c)
	}
	return strings.TrimSpace("+84 " + sb.String())
}
