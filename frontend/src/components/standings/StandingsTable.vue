<script setup lang="ts">
import type { StandingEntry } from '@/types/standings'

defineProps<{
  standings: StandingEntry[]
}>()
</script>

<template>
  <table class="w-full border-collapse text-2xl">
    <thead class="bg-gray-100">
    <tr>
      <th class="p-2 text-left">#</th>
      <th class="p-2 text-left">Team</th>
      <th class="p-2 text-left">Affiliation</th>
      <th class="p-2 text-center">W</th>
      <th class="p-2 text-center">D</th>
      <th class="p-2 text-center">L</th>
      <th class="p-2 text-center">GF</th>
      <th class="p-2 text-center">GA</th>
      <th class="p-2 text-center">+/-</th>
      <th class="p-2 text-center">Pts</th>
    </tr>
    </thead>

    <tbody>
    <tr
        v-for="entry in standings"
        :key="entry.team.id"
        class="border-b last:border-b-0"
    >
      <td class="p-2">{{ entry.rank }}</td>

      <td class="p-2 font-medium">
        {{ entry.team.name }}
      </td>

      <td class="p-2 text-gray-600">
        {{ entry.team.affiliation }}
      </td>

      <td class="p-2 text-center">{{ entry.wins }}</td>
      <td class="p-2 text-center">{{ entry.draws }}</td>
      <td class="p-2 text-center">{{ entry.losses }}</td>
      <td class="p-2 text-center">{{ entry.scored }}</td>
      <td class="p-2 text-center">{{ entry.conceded }}</td>
      <td class="p-2 text-center">{{ entry.difference }}</td>

      <td class="p-2 text-center font-semibold">
        {{ entry.points }}
      </td>
    </tr>
    </tbody>
  </table>
</template>
