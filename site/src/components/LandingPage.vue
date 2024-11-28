<script setup lang="ts">
import BlogPost from './BlogPost.vue'
import { ref } from 'vue';

interface Post {
  title: string;
  description: string;
}

const items = ref<Post[]>(
  [
    {
      title: "test1",
      description: "test2"
    },
    {
      title: "test3",
      description: "test4"
    }
  ]
);
const newPostTitle = ref<string>("");
const newPostDescription = ref<string>("");

function isPostValid(postTitle: string, postDescription: string) {
  if (postTitle.trim() === "") {
    return false;
  }
  if (postDescription.trim() === "") {
    return false;
  }
  return true;
}

const addPost = () => {
  if (isPostValid(newPostTitle.value, newPostDescription.value)) {
    let newPost: Post = { title: newPostTitle.value, description: newPostDescription.value };
    items.value.push(newPost);
    newPostTitle.value = "";
    newPostDescription.value = "";
  }
};

const popPost = () => {
  if (items.value.length > 0) {
    items.value.pop();
  }
};
</script>

<template>
  <div class="wrapper">
    <h1>This is the landing page</h1>
    <input type="text" v-model="newPostTitle" placeholder="Enter a blog title" />
    <br>
    <textarea rows="5" cols="50" v-model="newPostDescription" placeholder="Enter your blog description"
      @keyup.enter="addPost" />
    <br>
    <button @click="addPost">Add Post</button>
  </div>
  <div>
    <BlogPost v-for="item in items">
      <template #blogTitle>{{ item.title }}</template>
      <template #blogDescription>{{ item.description }}</template>
    </BlogPost>
  </div>
  <button @click="popPost">Remove Last Word</button>
</template>


<style scoped>
.wrapper {
  display: grid;
  justify-content: center;
  align-items: center;
}
</style>
