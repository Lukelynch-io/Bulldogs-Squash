<script setup lang="ts">
import { ref } from 'vue';
import Header from './components/Header.vue';
import LoginModal from './components/LoginModal.vue';

const isLoginActive = ref(false);

const toggleLoginModel = () => {
  isLoginActive.value = !isLoginActive.value
};

const storeBearerToken = (token: string) => {
  localStorage.setItem("bearerToken", token)
}
</script>

<template>
  <Header v-bind:is-login-showing="toggleLoginModel" />
  <LoginModal :flag="isLoginActive" :toggle-flag="toggleLoginModel" :storeBearerToken="storeBearerToken" />
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
