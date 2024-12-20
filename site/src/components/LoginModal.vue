<script setup lang="ts">

const foo = defineProps<{
  display: boolean
}>()
</script>

<template>
  <div class="modal" :class="{ 'is-visible': foo.display, 'is-hidden': !foo.display }">
    <div class="modal-content">
      <p>This is the login modal</p>
    </div>
  </div>
</template>

<style scoped>
@import url('../../colours.css');

.modal {
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  opacity: 0;
  /* Start hidden */
  visibility: hidden;
  /* Prevent clicks when hidden */
  transition: opacity 0.5s ease, visibility 0.5s ease;
  transition: opacity 0.3s cubic-bezier(0.4, 0, 0.2, 1), visibility 0.3s linear;
  /* cubic-bezier creates a smoother transition curve */
}

.modal.is-visible {
  opacity: 1;
  visibility: visible;
  /* Visible when displayed */
}

.modal.is-hidden {
  opacity: 0;
  visibility: hidden;
  /* Hidden when not displayed */
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto;
  /* 15% from the top and centered */
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
  /* Could be more or less, depending on screen size */
}

.displayblock {
  visibility: visible;
}
</style>
