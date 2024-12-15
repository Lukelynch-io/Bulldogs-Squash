<script setup lang="ts">
import type BlogPostData from '@/datatypes/BlogPost';
const foo = defineProps<BlogPostData>(
)
</script>
<template>
  <div class="blog-post">
    <h2>
      {{ foo.title }}
    </h2>
    <p>
      {{ foo.description }}
    </p>
    <img :src=foo.imageUrl alt="Bulldog">
  </div>
</template>

<style scoped>
img {
  max-width: 100%;
  display: block;
}

.blog-post {
  color: var(--blue-contrast);
  border: 1px solid #ccc;
  border-radius: 6px;
  padding: 2rem;
  text-align: center;
  display: grid;
  grid-template-rows: subgrid;
  grid-row: span 4;
}

.blog-post:hover {
  box-shadow: 0px 2px 10px 0px var(--shadow-color);
  transition: box-shadow 0.1s linear;
}
</style>
