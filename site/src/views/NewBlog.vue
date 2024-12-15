<script setup lang="ts">
import type BlogPostData from '@/datatypes/BlogPost';
import { ref } from 'vue';

const newPostTitle = ref<string>("");
const newPostDescription = ref<string>("");
function PostBlogPost(test: BlogPostData) {
  fetch("/api/blogposts", {
    method: "POST",
    body: JSON.stringify(test)
  })
}

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
    const newPost: BlogPostData = { title: newPostTitle.value, description: newPostDescription.value, imageUrl: "" };
    PostBlogPost(newPost);
    return;
  }
  console.error("Invalid Post");
};
</script>


<template>
  <div>
    <input type="text" v-model="newPostTitle" placeholder="Enter a blog title" />
    <br>
    <textarea rows="5" cols="50" v-model="newPostDescription" placeholder="Enter your blog description"
      @keyup.enter="addPost" />
    <br>
    <button @click="addPost">Add Post</button>
  </div>
</template>
