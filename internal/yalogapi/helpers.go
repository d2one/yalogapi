package yalogapi

import (
	"time"
)

func getDatePeriod(conf *Config) (string, string) {
	if conf.Mode == "" {
		return conf.StartDate, conf.EndDate
	}
	switch conf.Mode {
	case "regular":
		return conf.StartDate, conf.EndDate
	case "regular_early":
		return conf.StartDate, conf.EndDate
	case "history":
		return conf.StartDate, conf.EndDate
	case "":
		return conf.StartDate, conf.EndDate
	}

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
	return "", ""
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

// @TODO don`t validate date here
func getNewDates(from string, to string, daysInPeriod int, partNumber int) (string, string) {
	dateFrom, _ := time.Parse("2006-01-02", from)
	newDateFrom := dateFrom.AddDate(0, 0, partNumber*daysInPeriod)

	dateTo, _ := time.Parse("2006-01-02", to)
	newDateTo := dateFrom.AddDate(0, 0, (partNumber+1)*daysInPeriod-1)

	if newDateTo.After(dateTo) {
		newDateTo = dateTo
	}

	return newDateFrom.Format("2006-01-02"), newDateTo.Format("2006-01-02")
}
