<script>
import Team from '@/JS/Team';
// import LeagueTableTeam from './LeagueTableTeam.vue';
export default {
    name: "LeagueTable",
    props: {
        teams: Team.generateTeams
    },
    // components: { LeagueTableTeam }
}
</script>

<template>
  <div class="box">
    <div class="columns">
        <div class="column is-one-fifths">
            <h2 class="subtitle">Team Name</h2>
        </div>
        <div class="column is-one-fifths">
            <h2 class="subtitle">Games Played</h2>
        </div>
        <div class="column is-one-fifth">
            <h2 class="subtitle">Wins</h2>
        </div>
        <div class="column is-one-fifth">
            <h2 class="subtitle">Draws</h2>
        </div>
        <div class="column is-one-fifth">
            <h2 class="subtitle">Losses</h2>
        </div>
    </div>
    <h1 v-for="team in teams" :key="team">{{ team.name }}</h1>
    <!-- <LeagueTableTeam  v-for="team in teams"  :key="team.name" /> -->
    
  </div>
</template>