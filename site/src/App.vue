<script setup lang="ts">
import LandingPage from './components/LandingPage.vue'
</script>

<template>
  <header>
    <div class="bulldog-image-wrapper">
      <img alt="Vue logo" class="logo" src="../public/img/bulldogs.png" width="125" height="125" />
    </div>
  </header>

  <main>
    <LandingPage />
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
