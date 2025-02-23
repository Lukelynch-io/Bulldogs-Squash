<script setup lang="ts">
import { ref, provide, reactive, onBeforeMount } from 'vue';
import Header from './components/Header.vue';
import LoginForm from './components/LoginForm.vue';
import { GetUsername } from './api_calls';
import Modal from './components/Modal.vue';
import { CurrentUser, StoreUserJWT } from './auth/auth';

const isLoginActive = ref(false);
let currentUser = ref<CurrentUser | null>(null);
provide("currentUser", currentUser);
let currentUsername = ref("");

async function loadUserFromContentStore() {
  const tryToken = localStorage.getItem("UserJWT");
  if (tryToken !== null) {
    let username = await GetUsername(tryToken);

    var newCurrentUser = new CurrentUser(tryToken, username);
    currentUser.value = reactive(newCurrentUser);
    currentUsername.value = newCurrentUser.username
  }
}

onBeforeMount(async () => {
  await loadUserFromContentStore();

})

const toggleLoginModal = () => {
  isLoginActive.value = !isLoginActive.value
};

async function HandleTokenUpdate(newToken: string) {
  StoreUserJWT(newToken);
  await loadUserFromContentStore();
}

</script>

<template>
  <Header v-bind:is-login-showing="toggleLoginModal" />
  <Transition>
    <Modal v-if="isLoginActive" :closeModal="toggleLoginModal" :custom-content-style="''">
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
