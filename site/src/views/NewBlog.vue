<script setup lang="ts">
import type BlogPostData from '@/datatypes/BlogPost';
import TextArea from '@/components/TextArea.vue';
import { reactive, ref } from 'vue';
import TextInput from '@/components/TextInput.vue';
import FileInput from '@/components/FileInput.vue';

const newPostTitle = ref<string>("");
const newPostDescription = ref<string>("");
const textAreaPostDetail = "Post Detail"
let uploadedFile = reactive(File)
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
  <div>
    <form @submit.prevent="addPost">
      <TextInput :placeholder='"Test"' />
      <br>
      <TextArea :placeholder="textAreaPostDetail" @keyup.enter="addPost" />
      <br>
      <FileInput id="file" :acceptedFileTypes="acceptedFileTypes" v-on:fileChange="onFileChange" />
      <br>
      <button type="submit">Add Post</button>
    </form>
  </div>
</template>
