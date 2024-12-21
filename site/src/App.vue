<script setup lang="ts">
import { ref } from 'vue';
import Header from './components/Header.vue';
import LoginModal from './components/LoginModal.vue';

const isLoginActive = ref(false);

const toggleLoginModel = () => {
  isLoginActive.value = !isLoginActive.value
};
</script>

<template>
  <Header v-bind:is-login-showing="toggleLoginModel" />
  <LoginModal v-bind:flag="isLoginActive" :toggle-flag="toggleLoginModel" />
  <main>
    <RouterView />
  </main>
</template>

<style scoped>
.bulldog-image-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
}

main {
  margin-bottom: 1ch;
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
