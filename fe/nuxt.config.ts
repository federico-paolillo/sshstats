export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  app: {
    head: {
      title: "sshstats",
      link: [
        {
          rel: 'icon',
          type: 'image/png',
          href: 'favicon.png'
        }
      ]
    }
  },
  runtimeConfig: {
    public: {
      baseUrl: "<specify_in_env>"
    }
  }
})
