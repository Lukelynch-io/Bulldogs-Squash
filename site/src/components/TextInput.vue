<script setup lang="ts">
import { ref } from 'vue';

const emit = defineEmits<{
  (event: 'textChange', text: string): void
}>()
const props = defineProps<{
  placeholder: string,
}>()

const text = ref('')
function emitTextChange() {
  emit('textChange', text.value)
}
</script>

<template>
  <input type="text" :placeholder="props.placeholder" v-model="text" @input="emitTextChange()" />
</template>

<style scoped>
input,
button {
  box-sizing: border-box;
  display: block;
  border: none;
  background-color: lightgrey;
  border-radius: 5px 5px 5px 5px;
  font-size: 14px;
  padding: 7px 5px 7px 5px;
  font-weight: 500;
}
</style>
