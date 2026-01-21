<script setup lang="ts">
import type { Stage } from '@/types/matches.ts'
import MatchesTable from "@/components/matches/MatchesTable.vue";

defineProps<{
  stage: { league: string; stage: Stage } | undefined
}>()
</script>

<template>
  <div class="bg-white rounded shadow p-6">
    <h1 class="text-2xl font-bold mb-4">
      {{ stage?.stage.matches[0]?.league }} || {{ stage?.stage.name }}
    </h1>

    <MatchesTable :matches="stage" />

    <div class="mt-4 text-xs text-gray-500 text-left font-bold">
      Last published: {{ new Date(stage?.stage.standings_last_published ?? '').toLocaleString() }}
    </div>
  </div>
</template>
