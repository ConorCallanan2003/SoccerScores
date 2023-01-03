<script>
import Team from '@/JS/Team';
import LeagueTableTeam from './LeagueTableTeam.vue';
export default {
    name: "LeagueTable",
    data() {
        return {
            teams: Team.generateTeams()
        }
    },

    components: { LeagueTableTeam }
}
</script>

<template>
  <div class="box">
    <div class="columns">
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