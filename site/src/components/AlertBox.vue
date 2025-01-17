<script setup lang="ts">
import { MessageType } from '@/datatypes/MessageType';
import { onMounted, ref } from 'vue';

const { messageType, messageTitle, messageDescription } = defineProps<{
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

function determineAlertColour() {
  if (messageType == MessageType.Success)
    return "green-alert"
  switch (messageType) {
    case MessageType.Info:
      return "blue-alert";
    case MessageType.Warning:
      return "orange-alert"
    case MessageType.Error:
      return "red-alert"
  }
}
</script>

<template>
  <div class="alert-area" :class="{ show: showAlert }">
    <div class="alert-wrapper" :class="determineAlertColour()">
      <h3 class="alert-header">{{ messageTitle }}</h3>
      <p>{{ messageDescription }}</p>
    </div>
  </div>
</template>

<style>
.alert-header {
  text-align: center;
}

.green-alert {
  border: 5px solid greenyellow;
  background: green;
}

.blue-alert {
  border: 5px solid aliceblue;
  background: blue;
}

.orange-alert {
  border: 5px solid yellow;
  background: orange;
}

.red-alert {
  border: 5px solid #9d0000;
  background: red;
}

.alert-wrapper {
  color: white;
  border-radius: 5px 5px 5px 5px;
  padding-left: 1rem;
  padding-right: 1rem;
  box-shadow: 0px 5px 5px grey;
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
