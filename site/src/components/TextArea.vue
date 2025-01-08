<script setup lang="ts">
import { ref } from 'vue';

const emit = defineEmits<{
  (event: 'detailChange', payload: string): void
}>()
const props = defineProps<{
  placeholder: string,
}>()

const postDetail = ref('')
function emitDetailChange() {
  emit('detailChange', postDetail.value)
}
</script>

<template>
  <textarea type="text" :placeholder="props.placeholder" v-model="postDetail" @input="emitDetailChange" />
</template>

<style scoped>
textarea {
  width: 100%;
  height: 150px;
  box-sizing: border-box;
  display: block;
  background-color: lightgrey;
  border: 2px solid lightgrey;
  border-radius: 5px 5px 5px 5px;
  font-size: 14px;
  padding: 7px 10px 7px 10px;
  font-weight: 500;
}
</style>
