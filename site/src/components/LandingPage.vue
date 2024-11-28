<script setup lang="ts">
import BlogPost from './BlogPost.vue'
import { ref } from 'vue';

const items = ref<object[]>([{
  title: "test1",
  description: "test2"
},
{
  title: "test3",
  description: "test4"
}]);
const newWord = ref<string>("");

const addWord = () => {
  if (newWord.value.trim() !== "") {
    items.value.push(newWord.value);
    newWord.value = "";
  }
};

const popWord = () => {
  if (items.value.length > 0) {
    items.value.pop();
  }
};
</script>

<template>
  <h1>This is the landing page</h1>
  <input type="text" v-model="newWord" placeholder="Enter a word" @keyup.enter="addWord" />
  <button @click="addWord">Add Word</button>
  <div>
    <BlogPost v-for="item in items">
      <template #blogTitle>{{ item.title }}</template>
      <template #blogDescription>{{ item.description }}</template>
    </BlogPost>
  </div>
  <button @click="popWord">Remove Last Word</button>
</template>
