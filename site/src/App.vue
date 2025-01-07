<script setup lang="ts">
import { ref } from 'vue';
import Header from './components/Header.vue';
import LoginModal from './components/LoginModal.vue';
import { GetUsername } from './api_calls';

const isLoginActive = ref(false);
const token = ref("")
const loggedInUser = ref("");

const toggleLoginModel = () => {
  isLoginActive.value = !isLoginActive.value
};

async function HandleTokenUpdate(newToken: string) {
  token.value = newToken
  loggedInUser.value = await GetUsername(token.value)
}

</script>

<template>
  <Header v-bind:is-login-showing="toggleLoginModel" :loggedInUsername="loggedInUser" />
  <LoginModal :flag="isLoginActive" :toggle-flag="toggleLoginModel" @tokenUpdate="HandleTokenUpdate" />
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
