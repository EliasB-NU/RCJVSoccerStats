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
      class="w-1/3 p-4 space-y-3 border-r border-gray-700
         h-full overflow-y-auto"
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
