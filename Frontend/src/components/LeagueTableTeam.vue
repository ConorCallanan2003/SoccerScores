<script>
export default {
  name: 'LeagueTableTeam',
  props: {
    position: Number,
    badge: String,
    name: String,
    games_played: Number,
    wins: Number,
    losses: Number,
    draws: Number,
    points: Number
  }
}
</script>

<template>
    <div class="columns box is-mobile">
        <div class="column is-one-seventh">
            <h2 class="subtitle">{{ position }}</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">{{ name }}</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">{{ games_played }}</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">{{ wins }}</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">{{ draws }}</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">{{ losses }}</h2>
        </div>
        <div class="column is-one-seventh">
            <h2 class="subtitle">{{ points }}</h2>
        </div>
    </div>
</template>