<script setup lang="ts">

// TODO: Add pass in to modal content margin and height
const { elementId, closeModal, customContentStyle } = defineProps<{
  elementId: string
  // For the modal to close itself if it wants
  closeModal: Function
  customContentStyle: string
}>()

window.onclick = function (event) {
  var modal = document.getElementById(elementId)
  if (event.target == modal) {
    closeModal()
  }
}

</script>

<template>
  <div :id="elementId" class="modal">
    <div class="modal-content" :style=customContentStyle>
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
  background-color: rgba(0, 0, 0, 0.4);
  display: grid;
  place-items: center;
}

.v-enter-active,
.v-leave-active {
  transition: opacity 0.3s cubic-bezier(0.4, 0, 0.2, 1), visibility 0.3s linear;
  visibility: visible;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
  visibility: hidden;
}

.modal-content {
  background-color: #fefefe;
  box-shadow: 0px 0px 20px 0px black;
  padding: 20px;
  border: 1px solid #888;
  border-radius: 15px 15px 15px 15px;
  margin: 15%;
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
