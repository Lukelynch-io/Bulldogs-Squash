<script setup lang="ts">
import Header from './components/Header.vue'
import BlogPostCollection from './components/BlogPostCollection.vue';
</script>

<template>

  <Header />
  <!-- <header> -->
  <!--   <div class="bulldog-image-wrapper"> -->
  <!--     <img alt="Vue logo" class="logo" src="../public/img/bulldogs.png" width="125" height="125" /> -->
  <!--   </div> -->
  <!-- </header> -->

  <main>
    <BlogPostCollection />
  </main>
</template>

<style scoped>
.bulldog-image-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
}

.logo {
  max-width: 100%;
  /* Optional: keeps the image responsive */
}

.wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
