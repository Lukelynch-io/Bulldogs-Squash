<script setup lang="ts">
const props = defineProps<{
  // buttonAction: Function
}>()

function buttonAction() { }
</script>

<template>
  <button type="button" @click="buttonAction()">Login</button>
</template>

<style scoped>
@import url("../../colours.css");

button {
  background-color: var(--blue-9);
  border: none;
  border-radius: 5px 5px 5px 5px;
  color: white;
  padding: 3px 10px 3px 10px;
}
</style>
