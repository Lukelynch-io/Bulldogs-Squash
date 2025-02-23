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

</script>

<template>
  <div class="header">
    <div class="flex-1"></div>
    <h1>
      <RouterLink to="/">Bulldogs Squash Club</RouterLink>
    </h1>
    <div class="flex flex-1">
      <ul class="hamburger-menu">
        <li>
          <RouterLink class="hamburger-menu-item" to="/About">About</RouterLink>
        </li>
        <li>
          <RouterLink class="hamburger-menu-item" to="/Fixtures">Fixtures</RouterLink>
        </li>
        <li>
          <RouterLink class="hamburger-menu-item" to="/Contact">Contact Us</RouterLink>
        </li>
        <li>
          <p class="hamburger-menu-item" @click="isLoginShowing()">{{ currentUser?.username == "" ? "Login" :
            currentUser?.username
            }}</p>
          <!-- TODO: Show account menu options list item instead -->
        </li>
        <li>
          <img v-if="themeToggleTitle == 'Light Mode'" src="../../public/img/light_mode_icon.svg"
            @click="toggleDarkMode" />
          <img v-else src="../../public/img/dark_mode_icon.svg" @click="toggleDarkMode" />
        </li>
      </ul>
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
  padding-left: 1ch;
  display: flex;
  align-items: center;
  justify-content: flex-end;
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
}

.hamburger-menu-item {
  padding-left: 2.5px;
  padding-right: 2.5px;
  margin-right: 0.5rem;
  font-weight: 600;
  text-decoration: none;
  white-space: nowrap;
  overflow: hidden;
}

.hamburger-menu-item:hover {
  text-decoration-line: underline;
  cursor: pointer;
}
</style>
