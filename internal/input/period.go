package prompts

//func ParsePeriod(period string, startTime string, endTime string) (time.Time, time.Time, error) {
//	if period != "" { // if period is indicated
//		if startTime != "" || endTime != "" {
//			// return error if --start-time or --end-time are indicated
//			return ctx.Command.OnUsageError(
//				ctx,
//				errors.New("PERIOD cannot be indicated if you use --start-time or --end-time options"),
//				true,
//			)
//		}
//
//		var err error
//		startTime, endTime, err = dates.ParsePeriod(period)
//		if err != nil {
//			return err
//		}
//
//	} else {
//		if startTimeString == "" {
//			// return error if --start-time is not indicated
//			return ctx.Command.OnUsageError(
//				ctx,
//				errors.New("if PERIOD is not indicated, at least --start-time option is required"),
//				true,
//			)
//		}
//
//		var err error
//		startTime, err = dates.ParseDate(startTimeString)
//		if err != nil {
//			return err
//		}
//
//		if endTimeString == "" { // if --end-time is not indicated, use current time as the default value
//			endTime = time.Now()
//		}
//	}
//}
