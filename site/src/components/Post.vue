<script setup lang="ts">
import type PostData from '@/datatypes/PostData';
const postData = defineProps<PostData>()

const emit = defineEmits<{
  (event: 'expand-post', postId: string): void
}>();

</script>
<template>
  <div class="blog-post" @click="emit('expand-post', postData.id)">
    <h2>
      {{ postData.title }}
    </h2>
    <p>
      {{ postData.description }}
    </p>
    <img :src=postData.imageUrl alt="Bulldog">
  </div>
</template>

<style>
.blog-post {
  color: var(--shadow-color);
  background-color: white;
  border: 1px solid #ccc;
  border-radius: 6px;
  padding: 2rem;
  text-align: center;
  display: grid;
  grid-template-rows: subgrid;
  grid-row: span 4;
}


.blog-post img {
  max-width: 100%;
  display: block;
}

.blog-post:hover {
  box-shadow: 0px 2px 10px 0px var(--shadow-color);
  transition: box-shadow 0.1s linear;
  cursor: pointer;
}
</style>
