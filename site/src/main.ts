import { createApp } from 'vue'
import { createMemoryHistory, createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import BlogPostCollection from './views/BlogPostCollection.vue'
import NewBlog from './views/NewBlog.vue'
import AboutPage from './views/AboutPage.vue'
import FixturesPage from './views/FixturesPage.vue'
import ContactPage from './views/ContactPage.vue'

const routes = [
  { path: '/', component: BlogPostCollection },
  { path: '/Post', component: NewBlog },
  { path: '/About', component: AboutPage },
  { path: '/Fixtures', component: FixturesPage },
  { path: '/Contact', component: ContactPage },
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

createApp(App).use(router).mount('#app')
