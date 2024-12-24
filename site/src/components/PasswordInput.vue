<script setup lang="ts">
import { ref } from 'vue';

const emit = defineEmits<{
  (event: 'passwordChange', password: string): void
}>()
const props = defineProps<{
  placeholder: string
}>()

const password = ref('')
function emitPasswordChange() {
  emit('passwordChange', password.value)
}
</script>

<template>
  <input type="password" @input="emitPasswordChange" v-model="password" :placeholder="props.placeholder" />
</template>

<style scoped>
input,
button {
  display: block;
  border: none;
  background-color: lightgrey;
  border-radius: 5px 5px 5px 5px;
  font-size: 14px;
  padding: 7px 5px 7px 5px;
  font-weight: 500;
}
</style>
