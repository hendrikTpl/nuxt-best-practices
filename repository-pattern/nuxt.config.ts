// https://nuxt.com/docs/api/configuration/nuxt-config
import vuetify, { transformAssetUrls } from 'vite-plugin-vuetify'
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  app:{
    head: {
      title:"DeepX.ID Dev",
      meta: [
        { 
          name: 'viewport', 
          content: 'width=device-width, initial-scale=1' 
        }
      ],
      link: [
        { rel: 'SHORTCUT ICON', type: 'image/png', href: '/favicon.png' }
      ],
      script: [],      
      noscript: []
    }
  },
  alias:{
    assets:"<rootDir>/assets"
  },
  css:[
    '~/assets/css/main.css',
  ],
  build: {
    transpile: ['vuetify'],
  },
  modules: [
    (_options, nuxt) => {
      nuxt.hooks.hook('vite:extendConfig', (config) => {
        // @ts-expect-error
        config.plugins.push(vuetify({ autoImport: true }))
      })
    },
    //...
  ],
  vite: {
    vue: {
      template: {
        transformAssetUrls,
      },
    },
  },
  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.API_BASE_URI || 'http://localhost:5000',
    }
  }
})
