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
      class="p-4 rounded cursor-pointer transition"
      :class="[
      selected ? 'bg-blue-600' : 'bg-gray-800',
      upcoming ? 'border-2 border-red-500' : ''
    ]"
  >
    <h3 class="text-lg font-semibold">{{ name }}</h3>
  </div>
</template>
