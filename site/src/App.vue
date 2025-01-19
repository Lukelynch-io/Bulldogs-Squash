<script setup lang="ts">
import { ref } from 'vue';
import Header from './components/Header.vue';
import LoginForm from './components/LoginForm.vue';
import { GetUsername } from './api_calls';
import Modal from './components/Modal.vue';

const isLoginActive = ref(false);
const token = ref("")
const loggedInUser = ref("");

const toggleLoginModal = () => {
  isLoginActive.value = !isLoginActive.value
};

async function HandleTokenUpdate(newToken: string) {
  token.value = newToken
  loggedInUser.value = await GetUsername(token.value)
}

</script>

<template>
  <Header v-bind:is-login-showing="toggleLoginModal" :loggedInUsername="loggedInUser" />
  <Transition>
    <Modal v-if="isLoginActive" :elementId="'login-modal'" :closeModal="toggleLoginModal">
      <LoginForm @token-update="HandleTokenUpdate" />
    </Modal>
  </Transition>
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
