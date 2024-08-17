// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  runtimeConfig: {
    apiAuthHeaderKey: "<specify_in_.env>",
    apiAuthHeaderValue: "<specify_in_.env>"
  }
})
