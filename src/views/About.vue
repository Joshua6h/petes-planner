<template>
  <div class="about">
    <v-row class="mt-5">
      <v-col md="1"></v-col>
      <v-col md="10">
        <h1 class="text-h1">Welcome to Pete's Planner!</h1>
        <h1>We are dedicated to...</h1>
        <div class="list mt-5">
          <h2 class="text-h4 pa-2">- Preventing problematic procrastination</h2>
          <h2 class="text-h4 pa-2">- Providing a personal project planner</h2>
          <h2 class="text-h4 pa-2 text-decoration-underline font-weight-medium">- Profiting profusely for private and personal gains</h2>
        </div>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: 'About'
}
</script>

<style scoped>
.list {
  margin: auto;
  text-align: left;
}
</style>