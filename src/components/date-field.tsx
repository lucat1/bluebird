import * as React from "react";
import { useLocale } from "@react-aria/i18n";
import { DateFieldState, DateFieldStateOptions, DateSegment, useDateFieldState } from "@react-stately/datepicker";
import { AriaDatePickerProps, useDateField, useDateSegment } from "@react-aria/datepicker";
import { createCalendar } from "@internationalized/date";
import { DateValue } from "@react-types/calendar";

const DateSegmentView: React.FC<{ segment: DateSegment, state: DateFieldState }> = ({ segment, state }) => {
  let ref = React.useRef();
  let { segmentProps } = useDateSegment(segment, state, ref as any);

  return (
    <div
      {...segmentProps}
      ref={ref as any}
      style={{
        ...segmentProps.style,
        minWidth:
          segment.maxValue && String(segment.maxValue).length + "ch"
      }}
      className={`px-0.5 box-content tabular-nums text-right outline-none rounded-sm focus:bg-violet-600 focus:text-white group ${!segment.isEditable ? "text-gray-500" : "text-gray-800"
        }`}
    >
      {/* Always reserve space for the placeholder, to prevent layout shift when editing. */}
      <span
        aria-hidden="true"
        className="block w-full text-center italic text-gray-500 group-focus:text-white"
        style={{
          visibility: segment.isPlaceholder ? undefined : "hidden",
          height: segment.isPlaceholder ? "" : 0,
          pointerEvents: "none"
        }}
      >
        {segment.placeholder}
      </span>
      {segment.isPlaceholder ? "" : segment.text}
    </div>
  );
}

const DateField: React.FC<AriaDatePickerProps<DateValue>> = (props) => {
  let { locale } = useLocale();
  let state = useDateFieldState({
    ...props,
    locale,
    createCalendar
  });

  let ref = React.useRef();
  let { fieldProps } = useDateField(props, state, ref as any);

  return (
    <div {...fieldProps} ref={ref as any} className="flex">
      {state.segments.map((segment, i) => (
        <DateSegmentView key={i} segment={segment} state={state} />
      ))}
    </div>
  );
}


export default DateField
