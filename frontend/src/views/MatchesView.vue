<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'
import LoadingBar from '@/components/common/LoadingBar.vue'
import type {Stage, Stages} from '@/types/matches.ts'
import StageMatches from "@/components/matches/StageMatches.vue";

/* ---------------- Constants ---------------- */

const DISPLAY_TIME = 20000 // 20 seconds

/* ---------------- Routing ---------------- */

const route = useRoute()

const leagueAbbrevs = computed(() => {
  const param = route.query.leagues
  if (!param) return []
  if (Array.isArray(param)) return param
  return param.split(',').map(l => l.trim())
})

/* ---------------- State ---------------- */

const allStages = ref<
    { league: string; stage: Stages[number] }[]
>([])

const currentIndex = ref(0)
const current = computed(() => allStages.value[currentIndex.value])

const progress = ref(0)

/* ---------------- Data Fetching ---------------- */

async function fetchLeague(league: string): Promise<Stages | null>{
  let stages: Stages = []
  try {
    await axios.get(`/api/v1/matches/${league}`)
        .then(res => {
          stages = res.data.filter(
              (s: Stage) => s.matches !== null,
          )
        })
  } catch (error) {
    console.error(`Failed to fetch league ${league}:`, error)
    return null
  }
  for (const stage of stages) {
    stage.matches.filter(
        m => m.start !== null || m.team1 !== null && m.team2 !== null,
    ).sort((a, b) => {
      return new Date(a.start!).getTime() - new Date(b.start!).getTime()
    }
    )
  }
  return stages
}

const loadAllStages = async () => {
  const stagesList: { league: string; stage: Stages[number] }[] = []

  for (const league of leagueAbbrevs.value) {
    if (league === null) continue
    const stages = await fetchLeague(league)
    if (stages) {
      for (const stage of stages) {
        stagesList.push({ league, stage })
      }
    }
  }

  allStages.value = stagesList
}

/* ---------------- Progress Animation ---------------- */

let startTime = 0
let animationFrameId: number
let interval: number

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

const nextStage = () => {
  currentIndex.value =
      (currentIndex.value + 1) % allStages.value.length
  startProgress()
}

/* ---------------- Lifecycle ---------------- */

onMounted(async () => {
  await loadAllStages()
  setInterval(loadAllStages, 20 * 1000) // Refresh every 60 seconds

  if (!allStages.value.length) return

  startProgress()
  interval = window.setInterval(nextStage, DISPLAY_TIME)
})

onUnmounted(() => {
  clearInterval(interval)
  cancelAnimationFrame(animationFrameId)
})
</script>

<template>
  <div class="min-h-screen bg-white p-6 flex flex-col">
    <!-- Loading Bar -->
    <LoadingBar
        v-if="current"
        :progress="progress"
        :league-name="`${current.stage.matches[0]?.league} – ${current.stage.name}`"
        class="mb-6"
    />

    <!-- Matches Table -->
    <StageMatches
      :stage="current"
    />
  </div>
</template>
