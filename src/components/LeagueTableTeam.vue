<script>
export default {
  name: 'LeagueTableTeam',
  props: {
    badge: String,
    name: String,
    games_played: Number,
    wins: Number,
    losses: Number
  }
}
</script>

<template>
    <div class="columns box">
        <div class="column is-one-fifths">
            <h2 class="subtitle">{{ name }}</h2>
        </div>
        <div class="column is-one-fifth">
            <h2 class="subtitle">{{ games_played }}</h2>
        </div>
        <div class="column is-one-fifth">
            <h2 class="subtitle">{{ wins }}</h2>
        </div>
        <div class="column is-one-fifth">
            <h2 class="subtitle">{{ draws }}</h2>
        </div>
        <div class="column is-one-fifth">
            <h2 class="subtitle">{{ losses }}</h2>
        </div>
    </div>
</template>