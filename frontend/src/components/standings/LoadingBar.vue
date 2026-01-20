<script setup lang="ts">
defineProps<{
  progress: number
  leagueName: string
}>()
</script>

<template>
  <div class="w-full">
    <div class="flex justify-between text-sm text-gray-600 mb-1">
      <span>{{ leagueName }}</span>
      <span>{{ Math.floor(progress) }}%</span>
    </div>
    <div class="h-2 w-full bg-gray-200 rounded">
      <div
          class="h-2 bg-blue-500 rounded transition-[width] duration-75 ease-linear"
          :style="{ width: `${progress}%` }"
      />
    </div>
  </div>
</template>
