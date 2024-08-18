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
    apiAuthHeaderKey: "<specify_in_.env>",
    apiAuthHeaderValue: "<specify_in_.env>",
    apiEndpoint: "<specify_in_.env>",
    public: {
      nodenames: [
        'controlplane-1'
      ]
    }
  }
})
