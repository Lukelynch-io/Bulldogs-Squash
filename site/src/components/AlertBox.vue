<script setup lang="ts">
import type { MessageType } from '@/datatypes/MessageType';

const props = defineProps<{
  messageType: MessageType
  message: string
}>()
</script>

<template>
  <div>
    <p>{{ props.message }}</p>
  </div>
</template>

<style scoped>
p {
  border: 5px solid greenyellow;
  border-radius: 5px 5px 5px 5px;
}

div {
  display: grid;
  place-content: center;
  padding: 10px;
}
</style>
