<script setup lang="ts">
import TextArea from '@/components/TextArea.vue';
import { ref } from 'vue';
import TextInput from '@/components/TextInput.vue';
import FileInput from '@/components/FileInput.vue';
import axios from 'axios';
import type { NewPost } from '@/datatypes/PostData';

const newPostTitle = ref("");
const newPostDescription = ref("");
const textAreaPostDetail = "Post Detail"
let uploadedFile: Blob;

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
  if (!isPostValid(newPostTitle.value, newPostDescription.value)) {
    console.error("Invalid Post");
    return
  }
  const newPost: NewPost = { title: newPostTitle.value, description: newPostDescription.value, imageUrl: null };
  PostBlogPost(newPost, uploadedFile);
};

function PostBlogPost(newPost: NewPost, imageFile: Blob) {
  const formData = new FormData()
  formData.append("postTitle", newPost.title)
  formData.append("postDescription", newPost.description)
  formData.append("imageFile", imageFile)
  axios.post("/api/blogposts", formData)
}


function onFileChange(e: any) {
  var files = e.target.files || e.dataTransfer.files;
  if (!files.length)
    return;
  uploadedFile = files[0]
  console.log(uploadedFile)
}
const acceptedFileTypes = ["image/png", "image/jpeg"].join(",")
</script>

<template>
  <div class="form-wrapper">
    <form @submit.prevent="addPost">
      <TextInput :placeholder='"Test"' v-model="newPostTitle" />
      <br>
      <TextArea :placeholder="textAreaPostDetail" v-model="newPostDescription" @keyup.enter="addPost" />
      <br>
      <FileInput id="file" :acceptedFileTypes="acceptedFileTypes" v-on:fileChange="onFileChange" />
      <br>
      <button type="submit">Add Post</button>
    </form>
  </div>
</template>

<style scoped>
.form-wrapper {
  margin: 5px 25% 0px 25%;
}
</style>
