interface Err {
  message: string
  error: string
}

const f = async <T>(api: string): Promise<T> => {
  const req = await fetch(`/api/${api}`)
  let json = await req.json()
  if (req.status != 200) {
    json = json as Err
    console.group("Error from API: " + json.message)
    console.trace(req.status, req.url)
    console.error(json.error)
    console.groupEnd()
    throw new Error(json.message)
  }
  return json as T
}

export default f
