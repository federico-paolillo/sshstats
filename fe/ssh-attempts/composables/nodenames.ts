export function useNodenames(): string[] {
  const cfg = useRuntimeConfig()

  return cfg.public.nodenames ?? []
}