<script setup lang="ts">
import type { Stage } from '@/types/matches.ts'

defineProps<{
  matches: { league: string; stage: Stage } | undefined
}>()

/* ---------------- Helpers ---------------- */

const formatTime = (iso?: string) =>
    iso ? new Date(iso).toLocaleString([], { day: '2-digit', month: '2-digit', hour: '2-digit', minute: '2-digit' }) : '-'

const parseDurationToMs = (duration?: string | null | undefined): number => {
  if (!duration) return 0

  // Expected format: HH:MM:SS
  const [hours, minutes, seconds] = duration.split(':').map(Number)

  if (hours === undefined || minutes === undefined || seconds === undefined) {
    return 0
  }

  return ((hours * 60 + minutes) * 60 + seconds) * 1000
}


const isFinished = (start?: string, duration?: string | null | undefined) => {
  if (!start || !duration) return false

  const startTime = new Date(start).getTime()
  const durationMs = parseDurationToMs(duration)

  return Date.now() >= startTime + durationMs
}

const isLive = (start?: string, duration?: string | null | undefined) => {
  if (!start || !duration) return false

  const now = Date.now()
  const startTime = new Date(start).getTime()
  const endTime = startTime + parseDurationToMs(duration)

  return now >= startTime && now < endTime
}
</script>

<template>
  <div>
    <div
        v-if="matches"
        class="flex-1 overflow-y-auto"
    >
      <table class="w-full text-gray-900 border-collapse">
        <thead>
        <tr class="border-b border-gray-300 text-2xl">
          <th class="py-2 text-left">Time</th>
          <th class="py-2 text-left">Team 1</th>
          <th class="py-2 text-center">Score</th>
          <th class="py-2 text-left">Team 2</th>
          <th class="py-2 text-left">Field</th>
          <th class="py-2 text-center">Pts</th>
        </tr>
        </thead>

        <tbody>
        <tr
            v-for="match in matches.stage.matches"
            :key="match.number"
            class="border-b border-gray-200 text-xl"
            :class="[
              isFinished(match.start, match.duration)
               ? 'bg-green-50 text-gray-500'
              : isLive(match.start, match.duration)
               ? 'bg-yellow-100 font-semibold'
              : ''
            ]"
        >
          <td class="py-2">
            {{ formatTime(match.start) }}
          </td>

          <td class="py-2">
            {{ match.team1?.name ?? '—' }}
          </td>

          <td class="py-2 text-center font-semibold">
            {{ match.goals1 }} : {{ match.goals2 }}
          </td>

          <td class="py-2">
            {{ match.team2?.name ?? '—' }}
          </td>

          <td class="py-2">
            {{ match.pitch ?? '—' }}
          </td>

          <td class="py-2 text-center">
            {{ match.points1 }} : {{ match.points2 }}
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>