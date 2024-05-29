// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  pages: true,
  css: ['~/assets/css/main.css'],
  postcss:{
    plugins:{
      tailwindcss:{},
      autoprefixer:{},
    },
  },
  // devtools: { enabled: true },
  modules: [
    'nuxt-icon',
    '@pinia/nuxt',
    '@pinia-plugin-persistedstate/nuxt'
  ]
})
