package menu

import (
	"app/src/model"
	"app/src/model/constant"
	"context"
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_menuRepo_GetMenuList(t *testing.T) {
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name     string
		args     args
		want     []model.MenuItem
		wantErr  bool
		initMock func() (*sql.DB, sqlmock.Sqlmock, error)
	}{
		// TODO: Add test cases.
		{
			name: "success list menu",
			args: args{
				ctx:      context.Background(),
				menuType: "",
			},
			initMock: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "menu_items"`),
				).WillReturnRows(sqlmock.NewRows([]string{
					"name",
					"order_code",
					"price",
					"type",
				}).AddRow("Nasi Uduk", "nasi_uduk", 38000, constant.MenuTypeFood))
				return db, mock, err
			},
			want: []model.MenuItem{
				{
					Name:      "Nasi Uduk",
					OrderCode: "nasi_uduk",
					Price:     38000,
					Type:      constant.MenuTypeFood,
				},
			},
		},
		{
			name: "success list menu empty data",
			args: args{
				ctx:      context.Background(),
				menuType: "",
			},
			initMock: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "menu_items"`),
				).WillReturnRows(sqlmock.NewRows([]string{
					"name",
					"order_code",
					"price",
					"type",
				}))
				return db, mock, err
			},
			want: []model.MenuItem{},
		},
		{
			name: "fail list menu",
			args: args{
				ctx:      context.Background(),
				menuType: "",
			},
			initMock: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()
				mock.ExpectQuery(
					regexp.QuoteMeta(`SELECT * FROM "menu_items"`),
				).WillReturnError(errors.New("mock error"))
				return db, mock, err
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, err := tt.initMock()
			if err != nil {
				t.Error(err)
			}

			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{
				DSN:                  "sqlmock_db_0",
				DriverName:           "postgres",
				Conn:                 db,
				PreferSimpleProtocol: true,
			}))

			if err != nil {
				t.Error(err)
			}

			m := &menuRepo{
				db: gormDB,
			}

			got, err := m.GetMenuList(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuRepo.GetMenuList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuRepo.GetMenuList() = %v, want %v", got, tt.want)
			}

			if err := dbMock.ExpectationsWereMet(); err != nil {
				t.Errorf("expectations were not met: %s", err.Error())
			}
		})
	}
}
