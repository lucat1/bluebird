import * as React from "react";
import { AriaButtonProps, useButton } from "@react-aria/button";
import { useFocusRing } from "@react-aria/focus";
import { mergeProps } from "@react-aria/utils";

export const CalendarButton: React.FC<AriaButtonProps<"button">> = (props) => {
  let ref = React.useRef<HTMLButtonElement>();
  let { buttonProps } = useButton(props, ref as any);
  let { focusProps, isFocusVisible } = useFocusRing();

  return (
    <button
      {...mergeProps(buttonProps, focusProps)}
      ref={ref as any}
      className={`p-2 rounded-full ${props.isDisabled ? "text-gray-400" : ""} ${
        !props.isDisabled ? "hover:bg-sky-100 active:bg-sky-200" : ""
      } outline-none ${
        isFocusVisible ? "ring-2 ring-offset-2 ring-sky-600" : ""
      }`}
    >
      {props.children}
    </button>
  );
};

export const FieldButton: React.FC<
  AriaButtonProps<"button"> & { isPressed: boolean }
> = (props) => {
  let ref = React.useRef<HTMLButtonElement>();
  let { buttonProps, isPressed } = useButton(props, ref as any);

  return (
    <button
      {...buttonProps}
      ref={ref as any}
      className={`px-2 w-10  -ml-px border transition-colors rounded-r-md group-focus-within:border-sky-600 group-focus-within:group-hover:border-sky-600 outline-none ${
        isPressed || props.isPressed
          ? "bg-gray-200 border-gray-400"
          : "bg-gray-50 dark:bg-gray-700 dark:border-gray-600 border-gray-300 group-hover:border-gray-400"
      }`}
    >
      {props.children}
    </button>
  );
};
