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
      nodenames: [
        'controlplane-1',
        'missingnode-1'
      ],
      baseUrl: "<specify_in_env>"
    }
  }
})
