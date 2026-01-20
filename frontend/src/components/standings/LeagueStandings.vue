<script setup lang="ts">
import type { StandingsMessage } from '@/types/standings'
import StandingsTable from './StandingsTable.vue'

defineProps<{
  data: StandingsMessage
}>()
</script>

<template>
  <div class="bg-white rounded shadow p-6">
    <h1 class="text-2xl font-bold mb-4">
      {{ data.league }} || {{ data.stage_name }}
    </h1>

    <StandingsTable :standings="data.standings" />

    <div class="mt-4 text-xs text-gray-500 text-left font-bold">
      Last published: {{ data.last_published }}
    </div>
  </div>
</template>
