import { createApp } from 'vue'
import { createMemoryHistory, createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import PostCollection from './views/PostCollection.vue'
import NewPost from './views/NewPost.vue'
import AboutPage from './views/AboutPage.vue'
import FixturesPage from './views/FixturesPage.vue'
import ContactPage from './views/ContactPage.vue'

const routes = [
  { path: '/', component: PostCollection },
  { path: '/Post', component: NewPost },
  { path: '/About', component: AboutPage },
  { path: '/Fixtures', component: FixturesPage },
  { path: '/Contact', component: ContactPage },
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

createApp(App).use(router).mount('#app')
