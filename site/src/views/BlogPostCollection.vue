<script setup lang="ts">
import { onMounted, ref } from 'vue';
import type BlogPostData from '@/datatypes/BlogPost';
import BlogPost from '../components/BlogPost.vue';
import AlertBox from '@/components/AlertBox.vue';
import { MessageType } from '@/datatypes/MessageType';

const blogPosts = ref<BlogPostData[]>([]); // Reactive array for blog posts
const errorOccurred = ref(false);
const errorMessage = ref('')
const errorTitle = ref('Error Title')

function RefreshBlogPosts() {
  // Fetch blog posts from the API
  fetch("/api/blogposts")
    .then((response) => response.json())
    .then((data: BlogPostData[]) => {
      errorOccurred.value = false;
      blogPosts.value = data; // Update the reactive array
    })
    .catch((error) => {
      errorOccurred.value = true;
      errorMessage.value = "Error fetching blog posts: " + error;
    });
}

onMounted(() => {
  RefreshBlogPosts();
  setInterval(RefreshBlogPosts, 10000);
})

</script>

<template>
  <div class="card-collection">
    <BlogPost v-for="(post, index) in blogPosts" :key=index v-bind="post" />
  </div>
  <AlertBox v-if="errorOccurred" :messageType="MessageType.Info" :messageTitle="errorTitle"
    :messageDescription="errorMessage" />

</template>

<style scoped>
.card-collection {
  padding-left: 5%;
  padding-right: 5%;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  grid-template-rows: 100% repeat(auto-fit, minmax(150px, 1fr));
  gap: 1rem;
  margin-top: 5rem;
}

@media only screen and (min-width: 1200px) {
  .card-collection {
    grid-template-columns: repeat(3, minmax(350px, 1fr));
  }
}
</style>
