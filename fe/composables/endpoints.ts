export function useEndpoints() {
  const cfg = useRuntimeConfig()

  function endpoint(path: string): string {
    return new URL(path, cfg.public.baseUrl).toString()
  }

  return { endpoint }
}