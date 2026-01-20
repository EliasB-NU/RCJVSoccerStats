<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const now = ref(new Date())
let interval: number | null = null

onMounted(() => {
  interval = window.setInterval(() => {
    now.value = new Date()
  }, 1000)
})

onUnmounted(() => {
  if (interval) clearInterval(interval)
})
</script>

<template>
  <div class="text-9xl font-mono">
    {{ now.toLocaleTimeString() }}
  </div>
</template>
