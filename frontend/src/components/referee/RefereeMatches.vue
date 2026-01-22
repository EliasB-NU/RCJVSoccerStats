<script setup lang="ts">
import type { Match } from '@/types/referee'

defineProps<{
  matches: Match[]
}>()

const isPast = (start: string) =>
    new Date(start).getTime() < Date.now()

const isLive = (start: string) => {
  const diff = Date.now() - new Date(start).getTime()
  return diff >= 0 && diff <= 2 * 60 * 60 * 1000 // 2h window
}
</script>

<template>
  <div class="space-y-4 w-full max-w-2xl">
    <div
        v-for="(match, idx) in matches"
        :key="idx"
        class="p-4 rounded transition"
        :class="isPast(match.start)
        ? 'bg-green-900/40 opacity-70'
        : 'bg-gray-800'"
    >
      <div class="flex justify-between">
        <span class="font-semibold">{{ match.league }}</span>
        <span>{{ new Date(match.start).toLocaleTimeString() }}</span>
      </div>

      <div class="text-xl mt-2">
        {{ match.team1 }} vs {{ match.team2 }}
      </div>

      <div class="text-sm text-gray-400 flex justify-between">
        <span>
          Field: {{ match.field }} | 2nd Ref: {{ match.referees }}
        </span>

        <span
            v-if="isPast(match.start)"
            class="text-green-400 font-semibold"
        >
          Finished
        </span>
      </div>
    </div>
  </div>
</template>