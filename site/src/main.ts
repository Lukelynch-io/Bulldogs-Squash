import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import BlogPostCollection from './views/BlogPostCollection.vue'
import NewBlog from './views/NewBlog.vue'
import AboutPage from './views/AboutPage.vue'

const routes = [
  { path: '/', component: BlogPostCollection },
  { path: '/Post', component: NewBlog },
  { path: '/About', component: AboutPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

createApp(App).use(router).mount('#app')
