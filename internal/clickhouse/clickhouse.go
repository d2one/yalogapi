package clickhouse

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type clickhouse struct {
	config  *Config
	connect *sqlx.DB
}

func NewClickhouse() *clickhouse {
	return &clickhouse{}
}

func (clickhouse *clickhouse) conn() (*sqlx.DB, error) {
	return sqlx.Open("clickhouse", clickhouse.config.getDSN())
}

func (clickhouse *clickhouse) IsDataPresent(startDate string, endDate string, source string) (bool, error) {
	// connect, err := clickhouse.conn()
	// if err != nil {
	// 	return false, err
	// }
	// connect.Select("SEE")
	return true, nil
}

func main() {
	fmt.Println("clickhouse package")
}

func Run() {
	fmt.Println("clickhouse package")
}

func connect() {

}

// def get_clickhouse_data(query, host=CH_HOST):
// def upload(table, content, host=CH_HOST):
// def get_source_table_name(source, with_db=True):
// def get_tables():
// def get_dbs():
// def is_table_present(source):
// def is_db_present():
// def create_db():
// def get_ch_field_name(field_name):
// def drop_table(source):
// def create_table(source, fields):
// def save_data(source, fields, data):
// def is_data_present(start_date_str, end_date_str, source):
