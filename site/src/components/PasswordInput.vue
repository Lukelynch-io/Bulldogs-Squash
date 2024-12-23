<script setup lang="ts">
const props = defineProps<{
  placeholder: string
}>()
</script>

<template>
  <input type="password" :placeholder="props.placeholder" />
</template>

<style scoped>
input,
button {
  display: block;
  border: none;
  background-color: lightgrey;
  border-radius: 5px 5px 5px 5px;
  font-size: 14px;
  margin: 1ch;
  padding: 7px 5px 7px 5px;
}
</style>
