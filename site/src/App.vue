<script setup lang="ts">
import { ref } from 'vue';
import Header from './components/Header.vue';
import LoginForm from './components/LoginForm.vue';
import { GetUsername } from './api_calls';
import Modal from './components/Modal.vue';
import { StoreUserJWT } from './auth/auth';

const isLoginActive = ref(false);
const token = ref("")
const loggedInUser = ref("");

const toggleLoginModal = () => {
  isLoginActive.value = !isLoginActive.value
};

async function HandleTokenUpdate(newToken: string) {
  token.value = newToken
  StoreUserJWT(newToken);
  loggedInUser.value = await GetUsername(token.value)
}

</script>

<template>
  <Header v-bind:is-login-showing="toggleLoginModal" :loggedInUsername="loggedInUser" />
  <Transition>
    <Modal v-if="isLoginActive" :elementId="'login-modal'" :closeModal="toggleLoginModal" :custom-content-style="''">
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
}

.wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
