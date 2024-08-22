import nodenames from "~/assets/nodenames.json"

export function useNodenames(): string[] {
  return nodenames ?? [];
}