<script setup lang="ts">
import { ref } from 'vue';
import LoginButton from './LoginButton.vue';
import PasswordInput from './PasswordInput.vue';
import TextInput from './TextInput.vue';
import { RequestUserToken } from '@/api_calls';
import Modal from './Modal.vue';

const { flag, toggleFlag } = defineProps<{
  flag: boolean
  toggleFlag: Function
}>()
const emit = defineEmits<{
  (event: 'tokenUpdate', token: string): void
}>()

window.onclick = function (event) {
  var modal = document.getElementById("login-modal")
  if (event.target == modal) {
    toggleFlag()
  }
}

const username = ref('')
const password = ref('')
async function SendLoginRequest() {
  let token = await RequestUserToken(username.value, password.value)
  emit('tokenUpdate', token)
}
</script>

<template>
  <Modal :flag="flag" :toggleFlag="toggleFlag">
    <TextInput placeholder="Username" @usernameChange="(user: string) => username = user" />
    <PasswordInput placeholder="Password" @passwordChange="(pass: string) => password = pass" />
    <div style="display:grid; place-items: center;">
      <LoginButton @click="SendLoginRequest" />
    </div>
  </Modal>
</template>

<style scoped>
@import url('../../colours.css');


input,
button {
  display: block;
  margin: 1ch 1ch 0ch 1ch;
}

.modal-content-wrapper {
  display: block;
  text-align: center;
}
</style>
