<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import axios from 'axios'

import RefereeList from '@/components/referee/RefereeList.vue'
import RefereeMatches from '@/components/referee/RefereeMatches.vue'
import LiveClock from '@/components/common/LiveClock.vue'

import type { Match, RefereeResponse } from '@/types/referee'

/* ---------------- State ---------------- */

const referees = ref<string[]>([])
const selectedReferee = ref<string | null>(null)
const matches = ref<Match[]>([])
const showMatches = ref(false)
const upcomingReferees = ref<Set<string>>(new Set())

let hideTimeout: number | null = null
let pollInterval: number | null = null

/* ---------------- Helpers ---------------- */

const normalizeReferee = (name: string) =>
    name.toLowerCase().replace(/\s+/g, '')

const isWithin5Minutes = (start: string) => {
  const diff = new Date(start).getTime() - Date.now()
  return diff > 0 && diff <= 5 * 60 * 1000
}

/* ---------------- API ---------------- */

const loadReferees = async () => {
  const res = await axios.get<string[]>('/api/v1/referees')
  referees.value = res.data
      .slice()
      .sort((a, b) => a.localeCompare(b))
}

const loadMatches = async (referee: string) => {
  const res = await axios.get<RefereeResponse>(
      `/api/v1/referees/${normalizeReferee(referee)}/matches`
  )
  matches.value = res.data.matches
      .slice()
      .sort((a, b) => new Date(a.start).getTime() - new Date(b.start).getTime())
}

/* ---------------- Logic ---------------- */

const selectReferee = async (referee: string) => {
  selectedReferee.value = referee
  showMatches.value = true

  await loadMatches(referee)

  if (hideTimeout) clearTimeout(hideTimeout)
  hideTimeout = window.setTimeout(() => {
    showMatches.value = false
    selectedReferee.value = null
  }, 20000)
}

const pollUpcoming = async () => {
  const upcoming = new Set<string>()

  for (const ref of referees.value) {
    try {
      const res = await axios.get<RefereeResponse>(
          `/api/v1/referees/${normalizeReferee(ref)}/matches`
      )

      if (res.data.matches.some(m => isWithin5Minutes(m.start))) {
        upcoming.add(ref)
      }
    } catch {
      // ignore
    }
  }

  upcomingReferees.value = upcoming
}

/* ---------------- Lifecycle ---------------- */

onMounted(async () => {
  await loadReferees()
  pollInterval = window.setInterval(pollUpcoming, 60000)
})

onUnmounted(() => {
  if (hideTimeout) clearTimeout(hideTimeout)
  if (pollInterval) clearInterval(pollInterval)
})
</script>

<template>
  <div class="flex h-screen bg-gray-900 text-white">
    <RefereeList
        class="w-2/6 h-full"
        :referees="referees"
        :selected-referee="selectedReferee"
        :upcoming-referees="upcomingReferees"
        @select="selectReferee"
    />

    <main class="flex-1 h-full overflow-y-auto p-6 fade-bottom">
      <div class="flex justify-center">
        <RefereeMatches v-if="showMatches" :matches="matches" />
        <LiveClock v-else/>
      </div>
    </main>
  </div>
</template>

<style scoped>
.fade-bottom {
  -webkit-mask-image: linear-gradient(
      to bottom,
      black 85%,
      transparent 100%
  );
  mask-image: linear-gradient(
      to bottom,
      black 85%,
      transparent 100%
  );
}
</style>