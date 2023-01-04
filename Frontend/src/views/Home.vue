<template>
    <div class="columns">
        <div class="column is-one-sixth" v-for="(league, index) in leagues" v-bind:key="league.name">
            <button @click="handleNavigation(index)" class="button is-primary">{{ league.name }}</button>
        </div>
    </div>
    <div v-if="showLeague">
        <LeagueTableVue :name="this.currentLeague.name" :teams="this.currentLeague.teams" />
    </div>
</template>

<script>
import League from '@/JS/League';
import LeagueTableVue from '../components/LeagueTable.vue';

export default {
  name: 'App',
    data() {
        var leagues = League.retrieveLeagues(); 
    return {
        leagues: leagues,
        showLeague: false,
        currentLeague: leagues[0]
    }
  },
  methods: {
    handleNavigation(index) {
          console.log("CLICKED");
          this.currentLeague = this.leagues[index];
          this.showLeague = true;
    }
  },
    components: {
    LeagueTableVue
}
}
</script>