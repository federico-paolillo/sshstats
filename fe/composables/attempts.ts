import type { LoginAttempts } from "~/models/attempts"
import { useEndpoints } from "./endpoints"

export async function useLoginAttempts(nodename: string) {
  const { endpoint } = useEndpoints()

  const { data, error } = await useFetch<LoginAttempts>(
    endpoint(`data/attempts/${nodename}.json`)
  )

  return { data, error }
}