<script setup lang="ts">
import { RouterLink } from 'vue-router';
const props = defineProps<{
  isLoginShowing: Function
}>()

const toggleDarkMode = () => {
  document.body.classList.toggle('dark')
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
          <RouterLink class="hamburger-menu-item" to="/Post">Contact Us</RouterLink>
        </li>
        <li>
          <p class="hamburger-menu-item" @click="props.isLoginShowing()">Login</p>
        </li>
        <li>
          <button class="hamburger-menu-item" @click="toggleDarkMode">Dark Mode</button>
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

.flex {
  display: flex;
  justify-content: flex-end;
}

.flex-1 {
  flex: 1;
}

/* TODO: Clean this up*/
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
  font-weight: 600;
  text-decoration: none;
  white-space: nowrap;
  overflow: hidden;
}

.hamburger-menu-item:last-of-type {
  border-radius: 0px 0px 5px 0px;
  padding-right: 1ch;
}

.hamburger-menu-item:hover {
  text-decoration-line: underline;
}
</style>
