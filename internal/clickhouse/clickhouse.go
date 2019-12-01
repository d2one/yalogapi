package clickhouse

import (
	"fmt"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

type Clickhouse struct {
	config  *Config
	connect *sqlx.DB
}

func NewClickhouse(config *Config) (*Clickhouse, error) {

	clickhouse := &Clickhouse{config: config}
	connect, err := clickhouse.conn()
	if err != nil {
		return nil, err
	}
	clickhouse.connect = connect
	return clickhouse, nil
}

func (clickhouse *Clickhouse) conn() (*sqlx.DB, error) {
	return sqlx.Open("clickhouse", clickhouse.config.getDSN())
}

func (clickhouse *Clickhouse) IsDataPresent(startDate string, endDate string, source string) (bool, error) {
	// connect, err := clickhouse.conn()
	// if err != nil {
	// 	return false, err
	// }
	// connect.Select("SEE")
	return true, nil
}

// if not is_db_present():
// logger.info('Database created')
// create_db()

// if not is_table_present(source):
// logger.info('Table created')
// create_table(source, fields)

func (clickhouse *Clickhouse) Check(source string, fields map[string]string) {
	table := clickhouse.getTable(source)
	rows, err := clickhouse.Query("SHOW DATABASES")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		fmt.Println(rows)
	}

	clickhouse.createTable(table, fields)
}

func (clickhouse *Clickhouse) Query(query string) (*sqlx.Rows, error) {
	return clickhouse.connect.Queryx(query)
}

// def get_clickhouse_data(query, host=CH_HOST):
//     '''Returns ClickHouse response'''
//     logger.debug(query)
//     if (CH_USER == '') and (CH_PASSWORD == ''):
//         r = requests.post(host, data=query, verify=SSL_VERIFY)
//     else:
//         r = requests.post(host, data=query, auth=(CH_USER, CH_PASSWORD), verify=SSL_VERIFY)
//     if r.status_code == 200:
//         return r.text
//     else:
//         raise ValueError(r.text)

func (clickhouse *Clickhouse) SaveData(source string, fields []string) {
	// clickhouse.createTable(source, fields)
}

func (clickhouse *Clickhouse) createTable(table string, fields map[string]string) {

	// tmpl := `CREATE TABLE :table_name (:fields) ENGINE = :engine`
	// table := clickhouse.getTable(source)

	//     '''Creates table in ClickHouse for hits/visits with particular fields'''
	//     tmpl = '''
	//         CREATE TABLE {table_name} (
	//             {fields}
	//         ) ENGINE = {engine}
	//     '''
	//     field_tmpl = '{name} {type}'
	//     field_statements = []

	//     table_name = get_source_table_name(source)
	//     if source == 'hits':
	//         if ('ym:pv:date' in fields) and ('ym:pv:clientID' in fields):
	//             engine = 'MergeTree(Date, intHash32(ClientID), (Date, intHash32(ClientID)), 8192)'
	//         else:
	//             engine = 'Log'

	//     if source == 'visits':
	//         if ('ym:s:date' in fields) and ('ym:s:clientID' in fields):
	//             engine = 'MergeTree(Date, intHash32(ClientID), (Date, intHash32(ClientID)), 8192)'
	//         else:
	//             engine = 'Log'

	//     ch_field_types = utils.get_ch_fields_config()
	//     ch_fields = map(get_ch_field_name, fields)

	//     for i in range(len(fields)):
	//         field_statements.append(field_tmpl.format(name= ch_fields[i],
	//             type=ch_field_types[fields[i]]))

	//     field_statements = sorted(field_statements)
	//     query = tmpl.format(table_name=table_name,
	//                         engine=engine,
	//                         fields=',\n'.join(sorted(field_statements)))

	//     get_clickhouse_data(query)
	// get mapping

}

func (clickhouse *Clickhouse) getTable(source string) string {
	fmt.Println(clickhouse.config)
	switch source {
	case "hits":
		return clickhouse.config.HitsTable

	case "visits":
		return clickhouse.config.VisitsTable
	}
	return ""
}

// def create_table(source, fields):
//     '''Creates table in ClickHouse for hits/visits with particular fields'''
//     tmpl = '''
//         CREATE TABLE {table_name} (
//             {fields}
//         ) ENGINE = {engine}
//     '''
//     field_tmpl = '{name} {type}'
//     field_statements = []

//     table_name = get_source_table_name(source)
//     if source == 'hits':
//         if ('ym:pv:date' in fields) and ('ym:pv:clientID' in fields):
//             engine = 'MergeTree(Date, intHash32(ClientID), (Date, intHash32(ClientID)), 8192)'
//         else:
//             engine = 'Log'

//     if source == 'visits':
//         if ('ym:s:date' in fields) and ('ym:s:clientID' in fields):
//             engine = 'MergeTree(Date, intHash32(ClientID), (Date, intHash32(ClientID)), 8192)'
//         else:
//             engine = 'Log'

//     ch_field_types = utils.get_ch_fields_config()
//     ch_fields = map(get_ch_field_name, fields)

//     for i in range(len(fields)):
//         field_statements.append(field_tmpl.format(name= ch_fields[i],
//             type=ch_field_types[fields[i]]))

//     field_statements = sorted(field_statements)
//     query = tmpl.format(table_name=table_name,
//                         engine=engine,
//                         fields=',\n'.join(sorted(field_statements)))

//     get_clickhouse_data(query)

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
