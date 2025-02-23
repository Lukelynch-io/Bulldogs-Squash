<script setup lang="ts">
import type { CurrentUser } from '@/auth/auth';
import { inject, ref, type Ref } from 'vue';
import { RouterLink } from 'vue-router';
const { isLoginShowing } = defineProps<{
  isLoginShowing: Function
}>()
let currentUser = inject<Ref<CurrentUser | null>>("currentUser");
const themeToggleTitle = ref("Dark Mode")
const toggleDarkMode = () => {
  document.body.classList.toggle('dark')
  if (themeToggleTitle.value === "Dark Mode") {
    themeToggleTitle.value = "Light Mode"
  } else {
    themeToggleTitle.value = "Dark Mode"
  }
}

function logOut() {
  currentUser!.value = null;
  localStorage.removeItem("UserJWT");
}

</script>

<template>
  <div class="header">
    <div class="flex-1"></div>
    <h1>
      <RouterLink to="/">Bulldogs Squash Club</RouterLink>
    </h1>
    <div class="flex flex-1">
      <div class="hamburger-menu">
        <RouterLink class="hamburger-menu-item" to="/About">About</RouterLink>
        <RouterLink class="hamburger-menu-item" to="/Fixtures">Fixtures</RouterLink>
        <RouterLink class="hamburger-menu-item" to="/Contact">Contact Us</RouterLink>
        <div v-if="currentUser != null && currentUser.username != ''" class="dropdown">
          <RouterLink class="hamburger-menu-item" to="/Profile">{{ currentUser?.username }}</RouterLink>
          <div class="dropdown-content">
            <p @click="logOut()">Log out</p>
          </div>
        </div>
        <p v-else class="hamburger-menu-item" @click="isLoginShowing()">Login</p>
        <img style="width: 30px;" v-if="themeToggleTitle == 'Light Mode'" src="../../public/img/light_mode_icon.svg"
          @click="toggleDarkMode" />
        <img style="width: 30px;" v-else src="../../public/img/dark_mode_icon.svg" @click="toggleDarkMode" />
      </div>
    </div>
  </div>
</template>

<style scoped>
@import url('../../colours.css');

.header {
  background-color: var(--blue-9);
  color: var(--blue-contrast);
  border-radius: 0px 0px 5px 5px;
  width: 100%;
  height: 60px;
  padding-left: 1ch;
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

.dropdown {
  display: flex;
}

.dropdown:not(:hover) .dropdown-content {
  display: none;
}

.dropdown-content {
  display: block;
  position: absolute;
  top: 60px;
  background-color: white;
  color: black;
  box-shadow: 3px 3px 5px black;
}

.dropdown-content a,
.dropdown-content p {
  display: block;
  color: black;
  cursor: pointer;
  margin: 0;
  min-width: 50px;
  min-height: 50px;
}

img {
  cursor: pointer;
  filter: brightness(100%) !important;
  padding-right: 0.5rem;
}

.flex {
  display: flex;
  justify-content: flex-end;
}

.flex-1 {
  flex: 1;
  height: 100%;
}

/* TODO: Clean this up */
.header h1,
a {
  color: var(--blue-contrast);
  margin-top: 0.5ch;
  margin-bottom: 0.5ch;
  text-decoration: none;
}

.hamburger-menu {
  display: flex;
  list-style: none;
  justify-content: space-between;
  align-items: stretch;
  height: 100%;
}

.hamburger-menu-item {
  padding-left: 2.5px;
  padding-right: 2.5px;
  margin-right: 0.5rem;
  font-weight: 600;
  text-decoration: none;
  white-space: nowrap;
  overflow: hidden;
  display: flex;
  align-items: center;
}

.hamburger-menu-item:hover {
  text-decoration-line: underline;
  cursor: pointer;
}
</style>
