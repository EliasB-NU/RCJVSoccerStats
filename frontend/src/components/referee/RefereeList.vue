<script setup lang="ts">
import RefereeCard from './RefereeCard.vue'

defineProps<{
  referees: string[]
  selectedReferee: string | null
  upcomingReferees: Set<string>
}>()

defineEmits<{
  (e: 'select', referee: string): void
}>()
</script>

<template>
  <aside
      class="p-4 border-r border-gray-700
           grid gap-4
           h-full"
      style="grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));"
  >
    <RefereeCard
        v-for="ref in referees"
        :key="ref"
        :name="ref"
        :selected="selectedReferee === ref"
        :upcoming="upcomingReferees.has(ref)"
        @select="$emit('select', ref)"
    />
  </aside>
</template>

