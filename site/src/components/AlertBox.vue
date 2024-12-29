<script setup lang="ts">
import type { MessageType } from '@/datatypes/MessageType';
import { onMounted, ref } from 'vue';

const props = defineProps<{
  messageType: MessageType
  messageTitle: string
  messageDescription: string
}>()
const showAlert = ref(false);

onMounted(() => {
  setTimeout(() => {
    showAlert.value = true;
  }, 100);
})
</script>

<template>
  <div class="alert-area" :class="{ show: showAlert }">
    <div class="alert-wrapper">
      <h3 class="alert-header">{{ props.messageTitle }}</h3>
      <p>{{ props.messageDescription }}</p>
    </div>
  </div>
</template>

<style>
.alert-header {
  text-align: center;
}

.alert-wrapper {
  border: 5px solid greenyellow;
  background: green;
  color: white;
  border-radius: 5px 5px 5px 5px;
  padding-left: 1rem;
  padding-right: 1rem;
}

.alert-area {
  display: grid;
  position: fixed;
  width: 100%;
  bottom: 0;
  z-index: 1;

  place-content: center;
  padding: 10px;

  transform: translateY(100%);
  transition: transform 0.5s ease-in-out;
}

.alert-area.show {
  transform: translateY(0);
}
</style>
