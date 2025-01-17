<script setup lang="ts">
import { onMounted, onUnmounted, ref } from 'vue';
import type PostData from '@/datatypes/PostData';
import Post from '../components/Post.vue';
import AlertBox from '@/components/AlertBox.vue';
import { MessageType } from '@/datatypes/MessageType';
import { GetBlogPosts } from '@/api_calls';

const postsArray = ref<PostData[]>([]); // Reactive array for blog posts
const errorOccurred = ref(false);
const errorMessage = ref('')
const errorTitle = ref('Error Title')
let postLoadInterval = 0;
async function RefreshBlogPosts() {
  // Fetch blog posts from the API
  let [returnArray, returnError] = await GetBlogPosts();
  if (returnError != "") {
    errorOccurred.value = true;
    errorMessage.value = returnError
    return;
  }
  errorOccurred.value = false;
  postsArray.value = returnArray;
}

onMounted(() => {
  RefreshBlogPosts();
  postLoadInterval = setInterval(RefreshBlogPosts, 10000);
})

onUnmounted(() => {
  clearInterval(postLoadInterval)
})

</script>

<template>
  <div class="card-collection">
    <Post v-for="(post, index) in postsArray" :key=index v-bind="post" />
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
