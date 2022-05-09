<template>
    <v-container fill-height fluid class="about">
        <v-img
        class="bg-img"
        height="100%"
        gradient="to top, rgba(255,255,255,0.5), rgba(255,255,255,0.7)"
        src="../assets/calendar-background.jpg"
        ></v-img>
        <v-row class="content mt-5">
        <v-col md="1"></v-col>
        <v-col md="10">
            <h1 class="text-h1">Welcome to Pete's Planner!</h1>
            <h1 class="banner">We are dedicated to...</h1>
            <div class="list mt-5">
            <h2 class="text-h4 pa-2">- Preventing problematic procrastination</h2>
            <h2 class="text-h4 pa-2">- Providing a personal project planner</h2>
            <h2 class="text-h4 pa-2 text-decoration-underline font-weight-medium">- Profiting profusely for private and personal gains</h2>
            </div>
        </v-col>
        </v-row>
    </v-container>
</template>

<script>
export default {
    name: "AboutHero"
}
</script>

<style scoped>
.list {
  margin: auto;
  text-align: left;
}

.banner {
  color: #EE4E34;
}

.bg-img {
  position: absolute;
}

.content {
  position: relative;
}

.about {
  background-size: cover;
  background-repeat: no-repeat;
  max-height: 100vh;
  padding: 0;
}
</style>