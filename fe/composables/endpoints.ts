export function useEndpoints() {
  const cfg = useRuntimeConfig()

  function endpoint(path: string): string {
    return `${cfg.public.baseUrl}/${path}`;
  }

  return { endpoint }
}