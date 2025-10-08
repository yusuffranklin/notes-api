package test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/yusuffranklin/notes-api/database"
	"github.com/yusuffranklin/notes-api/handlers"
	"github.com/yusuffranklin/notes-api/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateNoteHandler(t *testing.T) {
	tests := []struct {
		name           string
		payload        string
		mockSetup      func(sqlmock.Sqlmock)
		expectedStatus int
		expectedNote   *models.Note
	}{
		{
			name:    "success",
			payload: `{"title":"test-title","content":"test-content"}`,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO notes").
					WithArgs("test-title", "test-content").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			expectedStatus: http.StatusCreated,
			expectedNote:   &models.Note{ID: 1, Title: "test-title", Content: "test-content"},
		},
		{
			name:    "invalid JSON",
			payload: `{invalid-json`,
			mockSetup: func(mock sqlmock.Sqlmock) {
				// no DB call expected
			},
			expectedStatus: http.StatusBadRequest,
			expectedNote:   nil,
		},
		{
			name:    "DB error",
			payload: `{"title":"fail-title","content":"fail-content"}`,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO notes").
					WithArgs("fail-title", "fail-content").
					WillReturnError(sql.ErrConnDone)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedNote:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Setup sqlmock
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()
			database.Db = db

			if test.mockSetup != nil {
				test.mockSetup(mock)
			}

			req := httptest.NewRequest(http.MethodPost, "/notes", bytes.NewBufferString(test.payload))
			w := httptest.NewRecorder()

			handlers.CreateNoteHandler(w, req)

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, test.expectedStatus, res.StatusCode)

			if test.expectedNote != nil {
				var note models.Note
				err := json.NewDecoder(res.Body).Decode(&note)
				assert.NoError(t, err)
				assert.Equal(t, *test.expectedNote, note)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetNoteHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	database.Db = db

	rows := sqlmock.NewRows([]string{"id", "title", "content"}).AddRow(1, "title-1", "content-1")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, "/notes/1", nil)
	w := httptest.NewRecorder()

	handlers.GetNoteHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, res.StatusCode)
}
