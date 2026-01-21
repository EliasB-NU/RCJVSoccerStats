<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import type { LeagueResponse, StandingsMessage } from '@/types/standings'
import LeagueStandings from '@/components/standings/LeagueStandings.vue'
import LoadingBar from '@/components/common/LoadingBar.vue'
import axios from 'axios'

const api = axios.create({
  baseURL: '/api/v1',
})

const leagues = ref<LeagueResponse[]>([])
const currentLeagueIndex = ref(0)
const standings = ref<StandingsMessage | null>(null)
const progress = ref(0)

let interval: number

const DISPLAY_TIME = 20000 // 20 seconds

const fetchLeagues = async (): Promise<LeagueResponse[]> => {
  const { data } = await api.get<LeagueResponse[]>('/leagues')
  return data
}

const fetchStandings = async (
    leagueAbbrev: string
): Promise<StandingsMessage> => {
  let { data } = await api.get<StandingsMessage>(
      `/standings/${leagueAbbrev}`
  )
  data.standings = data.standings
      .slice()
      .sort((a, b) => a.rank - b.rank)
  return data
}

const loadLeague = async () => {
  const league = leagues.value[currentLeagueIndex.value]
  standings.value = await fetchStandings(league?.abbreviation ?? '')
}

let startTime = 0
let animationFrameId: number

const animateProgress = (timestamp: number) => {
  if (!startTime) startTime = timestamp

  const elapsed = timestamp - startTime
  progress.value = Math.min((elapsed / DISPLAY_TIME) * 100, 100)

  if (elapsed < DISPLAY_TIME) {
    animationFrameId = requestAnimationFrame(animateProgress)
  }
}

const startProgress = () => {
  progress.value = 0
  startTime = 0
  cancelAnimationFrame(animationFrameId)
  animationFrameId = requestAnimationFrame(animateProgress)
}

const nextLeague = async () => {
  currentLeagueIndex.value =
      (currentLeagueIndex.value + 1) % leagues.value.length

  await loadLeague()
  startProgress()
}


onMounted(async () => {
  leagues.value = await fetchLeagues()
  await loadLeague()
  startProgress()

  interval = window.setInterval(nextLeague, DISPLAY_TIME)
})

onUnmounted(() => {
  clearInterval(interval)
  cancelAnimationFrame(animationFrameId)
})
</script>

<template>
  <div class="min-h-screen bg-gray-50 p-6">
    <LoadingBar
        v-if="leagues.length"
        :progress="progress"
        :league-name="leagues[currentLeagueIndex]?.league_name ?? 'League not found'"
        class="mb-6"
    />

    <LeagueStandings
        v-if="standings"
        :data="standings"
    />
  </div>
</template>
