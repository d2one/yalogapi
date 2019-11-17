package yalogapi

import (
	"fmt"
)

type YaLogApi struct {
	config *Config
}

func main() {

}

func NewYaLogApi(config *Config) *YaLogApi {
	fmt.Println(config)
	return &YaLogApi{config: config}
}

func (yalogapi *YaLogApi) Run() {
	fmt.Println(yalogapi.config)
	fmt.Println("YaLogApi")
}

func (yalogapi *YaLogApi) buildUderRequest() {

}

// def build_user_request(config):
//     options = utils.get_cli_options()
//     logger.info('CLI Options: ' + str(options))

//     start_date_str, end_date_str = get_date_period(options)
//     source = options.source

//     # Validate that fields are present in config
//     assert '{source}_fields'.format(source=source) in config, \
//         'Fields must be specified in config'
//     fields = config['{source}_fields'.format(source=source)]

//     # Creating data structure (immutable tuple) with initial user request
//     UserRequest = namedtuple(
//         "UserRequest",
//         "token counter_id start_date_str end_date_str source fields"
//     )

//     user_request = UserRequest(
//         token=config['token'],
//         counter_id=config['counter_id'],
//         start_date_str=start_date_str,
//         end_date_str=end_date_str,
//         source=source,
//         fields=tuple(fields),
//     )

//     logger.info(user_request)
//     utils.validate_user_request(user_request)
//     return user_request

// def get_date_period(options):
//     if options.mode is None:
//         start_date_str = options.start_date
//         end_date_str = options.end_date
//     else:
//         if options.mode == 'regular':
//             start_date_str = (datetime.datetime.today() - datetime.timedelta(2)) \
//                 .strftime(utils.DATE_FORMAT)
//             end_date_str = (datetime.datetime.today() - datetime.timedelta(2)) \
//                 .strftime(utils.DATE_FORMAT)
//         elif options.mode == 'regular_early':
//             start_date_str = (datetime.datetime.today() - datetime.timedelta(1)) \
//                 .strftime(utils.DATE_FORMAT)
//             end_date_str = (datetime.datetime.today() - datetime.timedelta(1)) \
//                 .strftime(utils.DATE_FORMAT)
//         elif options.mode == 'history':
//             start_date_str = utils.get_counter_creation_date(
//                 config['counter_id'],
//                 config['token']
//             )
//             end_date_str = (datetime.datetime.today() - datetime.timedelta(2)) \
//                 .strftime(utils.DATE_FORMAT)
//     return start_date_str, end_date_str
