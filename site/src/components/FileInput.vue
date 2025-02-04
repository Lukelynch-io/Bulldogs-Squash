<script setup lang="ts">
const { acceptedFileTypes } = defineProps<{
  acceptedFileTypes: string
  onFileChange: (event: Event) => void
}>()

</script>

<template>
  <input type="file" :accept="acceptedFileTypes" v-on:change="onFileChange" />
</template>
