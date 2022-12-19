import type { DateRange } from "@react-types/datepicker";

interface Err {
  message: string;
  error: string;
}

export enum QueryType {
  Keyword = "keyword",
  User = "user",
}

export interface Query {
  type: QueryType;
  query: string;
  timeRange: DateRange;
}

export const searchURL = (
  url: string,
  { type, query, timeRange }: Query
): string => {
  if (!type || !query) return url;

  let base = `${url}?type=${type}&query=${encodeURIComponent(
    query
  )}&amount=100`;
  if (timeRange) {
    const start = timeRange.start
      .subtract({ hours: 1 })
      .toDate("utc")
      .toISOString();
    const end = timeRange.end
      .subtract({ hours: 1 })
      .toDate("utc")
      .toISOString();
    base += `&startTime=${start}&endTime=${end}`;
  }
  return base;
};

const f = async <T>(api: string, options?: RequestInit): Promise<T> => {
  const req = await fetch(
    `${import.meta.env.DEV ? "http://localhost:8080" : ""}/api/${api}`,
    options
  );
  let json = await req.json();
  if (req.status != 200) {
    json = json as Err;
    console.group("Error from API: " + json.message);
    console.trace(req.status, req.url);
    console.error(json.error);
    console.groupEnd();
    throw new Error(json.message);
  }

  return json as T;
};

export const withJSON = (method: string, obj: Object): RequestInit => ({
  method,
  body: JSON.stringify(obj),
  headers: {
    "Content-Type": "application/json",
  },
});

export default f;
