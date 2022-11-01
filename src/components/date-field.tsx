import { useRef } from "react";
import { useLocale } from "@react-aria/i18n";
import { useDateFieldState } from "@react-stately/datepicker";
import { useDateField, useDateSegment } from "@react-aria/datepicker";
import { createCalendar } from "@internationalized/date";

export function DateField(props) {
  let { locale } = useLocale();
  let state = useDateFieldState({
    ...props,
    locale,
    createCalendar
  });

  let ref = useRef();
  let { fieldProps } = useDateField(props, state, ref);

  return (
    <div {...fieldProps} ref={ref} className="flex">
      {state.segments.map((segment, i) => (
        <DateSegment key={i} segment={segment} state={state} />
      ))}
    </div>
  );
}

function DateSegment({ segment, state }) {
  let ref = useRef();
  let { segmentProps } = useDateSegment(segment, state, ref);

  return (
    <div
      {...segmentProps}
      ref={ref}
      style={{
        ...segmentProps.style,
        minWidth:
          segment.maxValue != null && String(segment.maxValue).length + "ch"
      }}
      className={`px-0.5 box-content tabular-nums text-right outline-none rounded-sm focus:bg-violet-600 focus:text-white group ${
        !segment.isEditable ? "text-gray-500" : "text-gray-800"
      }`}
    >
      {/* Always reserve space for the placeholder, to prevent layout shift when editing. */}
      <span
        aria-hidden="true"
        className="block w-full text-center italic text-gray-500 group-focus:text-white"
        style={{
          visibility: segment.isPlaceholder ? "" : "hidden",
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
