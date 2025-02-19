<script setup lang="ts">
import { ref } from 'vue';
import LoginButton from './LoginButton.vue';
import PasswordInput from './PasswordInput.vue';
import TextInput from './TextInput.vue';
import { RequestUserToken } from '@/api_calls';
import { StoreUserJWT } from '@/auth/auth';

const emit = defineEmits<{
  (event: 'token-update', token: string): void
}>()

const username = ref('')
const password = ref('')
async function SendLoginRequest() {
  let token = await RequestUserToken(username.value, password.value)
  StoreUserJWT(token);
  emit('token-update', token)
}
</script>

<template>
  <TextInput placeholder="Username" @textChange="(user: string) => username = user" />
  <PasswordInput placeholder="Password" @passwordChange="(pass: string) => password = pass" />
  <div style="display:grid; place-items: center;">
    <LoginButton @click="SendLoginRequest" />
  </div>
</template>

<style scoped>
@import url('../../colours.css');

input,
button {
  display: block;
  margin: 1ch 1ch 0ch 1ch;
}
</style>
