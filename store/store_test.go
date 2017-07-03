package store

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func newMockFloresta() (Store, sqlmock.Sqlmock) {
	// Everything should work fine
	db, mock, _ := sqlmock.New()

	floresta := NewStore(db)

	return floresta, mock
}

func Test_NewStore(t *testing.T) {
	db, _, err := sqlmock.New()

	if err != nil {
		t.Errorf("Could not create a mock database: %q", err.Error())
	}

	floresta := NewStore(db)

	if _, ok := floresta.(Store); !ok {
		t.Errorf("NewStore must return an store.Store")
	}

	defer db.Close()
}

func Test_GetWord(t *testing.T) {
	floresta, mock := newMockFloresta()

	rows := sqlmock.NewRows([]string{
		"id",
		"arvore",
		"palavra",
		"lema",
		"funcao",
		"froma",
		"morfo",
		"pai",
	}).AddRow(
		42,
		420,
		"palavra",
		"lema",
		"funcao",
		"forma",
		"morfo",
		"pai",
	)

	mock.ExpectQuery(getWordQuery).WithArgs(anyString{}).WillReturnRows(rows)

	_, err := floresta.GetWord("some word")

	if err != nil {
		t.Errorf("Did not expect an error: %q", err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_GetWord_queryError(t *testing.T) {
	floresta, mock := newMockFloresta()

	mock.ExpectQuery(getWordQuery).WithArgs(anyString{}).WillReturnError(errors.New("Some error"))

	_, err := floresta.GetWord("some word")

	if err == nil {
		t.Error("Expecting an error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_GetWord_RowsError(t *testing.T) {
	floresta, mock := newMockFloresta()

	rows := sqlmock.NewRows([]string{
		"id",
		"arvore",
		"palavra",
		"lema",
		"funcao",
		"froma",
		"morfo",
		"pai",
	})

	for i := 0; i < 5; i++ {
		rows.AddRow(
			42+i,
			420,
			fmt.Sprintf("palavra-%d", i),
			"lema",
			"funcao",
			"forma",
			"morfo",
			"pai",
		)
	}

	rows = rows.RowError(2, errors.New("Some row error"))

	mock.ExpectQuery(getWordQuery).WithArgs(anyString{}).WillReturnRows(rows)

	ramos, err := floresta.GetWord("some word")

	if err == nil {
		t.Error("Expecting an error")
	}

	if len(ramos) != 2 {
		t.Errorf("len(ramos) = %d, expected: %d", len(ramos), 2)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_GetWord_ScanError(t *testing.T) {
	floresta, mock := newMockFloresta()

	rows := sqlmock.NewRows([]string{
		"id",
		"arvore",
		"palavra",
		"lema",
		"funcao",
		"froma",
		"morfo",
	}).AddRow(
		42,
		420,
		"palavra",
		"lema",
		"funcao",
		"forma",
		"morfo",
	)

	mock.ExpectQuery(getWordQuery).WithArgs(anyString{}).WillReturnRows(rows)

	ramos, err := floresta.GetWord("some word")

	if err == nil {
		t.Error("Expecting an error")
	}

	if len(ramos) != 0 {
		t.Errorf("len(ramos) = %d, expected: %d", len(ramos), 2)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_GetLema(t *testing.T) {
	floresta, mock := newMockFloresta()

	rows := sqlmock.NewRows([]string{
		"id",
		"arvore",
		"palavra",
		"lema",
		"funcao",
		"froma",
		"morfo",
		"pai",
	}).AddRow(
		42,
		420,
		"palavra",
		"lema",
		"funcao",
		"forma",
		"morfo",
		"pai",
	)

	mock.ExpectQuery(getLemaQuery).WithArgs(anyString{}).WillReturnRows(rows)

	_, err := floresta.GetLema("some lema")

	if err != nil {
		t.Errorf("Did not expect an error: %q", err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_GetLema_queryError(t *testing.T) {
	floresta, mock := newMockFloresta()

	mock.ExpectQuery(getLemaQuery).WithArgs(anyString{}).WillReturnError(errors.New("Some error"))

	_, err := floresta.GetLema("some lema")

	if err == nil {
		t.Error("Expecting an error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_GetLema_RowsError(t *testing.T) {
	floresta, mock := newMockFloresta()

	rows := sqlmock.NewRows([]string{
		"id",
		"arvore",
		"palavra",
		"lema",
		"funcao",
		"froma",
		"morfo",
		"pai",
	})

	for i := 0; i < 5; i++ {
		rows.AddRow(
			42+i,
			420,
			fmt.Sprintf("palavra-%d", i),
			"lema",
			"funcao",
			"forma",
			"morfo",
			"pai",
		)
	}

	rows = rows.RowError(2, errors.New("Some row error"))

	mock.ExpectQuery(getLemaQuery).WithArgs(anyString{}).WillReturnRows(rows)

	ramos, err := floresta.GetLema("some lema")

	if err == nil {
		t.Error("Expecting an error")
	}

	if len(ramos) != 2 {
		t.Errorf("len(ramos) = %d, expected: %d", len(ramos), 2)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_GetLema_ScanError(t *testing.T) {
	floresta, mock := newMockFloresta()

	rows := sqlmock.NewRows([]string{
		"id",
		"arvore",
		"palavra",
		"lema",
		"funcao",
		"froma",
		"morfo",
	}).AddRow(
		42,
		420,
		"palavra",
		"lema",
		"funcao",
		"forma",
		"morfo",
	)

	mock.ExpectQuery(getLemaQuery).WithArgs(anyString{}).WillReturnRows(rows)

	ramos, err := floresta.GetLema("some word")

	if err == nil {
		t.Error("Expecting an error")
	}

	if len(ramos) != 0 {
		t.Errorf("len(ramos) = %d, expected: %d", len(ramos), 2)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_ReadArvores(t *testing.T) {
	floresta, mock := newMockFloresta()

	rows := sqlmock.NewRows([]string{
		"id",
		"referencia",
		"n",
		"cad",
		"sec",
		"sem",
		"texto",
		"analise",
	}).AddRow(
		42,
		"referencia",
		10,
		"cad",
		"sec",
		"sem",
		"texto",
		"analise",
	)

	mock.ExpectQuery(readArvoresQuery).WithArgs(1).WillReturnRows(rows)

	_, err := floresta.ReadArvores(1)

	if err != nil {
		t.Errorf("Did not expect an error: %q", err.Error())
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_ReadArvores_QueryError(t *testing.T) {
	floresta, mock := newMockFloresta()

	mock.ExpectQuery(readArvoresQuery).WithArgs(1).WillReturnError(errors.New("some error"))

	_, err := floresta.ReadArvores(1)

	if err == nil {
		t.Error("Expecting an error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_ReadArvores_RowsError(t *testing.T) {
	floresta, mock := newMockFloresta()

	rows := sqlmock.NewRows([]string{
		"id",
		"referencia",
		"n",
		"cad",
		"sec",
		"sem",
		"texto",
		"analise",
	})

	for i := 0; i < 5; i++ {
		rows.AddRow(
			42+i,
			"referencia",
			10,
			"cad",
			"sec",
			"sem",
			"texto",
			"analise")
	}

	rows.RowError(3, errors.New("Some row error"))

	mock.ExpectQuery(readArvoresQuery).WithArgs(1).WillReturnRows(rows)

	arvores, err := floresta.ReadArvores(1)

	if err == nil {
		t.Error("Expecting an error")
	}

	if len(arvores) != 3 {
		t.Errorf("len(arvores) = %s, expecting: %d", len(arvores), 3)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

func Test_ReadArvores_ScanError(t *testing.T) {
	floresta, mock := newMockFloresta()

	rows := sqlmock.NewRows([]string{
		"id",
		"referencia",
		"n",
		"cad",
		"sec",
		"sem",
	})

	for i := 0; i < 5; i++ {
		rows.AddRow(
			42+i,
			"referencia",
			10,
			"cad",
			"sec",
			"sem",
		)
	}

	mock.ExpectQuery(readArvoresQuery).WithArgs(1).WillReturnRows(rows)

	arvores, err := floresta.ReadArvores(1)

	if err == nil {
		t.Error("Expecting an error")
	}

	if len(arvores) != 0 {
		t.Errorf("len(arvores) = %s, expecting: %d", len(arvores), 0)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}
}

var getWordQuery = escapeQuery("SELECT id, arvore, palavra, lema, funcao, forma, morfo, pai FROM ramos WHERE palavra = $1")

var getLemaQuery = escapeQuery("SELECT id, arvore, palavra, lema, funcao, forma, morfo, pai FROM ramos WHERE lema = $1")

var readArvoresQuery = escapeQuery("SELECT id, referencia, n, cad, sec, sem, texto, analise FROM arvores LIMIT $1")

type anyString struct{}

func (a anyString) Match(v driver.Value) bool {
	_, ok := v.(string)
	return ok
}

func escapeQuery(s string) string {
	re := regexp.MustCompile("\\s+")

	s1 := regexp.QuoteMeta(s)
	s1 = strings.TrimSpace(re.ReplaceAllString(s1, " "))
	return s1
}
