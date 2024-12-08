import { createApp } from 'vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import App from './App.vue'
import BlogPostCollection from './components/BlogPostCollection.vue'
import LandingPage from './components/LandingPage.vue'

const routes = [
  { path: '/', component: BlogPostCollection },
  { path: '/Login', component: LandingPage }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

createApp(App)
  .use(router)
  .mount('#app')
