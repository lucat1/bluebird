import { useRef } from "react";
import { useRangeCalendarState } from "@react-stately/calendar";
import { useRangeCalendar } from "@react-aria/calendar";
import { useLocale } from "@react-aria/i18n";
import { createCalendar } from "@internationalized/date";
import { CalendarButton } from "./button";
import { CalendarGrid } from "./calendar-grid";
import { ChevronLeftIcon, ChevronRightIcon } from "@heroicons/react/24/outline";

export function RangeCalendar(props) {
  let { locale } = useLocale();
  let state = useRangeCalendarState({
    ...props,
    locale,
    createCalendar
  });

  let ref = useRef();
  let {
    calendarProps,
    prevButtonProps,
    nextButtonProps,
    title
  } = useRangeCalendar(props, state, ref);

  return (
    <div {...calendarProps} ref={ref} className="inline-block">
      <div className="flex items-center pb-4">
        <h2 className="flex-1 font-bold text-xl ml-2">{title}</h2>
        <CalendarButton {...prevButtonProps}>
          <ChevronLeftIcon className="h-6 w-6" />
        </CalendarButton>
        <CalendarButton {...nextButtonProps}>
          <ChevronRightIcon className="h-6 w-6" />
        </CalendarButton>
      </div>
      <CalendarGrid state={state} />
    </div>
  );
}
