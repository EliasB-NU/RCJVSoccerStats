<script setup lang="ts">
defineProps<{
  name: string
  selected: boolean
  upcoming: boolean
}>()

defineEmits<{
  (e: 'select'): void
}>()
</script>

<template>
  <div
      @click="$emit('select')"
      class="rounded cursor-pointer transition
           flex items-center justify-center text-center
           min-h-20 p-4"
      :class="[
      selected ? 'bg-blue-600' : 'bg-gray-800',
      upcoming ? 'border-2 border-red-500 animate-bounce' : ''
    ]"
  >
    <h3 class="text-lg font-semibold">
      {{ name }}
    </h3>
  </div>
</template>

