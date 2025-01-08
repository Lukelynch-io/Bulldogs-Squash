<script setup lang="ts">
const { acceptedFileTypes } = defineProps<{
  acceptedFileTypes: string
  onFileChange: Function
}>()

</script>

<template>
  <input type="file" :accept="acceptedFileTypes" v-on:change="onFileChange" />
</template>
