<script setup lang="ts">
import { ref } from 'vue';
import Header from './components/Header.vue';
import LoginModal from './components/LoginModal.vue';
import { jwtDecode } from 'jwt-decode';
import axios from 'axios';

const isLoginActive = ref(false);
const loggedInUser = ref("");

const toggleLoginModel = () => {
  isLoginActive.value = !isLoginActive.value
};

const storeBearerToken = (token: string) => {
  localStorage.setItem("bearerToken", token)
  const decoded = jwtDecode(token)
  console.log(decoded)
  GetUsername(token)
}

function GetUsername(token: string) {
  axios.get("/api/auth/user", {
    headers: {
      "Authorization": "Bearer " + token
    }
  }).then((response) => {
    loggedInUser.value = response.data.username
  })
}

</script>

<template>
  <Header v-bind:is-login-showing="toggleLoginModel" :loggedInUsername="loggedInUser" />
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
