<script setup lang="ts">
import { ref } from 'vue';

const { flag, toggleFlag } = defineProps<{
  flag: boolean
  toggleFlag: Function
}>()

window.onclick = function (event) {
  var modal = document.getElementById("login-modal")
  if (event.target == modal) {
    toggleFlag()
  }
}

</script>

<template>
  <div id="login-modal" class="modal" :class="{ 'is-visible': flag, 'is-hidden': !flag }">
    <div class="modal-content">
      <div class="modal-content-wrapper">
        <slot></slot>
      </div>
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
  /* cubic-bezier creates a smoother transition curve */
  transition: opacity 0.3s cubic-bezier(0.4, 0, 0.2, 1), visibility 0.3s linear;
  background-color: rgba(0, 0, 0, 0.4);

  display: grid;
  place-items: center;
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
  box-shadow: 0px 0px 20px 0px black;
  margin: 15% auto;
  /* 15% from the top and centered */
  padding: 20px;
  border: 1px solid #888;
  border-radius: 15px 15px 15px 15px;
  height: 30%;
  /* Could be more or less, depending on screen size */
  display: grid;
  place-items: center;
}

.modal-content-wrapper {
  display: block;
  text-align: center;
}

.displayblock {
  visibility: visible;
}
</style>
