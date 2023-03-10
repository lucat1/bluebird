import * as React from "react";
import {
  DateRangePickerStateOptions,
  useDateRangePickerState,
} from "@react-stately/datepicker";
import {
  AriaDateRangePickerProps,
  useDateRangePicker,
} from "@react-aria/datepicker";
import { FieldButton } from "./button";
import RangeCalendar from "./range-calendar";
import DateField from "./date-field";
import Popover from "./popover";
import {
  CalendarIcon,
  ExclamationTriangleIcon,
} from "@heroicons/react/24/outline";
import { DateValue } from "@react-types/calendar";

const DateRangePicker: React.FC<
  AriaDateRangePickerProps<DateValue> & DateRangePickerStateOptions
> = (props) => {
  let state = useDateRangePickerState(props);
  let ref = React.useRef<Element>();
  let {
    groupProps,
    labelProps,
    startFieldProps,
    endFieldProps,
    buttonProps,
    dialogProps,
    calendarProps,
  } = useDateRangePicker(props, state, ref as any);

  return (
    <div className="relative inline-flex flex-col text-left">
      <span {...labelProps} className="hidden">
        {props.label}
      </span>
      <div
        {...groupProps}
        ref={ref as any}
        className="flex  max-w-screen-sm group"
      >
        <div className="flex sm:flex-row flex-col bg-white dark:bg-gray-700 border dark:border-gray-600 border-gray-300 group-hover:border-gray-400 transition-colors rounded-l-md pr-1 group-focus-within:border-sky-500 group-focus-within:group-hover:border-sky-600 p-1 relative">
          <DateField {...startFieldProps} />
          <span aria-hidden="true" className="px-2">
            -
          </span>
          <DateField {...endFieldProps} />
          {state.validationState === "invalid" && (
            <ExclamationTriangleIcon className="w-6 h-6 text-red-500 absolute right-1" />
          )}
        </div>
        <FieldButton {...buttonProps} isPressed={state.isOpen}>
          <CalendarIcon className="w-5 h-5 text-gray-700 group-focus-within:text-sky-700 dark:text-white" />
        </FieldButton>
      </div>
      {state.isOpen && (
        <Popover
          {...dialogProps}
          isOpen={state.isOpen}
          onClose={() => state.setOpen(false)}
        >
          <RangeCalendar {...calendarProps} />
        </Popover>
      )}
    </div>
  );
};

export default DateRangePicker;
