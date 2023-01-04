<script>
// import Team from '@/JS/Team';
import LeagueTableTeam from './LeagueTableTeam.vue';
export default {
    name: "LeagueTable",
    props: {
        name: String,
        teams: {
            type: Array,
            default: () => ([])
        }
    },
    components: { LeagueTableTeam }
}
</script>

<template>
    <div class="title">{{ name }}</div>
  <div class="mr-3 ml-3 box">
    <div class="mr-2 ml-2 columns is-mobile">
        <div class="column is-one-seventh">
            <h2 class="subtitle">Position</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">Team Name</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">Games Played</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">Wins</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">Draws</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">Losses</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">Points</h2>
        </div>
    </div>
    <LeagueTableTeam  v-for="(team, index) in teams"  :key="team.name" :position="index+=1" :name="team.name" :games_played="team.gamesPlayed" :wins="team.wins" :draws="team.draws" :losses="team.losses" :points="team.points"/>
    
  </div>
</template>