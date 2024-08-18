export function useEndpoints() {
  const cfg = useRuntimeConfig()

  function endpoint(path: string): string {
    return new URL(path, cfg.apiEndpoint).toString()
  }

  return { endpoint }
}