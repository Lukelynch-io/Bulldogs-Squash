import { createApp } from 'vue'
import { createMemoryHistory, createRouter } from 'vue-router'
import App from './App.vue'
import BlogPostCollection from './views/BlogPostCollection.vue'
import NewBlog from './views/NewBlog.vue'

const routes = [
  { path: '/', component: BlogPostCollection },
  { path: '/Post', component: NewBlog }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes
})

createApp(App)
  .use(router)
  .mount('#app')
