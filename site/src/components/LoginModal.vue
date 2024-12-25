<script setup lang="ts">
import { ref } from 'vue';
import LoginButton from './LoginButton.vue';
import PasswordInput from './PasswordInput.vue';
import TextInput from './TextInput.vue';

const foo = defineProps<{
  flag: boolean
  toggleFlag: Function
}>()

window.onclick = function (event) {
  var modal = document.getElementById("login-modal")
  if (event.target == modal) {
    foo.toggleFlag()
  }
}

const username = ref('')
const password = ref('')
function SendLoginRequest() {
  console.log("Username emitted: " + username.value);
  console.log("Password: " + password.value)
}
</script>

<template>
  <div id="login-modal" class="modal" :class="{ 'is-visible': foo.flag, 'is-hidden': !foo.flag }">
    <div class="modal-content">
      <div class="modal-content-wrapper">
        <TextInput placeholder="Username" @usernameChange="(user: string) => username = user" />
        <PasswordInput placeholder="Password" @passwordChange="(pass: string) => password = pass" />
        <div style="display:grid; place-items: center;">
          <LoginButton @click="SendLoginRequest" />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url('../../colours.css');

.modal {
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  /* cubic-bezier creates a smoother transition curve */
  transition: opacity 0.3s cubic-bezier(0.4, 0, 0.2, 1), visibility 0.3s linear;
  background-color: rgba(0, 0, 0, 0.4);

  display: grid;
  place-items: center;
}

.modal.is-visible {
  opacity: 1;
  visibility: visible;
  /* Visible when displayed */
}

.modal.is-hidden {
  opacity: 0;
  visibility: hidden;
  /* Hidden when not displayed */
}

.modal-content {
  background-color: #fefefe;
  box-shadow: 0px 0px 20px 0px black;
  margin: 15% auto;
  /* 15% from the top and centered */
  padding: 20px;
  border: 1px solid #888;
  border-radius: 15px 15px 15px 15px;
  height: 30%;
  /* Could be more or less, depending on screen size */
  display: grid;
  place-items: center;
}

.modal-content input,
.modal-content button {
  display: block;
  margin: 1ch 1ch 0ch 1ch;
}

.modal-content-wrapper {
  display: block;
  text-align: center;
}

.displayblock {
  visibility: visible;
}
</style>
