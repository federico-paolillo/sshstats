import type { LoginAttempt } from "~/models/attempts"
import { useEndpoints } from "./endpoints"

export async function useLoginAttempts(nodename: string) {
  const cfg = useRuntimeConfig()

  const { endpoint } = useEndpoints()

  const { data, error } = await useFetch<LoginAttempt>(
    endpoint(`attempts/${nodename}`),
    {
      method: 'GET',
      headers: {
        [cfg.apiAuthHeaderKey]: cfg.apiAuthHeaderValue
      }
    })

  return { data, error }
}