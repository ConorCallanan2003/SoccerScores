<template>
  <div>
    <LeagueTable v-for="league in leagues" v-bind:key="league.name" :name="league.name" :teams="league.teams" />
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import LeagueTable from './components/LeagueTable.vue';
import League from './JS/League';

export default {
  name: 'App',
  data() {
    return {
      leagues: League.retrieveLeagues()
    }
  },
  components: {
    // HelloWorld,
    LeagueTable
}
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
