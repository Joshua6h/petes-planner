<template>
    <span>
        <v-toolbar elevation="3" color="#FCEDDA">
            <v-app-bar-nav-icon class="hidden-md-and-up"></v-app-bar-nav-icon>
            <router-link to="/">
                <v-app-bar-title>Pete's Planner</v-app-bar-title>
            </router-link>
            <v-toolbar flat color="#FCEDDA">
                <v-tabs slider-color="#EE4E34">
                    <v-tab to="/">Dashboard</v-tab>
                    <v-tab to="/about">About</v-tab>
                    <v-tab to="/calendar">Calendar</v-tab>
                    <v-tab to="/add-events">Add Events</v-tab>
                </v-tabs>
            </v-toolbar>
            <v-spacer></v-spacer>
            <v-btn plain outlined to="/login">Log In</v-btn>
            <div class="mx-3">Username</div>
            <v-btn to="/profile" icon> <!-- Make the sign in and account buttons alternate visibility based on authentication -->
                <v-icon>mdi-account-circle</v-icon>
            </v-btn>
        </v-toolbar>
    </span>
</template>

<script>
export default {
    name: 'AppNavigation',
    computed: {
    isAuthenticated() {
      return this.$store.getters.isAuthenticated;
    }
  }
}
</script>

<style scoped>
a, a:visited {
    color: black;
    text-decoration: none;
}
</style>